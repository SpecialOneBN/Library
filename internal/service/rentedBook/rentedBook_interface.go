package service

import (
	"context"
)

type RentedBookService interface {
	IsBookIssued(ctx context.Context, bookID int64) (bool, error)
	IsBookRentedByUser(ctx context.Context, userID, bookID int64) (bool, error)
	IssueBook(ctx context.Context, userID, bookID int64) error
	ReturnBook(ctx context.Context, userID, bookID int64) error
}
