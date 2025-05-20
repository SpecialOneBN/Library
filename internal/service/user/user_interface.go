package user

import (
	"Library/internal/models"
	"context"
)

type UserService interface {
	GetAll(ctx context.Context) ([]models.User, error)
	//	GetUserByID(ctx context.Context, id int64) (models.User, error)
	GetAllUsersWithBooksJoin(ctx context.Context) ([]models.User, error)
	GetAllUsersWithBooksSubqueries(ctx context.Context) ([]models.User, error)
}
