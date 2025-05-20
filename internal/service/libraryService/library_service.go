package service

import (
	"Library/internal/repository"
	"context"
	"fmt"
)

// LibraryService - суперсервис, объединяющий работу с книгами, пользователями и арендой
// Этот слой инкапсулирует бизнес-логику и использует репозитории как интерфейсы

type libraryServiceImpl struct {
	UserRepo   repository.UserRepository
	BookRepo   repository.BookRepository
	AuthorRepo repository.AuthorRepository
	RentedRepo repository.RentedBookRepository
}

// NewLibraryService - конструктор суперсервиса
func NewLibraryService(userRepo repository.UserRepository, bookRepo repository.BookRepository, authorRepo repository.AuthorRepository, rentedRepo repository.RentedBookRepository) LibraryService {
	return &libraryServiceImpl{
		UserRepo:   userRepo,
		BookRepo:   bookRepo,
		AuthorRepo: authorRepo,
		RentedRepo: rentedRepo,
	}
}

// GiveBook - логика выдачи книги пользователю
func (s *libraryServiceImpl) GiveBook(ctx context.Context, userID, bookID int64) error {
	// Проверка существования пользователя
	user, err := s.UserRepo.GetByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("пользователь не найден: %w", err)
	}

	// Проверка существования книги
	book, err := s.BookRepo.GetByID(ctx, bookID)
	if err != nil {
		return fmt.Errorf("книга не найдена: %w", err)
	}

	// Проверка, выдана ли уже книга
	issued, err := s.RentedRepo.IsBookIssued(ctx, bookID)
	if err != nil {
		return fmt.Errorf("ошибка при проверке аренды: %w", err)
	}
	if issued {
		return fmt.Errorf("книга уже выдана другому пользователю")
	}

	// Всё ок — записываем аренду
	return s.RentedRepo.IssueBook(ctx, user.ID, book.ID)
}

// ReturnBook - логика возврата книги пользователем
func (s *libraryServiceImpl) ReturnBook(ctx context.Context, userID, bookID int64) error {
	// Проверяем, что пользователь вообще арендовал эту книгу
	ok, err := s.RentedRepo.IsBookRentedByUser(ctx, userID, bookID)
	if err != nil {
		return fmt.Errorf("ошибка при проверке аренды: %w", err)
	}
	if !ok {
		return fmt.Errorf("книга не числится за пользователем")
	}

	// Удаляем запись из rented_books
	return s.RentedRepo.ReturnBook(ctx, userID, bookID)
}
