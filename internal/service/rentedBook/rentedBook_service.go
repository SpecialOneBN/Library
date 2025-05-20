package service

import (
	"Library/internal/repository"
	"context"
	"fmt"
)

type rentedBookServiceImpl struct {
	repo repository.RentedBookRepository
}

func NewRentedBookService(repo rentedBookServiceImpl) RentedBookService {
	return &rentedBookServiceImpl{repo: repo}
}

// Проверка: выдана ли книга
func (s *rentedBookServiceImpl) IsBookIssued(ctx context.Context, bookID int64) (bool, error) {
	return s.repo.IsBookIssued(ctx, bookID)
}

// Проверка: арендовал ли пользователь эту книгу
func (s *rentedBookServiceImpl) IsBookRentedByUser(ctx context.Context, userID, bookID int64) (bool, error) {
	return s.repo.IsBookRentedByUser(ctx, userID, bookID)
}

// Выдача книги пользователю
func (s *rentedBookServiceImpl) IssueBook(ctx context.Context, userID, bookID int64) error {
	issued, err := s.repo.IsBookIssued(ctx, bookID)
	if err != nil {
		return fmt.Errorf("ошибка при проверке выдачи: %w", err)
	}
	if issued {
		return fmt.Errorf("книга уже выдана")
	}
	return s.repo.IssueBook(ctx, userID, bookID)
}

// Возврат книги
func (s *rentedBookServiceImpl) ReturnBook(ctx context.Context, userID, bookID int64) error {
	ok, err := s.repo.IsBookRentedByUser(ctx, userID, bookID)
	if err != nil {
		return fmt.Errorf("ошибка при проверке: %w", err)
	}
	if !ok {
		return fmt.Errorf("книга не числится за пользователем")
	}
	return s.repo.ReturnBook(ctx, userID, bookID)
}
