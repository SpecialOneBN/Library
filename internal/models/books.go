package models

type Book struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	AuthorID int64   `json:"author_id"`
	Author   *Author `json:"author,omitempty"`
}

type CreateBookRequest struct {
	Name     string `json:"name" example:"Война и мир"`
	AuthorID int64  `json:"author_id" example:"1"`
}
