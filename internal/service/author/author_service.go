package author

import (
	"Library/internal/models"
	"Library/internal/repository"
	"context"
)

type authorServiceImpl struct {
	repo repository.AuthorRepository
}

func NewAuthorService(repo repository.AuthorRepository) AuthorService {
	return &authorServiceImpl{repo: repo}
}

func (s *authorServiceImpl) CreateAuthor(ctx context.Context, author *models.Author) error {
	return s.repo.Create(ctx, author)
}

func (s *authorServiceImpl) GetAllAuthors(ctx context.Context) ([]models.Author, error) {
	return s.repo.GetAll(ctx)
}
