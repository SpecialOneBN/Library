package models

type Author struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Books []Book `json:"books,omitempty"`
}

type CreateAuthorRequest struct {
	Name string `json:"name"`
}
