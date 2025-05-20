package user

import (
	"Library/internal/models"
	"Library/internal/repository"
	"context"
)

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) GetAll(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *userServiceImpl) GetAllUsersWithBooksJoin(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAllUsersWithBooksJoin(ctx)
}

func (s *userServiceImpl) GetAllUsersWithBooksSubqueries(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAllUsersWithBooksSubqueries(ctx)
}

/*func (s *userServiceImpl) GetUserByID(ctx context.Context, id int64) (models.User, error) {
	return s.repo.GetByID(ctx, id)
}*/
