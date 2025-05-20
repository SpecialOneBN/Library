package book

import (
	"Library/internal/models"
	"context"
)

type BookService interface {
	CreateBook(ctx context.Context, b models.Book) error
	GetAllBooks(ctx context.Context) ([]models.Book, error)
}
