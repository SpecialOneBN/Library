package repository

import (
	"Library/internal/models"
	"context"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	GetAllUsersWithBooksJoin(ctx context.Context) ([]models.User, error)
	GetAllUsersWithBooksSubqueries(ctx context.Context) ([]models.User, error)
}
