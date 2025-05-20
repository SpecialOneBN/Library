package facade

import (
	service "Library/internal/service/libraryService"
	"context"
)

type LibraryFacade struct {
	library service.LibraryService
}

func NewLibraryFacade(library service.LibraryService) *LibraryFacade {
	return &LibraryFacade{library: library}
}

func (f *LibraryFacade) TakeBook(ctx context.Context, userID, bookID int64) error {
	return f.library.GiveBook(ctx, userID, bookID)
}

func (f *LibraryFacade) GiveBackBook(ctx context.Context, userID, bookID int64) error {
	return f.library.ReturnBook(ctx, userID, bookID)
}
