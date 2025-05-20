package handler

import (
	"Library/internal/facade"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type LibraryHandler struct {
	facade *facade.LibraryFacade
}

func NewLibraryHandler(f *facade.LibraryFacade) *LibraryHandler {
	return &LibraryHandler{facade: f}
}

// TakeBook godoc
// @Summary Выдать книгу пользователю
// @Tags library
// @Param user_id query int true "ID пользователя"
// @Param book_id query int true "ID книги"
// @Success 200 {object} models.Rental
// @Success 200 {string} string "Книга успешно выдана"
// @Failure 400 {string} string "Неверные параметры"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /library/take [post]
func (h *LibraryHandler) TakeBook(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		http.Error(w, "Неверный user_id", http.StatusBadRequest)
		return
	}

	bookID, err := strconv.ParseInt(r.URL.Query().Get("book_id"), 10, 64)
	if err != nil {
		http.Error(w, "Неверный book_id", http.StatusBadRequest)
		return
	}

	if err := h.facade.TakeBook(r.Context(), userID, bookID); err != nil {
		log.Printf("Ошибка при выдаче книги: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Книга успешно выдана")
}

// ReturnBook godoc
// @Summary Вернуть книгу
// @Tags library
// @Param user_id query int true "ID пользователя"
// @Param book_id query int true "ID книги"
// @Success 200 {object} models.Rental
// @Success 200 {string} string "Книга успешно возвращена"
// @Failure 400 {string} string "Неверные параметры"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /library/return [post]
func (h *LibraryHandler) ReturnBook(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		http.Error(w, "Неверный user_id", http.StatusBadRequest)
		return
	}

	bookID, err := strconv.ParseInt(r.URL.Query().Get("book_id"), 10, 64)
	if err != nil {
		http.Error(w, "Неверный book_id", http.StatusBadRequest)
		return
	}

	if err := h.facade.GiveBackBook(r.Context(), userID, bookID); err != nil {
		log.Printf("Ошибка при возврате книги: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Книга успешно возвращена")
}
