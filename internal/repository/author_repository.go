package repository

import (
	"Library/internal/models"
	"context"
)

type AuthorRepository interface {
	GetByID(ctx context.Context, id int64) (models.Author, error)
	GetAll(ctx context.Context) ([]models.Author, error)
	Create(ctx context.Context, author *models.Author) error
}
