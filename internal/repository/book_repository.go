package repository

import (
	"Library/internal/models"
	"context"
)

type BookRepository interface {
	GetByID(ctx context.Context, id int64) (models.Book, error)
	GetAll(ctx context.Context) ([]models.Book, error)
	Create(ctx context.Context, book models.Book) error
	GetByAuthorID(ctx context.Context, authorID int64) ([]models.Book, error)
}
