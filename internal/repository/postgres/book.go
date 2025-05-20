package repository

import (
	"Library/internal/models"
	"Library/internal/repository"
	"context"
	"database/sql"
	"fmt"
)

type PostgreSQLBookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) repository.BookRepository {
	return &PostgreSQLBookRepository{db: db}
}

func (r *PostgreSQLBookRepository) Create(ctx context.Context, book models.Book) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO books (name, author_id)
		VALUES ($1, $2)
	`, book.Name, book.Author.ID)

	if err != nil {
		return fmt.Errorf("ошибка при создании книги: %w", err)
	}
	return nil
}

func (r *PostgreSQLBookRepository) GetByID(ctx context.Context, id int64) (models.Book, error) {
	var book models.Book
	var author models.Author

	err := r.db.QueryRowContext(ctx, `
		SELECT b.id, b.name, a.id, a.name
		FROM books b
		JOIN authors a ON b.author_id = a.id
		WHERE b.id = $1
	`, id).Scan(&book.ID, &book.Name, &author.ID, &author.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Book{}, nil
		}
		return models.Book{}, fmt.Errorf("ошибка при получении книги по id: %w", err)
	}

	book.Author = &author
	return book, nil
}

func (r *PostgreSQLBookRepository) GetAll(ctx context.Context) ([]models.Book, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT b.id, b.name, a.id, a.name
		FROM books b
		JOIN authors a ON b.author_id = a.id
	`)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении всех книг: %w", err)
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		var author models.Author

		if err := rows.Scan(&book.ID, &book.Name, &author.ID, &author.Name); err != nil {
			return nil, fmt.Errorf("ошибка при сканировании книги: %w", err)
		}

		book.Author = &author
		books = append(books, book)
	}

	return books, nil
}

func (r *PostgreSQLBookRepository) GetByAuthorID(ctx context.Context, authorID int64) ([]models.Book, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT b.id, b.name, a.id, a.name
		FROM books b
		JOIN authors a ON b.author_id = a.id
		WHERE b.author_id = $1
	`, authorID)

	if err != nil {
		return nil, fmt.Errorf("ошибка при получении книг по автору: %w", err)
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		var author models.Author

		if err := rows.Scan(&book.ID, &book.Name, &author.ID, &author.Name); err != nil {
			return nil, fmt.Errorf("ошибка при сканировании книги: %w", err)
		}

		book.Author = &author
		books = append(books, book)
	}

	return books, nil
}
