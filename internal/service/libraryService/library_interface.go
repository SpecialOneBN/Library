package service

import "context"

type LibraryService interface {
	GiveBook(ctx context.Context, userID, bookID int64) error
	ReturnBook(ctx context.Context, userID, bookID int64) error
}
