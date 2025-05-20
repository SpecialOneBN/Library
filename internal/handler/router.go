package handler

import (
	"Library/internal/facade"
	"Library/internal/service/author"
	"Library/internal/service/book"
	"Library/internal/service/user"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func NewRouter(
	userService user.UserService,
	bookService book.BookService,
	authorService author.AuthorService,
	libraryFacade *facade.LibraryFacade,
) http.Handler {
	r := chi.NewRouter()

	userHandler := NewUserHandler(userService)
	bookHandler := NewBookHandler(bookService)
	authorHandler := NewAuthorHandler(authorService)
	libraryHandler := NewLibraryHandler(libraryFacade)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Get("/users/join", userHandler.GetAllUsersWithJoin)
	r.Get("/users/subquery", userHandler.GetAllUsersWithSubquery)

	r.Get("/books", bookHandler.GetAllBooks)
	r.Post("/books", bookHandler.CreateBook)

	r.Get("/authors", authorHandler.GetAllAuthors)
	r.Post("/authors", authorHandler.CreateAuthor)

	// Маршруты библиотеки (фасад)
	r.Post("/library/take", libraryHandler.TakeBook)
	r.Post("/library/return", libraryHandler.ReturnBook)

	return r
}
