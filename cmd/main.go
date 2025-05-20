package main

import (
	_ "Library/docs"
	"Library/internal/config"
	"Library/internal/facade"
	"Library/internal/handler"
	"Library/internal/initdata"
	"Library/internal/repository/postgres"
	"Library/internal/service/author"
	"Library/internal/service/book"
	service "Library/internal/service/libraryService"
	"Library/internal/service/user"
	"Library/pkg/database"
	"log"
	"net/http"
)

// @title Library API
// @version 1.0
// @description API для библиотеки с суперсервисами и фасадом. Выдача, возврат, авторы и книги.

// @contact.name Nikolay Korotaev
// @contact.url http://localhost:8080
// @contact.email Korotallie@gmail.com

// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.LoadConfig()

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}
	defer db.Close()

	log.Println("База успешно подключена")

	if err := database.RunMigrations(db, "./migrations"); err != nil {
		log.Fatalf("Миграция не удалась: %v", err)
	}

	initdata.InitDatabase(db)

	log.Println("База успешно подключена и миграции применены")

	// Репозитории
	userRepo := repository.NewUserRepository(db)
	bookRepo := repository.NewBookRepository(db)
	authorRepo := repository.NewAuthorRepository(db)
	rentedRepo := repository.NewRentedBookRepository(db)

	// Сервисы
	uService := user.NewUserService(userRepo)
	bService := book.NewBookService(bookRepo)
	aService := author.NewAuthorService(authorRepo)
	librarySvc := service.NewLibraryService(userRepo, bookRepo, authorRepo, rentedRepo)

	// Фасад
	libraryFacade := facade.NewLibraryFacade(librarySvc)

	// Роутер
	router := handler.NewRouter(uService, bService, aService, libraryFacade)

	log.Println("Сервер запущен на :8080 🚀")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}
