package book

import (
	"Library/internal/models"
	"Library/internal/repository"
	"context"
	"fmt"
)

type bookServiceImpl struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookServiceImpl{repo: repo}
}

func (s *bookServiceImpl) CreateBook(ctx context.Context, book models.Book) error {
	if book.Name == "" {
		return fmt.Errorf("название книги не может быть пустым")
	}
	if book.Author == nil || book.Author.ID == 0 {
		return fmt.Errorf("не указан автор")
	}

	// Проверяем, существует ли автор
	_, err := s.repo.GetByID(ctx, book.Author.ID)
	if err != nil {
		return fmt.Errorf("автор с ID %d не найден", book.Author.ID)
	}

	return s.repo.Create(ctx, book)
}

func (s *bookServiceImpl) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	return s.repo.GetAll(ctx)
}
