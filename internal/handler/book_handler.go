package handler

import (
	"Library/internal/models"
	"Library/internal/service/book"
	"encoding/json"
	"net/http"
)

type BookHandler struct {
	service book.BookService
}

func NewBookHandler(service book.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// CreateBook godoc
// @Summary Добавить новую книгу
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.CreateBookRequest true "Книга"
// @Success 201 {object} models.Book
// @Failure 400 {string} string "Неверные данные"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /books [post]
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var req models.CreateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}

	book := models.Book{
		Name: req.Name,
		Author: &models.Author{
			ID: req.AuthorID,
		},
	}

	if err := h.service.CreateBook(r.Context(), book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(book)
}

// GetAllBooks godoc
// @Summary Получить список всех книг
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetAllBooks(r.Context())
	if err != nil {
		http.Error(w, "Failed to get books", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(books)
}
