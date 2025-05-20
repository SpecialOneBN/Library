package repository

import (
	"Library/internal/repository"
	"context"
	"database/sql"
	"fmt"
)

type PostgreSQLRentedBookRepository struct {
	db *sql.DB
}

func NewRentedBookRepository(db *sql.DB) repository.RentedBookRepository {
	return &PostgreSQLRentedBookRepository{db: db}
}

// Проверяет, выдана ли книга (есть ли запись по этому book_id)
func (r *PostgreSQLRentedBookRepository) IsBookIssued(ctx context.Context, bookID int64) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS(
			SELECT 1 FROM rented_books WHERE book_id = $1
		)
	`
	err := r.db.QueryRowContext(ctx, query, bookID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("ошибка при проверке аренды книги: %w", err)
	}
	return exists, nil
}

// Проверяет, выдана ли книга конкретному пользователю
func (r *PostgreSQLRentedBookRepository) IsBookRentedByUser(ctx context.Context, userID, bookID int64) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS(
			SELECT 1 FROM rented_books WHERE user_id = $1 AND book_id = $2
		)
	`
	err := r.db.QueryRowContext(ctx, query, userID, bookID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("ошибка при проверке аренды книги пользователем: %w", err)
	}
	return exists, nil
}

// Выдаёт книгу пользователю (создаёт запись)
func (r *PostgreSQLRentedBookRepository) IssueBook(ctx context.Context, userID, bookID int64) error {
	query := `
		INSERT INTO rented_books (user_id, book_id)
		VALUES ($1, $2)
	`
	_, err := r.db.ExecContext(ctx, query, userID, bookID)
	if err != nil {
		return fmt.Errorf("ошибка при выдаче книги: %w", err)
	}
	return nil
}

// Возвращает книгу (удаляет запись)
func (r *PostgreSQLRentedBookRepository) ReturnBook(ctx context.Context, userID, bookID int64) error {
	query := `
		DELETE FROM rented_books
		WHERE user_id = $1 AND book_id = $2
	`
	_, err := r.db.ExecContext(ctx, query, userID, bookID)
	if err != nil {
		return fmt.Errorf("ошибка при возврате книги: %w", err)
	}
	return nil
}
