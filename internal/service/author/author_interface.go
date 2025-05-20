package author

import (
	"Library/internal/models"
	"context"
)

type AuthorService interface {
	CreateAuthor(ctx context.Context, author *models.Author) error
	GetAllAuthors(ctx context.Context) ([]models.Author, error)
}
