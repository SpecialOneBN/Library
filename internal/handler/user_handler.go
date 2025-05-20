package handler

import (
	"Library/internal/service/user"
	"encoding/json"
	"log"
	"net/http"
)

type UserHandler struct {
	service user.UserService
}

func NewUserHandler(service user.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAllUsersWithJoin(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsersWithBooksJoin(r.Context())
	if err != nil {
		log.Printf("Ошибка при получении пользователей с JOIN: %v", err)
		http.Error(w, "Не удалось получить пользователей", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(users)
}

// GetAllUsersWithSubquery godoc
// @Summary Получить всех пользователей с книгами (через подзапросы)
// @Description Возвращает всех пользователей, включая арендуемые ими книги
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {string} string "Не удалось получить пользователей"
// @Router /users/subquery [get]
func (h *UserHandler) GetAllUsersWithSubquery(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsersWithBooksSubqueries(r.Context())
	if err != nil {
		log.Printf("Ошибка при получении пользователей с подзапросами: %v", err)
		http.Error(w, "Не удалось получить пользователей", http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(users)
}
