package handler

import (
	"Library/internal/models"
	"Library/internal/service/author"
	"encoding/json"
	"net/http"
)

type AuthorHandler struct {
	service author.AuthorService
}

func NewAuthorHandler(service author.AuthorService) *AuthorHandler {
	return &AuthorHandler{service: service}
}

// CreateAuthor godoc
// @Summary Добавить нового автора
// @Tags authors
// @Accept json
// @Produce json
// @Param author body models.CreateAuthorRequest true "Автор"
// @Success 201 {object} models.Author
// @Failure 400 {string} string "Неверный ввод"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /authors [post]
func (h *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var req models.CreateAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	author := &models.Author{Name: req.Name}
	if err := h.service.CreateAuthor(r.Context(), author); err != nil {
		http.Error(w, "Failed to create author", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(author)
}

// GetAllAuthors godoc
// @Summary Получить список всех авторов
// @Tags authors
// @Produce json
// @Success 200 {array} models.Author
// @Router /authors [get]
func (h *AuthorHandler) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := h.service.GetAllAuthors(r.Context())
	if err != nil {
		http.Error(w, "Failed to get authors", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(authors)
}
