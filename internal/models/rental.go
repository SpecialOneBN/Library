package models

type Rental struct {
	UserID   int64  `json:"user_id"`
	BookID   int64  `json:"book_id"`
	RentedAt string `json:"rented_at"`
}
