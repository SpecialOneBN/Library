package repository

import (
	"Library/internal/models"
	"Library/internal/repository"
	"context"
	"database/sql"
	"fmt"
)

type PostgreSQLAuthorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) repository.AuthorRepository {
	return &PostgreSQLAuthorRepository{db: db}
}

func (r *PostgreSQLAuthorRepository) GetAll(ctx context.Context) ([]models.Author, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT a.id, a.name, b.id, b.name
		FROM authors a
		LEFT JOIN books b ON a.id = b.author_id
	`)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении авторов и книг: %w", err)
	}
	defer rows.Close()

	authorsMap := make(map[int64]*models.Author)

	for rows.Next() {
		var authorID int64
		var authorName string
		var bookID sql.NullInt64
		var bookName sql.NullString

		if err := rows.Scan(&authorID, &authorName, &bookID, &bookName); err != nil {
			return nil, err
		}

		author, exists := authorsMap[authorID]
		if !exists {
			author = &models.Author{
				ID:    authorID,
				Name:  authorName,
				Books: []models.Book{},
			}
			authorsMap[authorID] = author
		}

		if bookID.Valid {
			author.Books = append(author.Books, models.Book{
				ID:       bookID.Int64,
				Name:     bookName.String,
				AuthorID: authorID,
			})
		}
	}

	var authors []models.Author
	for _, a := range authorsMap {
		authors = append(authors, *a)
	}

	return authors, nil
}

func (r *PostgreSQLAuthorRepository) GetByID(ctx context.Context, id int64) (models.Author, error) {
	var author models.Author
	err := r.db.QueryRowContext(ctx, `SELECT id, name FROM authors WHERE id = $1`, id).Scan(&author.ID, &author.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Author{}, nil
		}
		return models.Author{}, fmt.Errorf("ошибка при получении автора по ID: %w", err)
	}
	return author, nil
}

func (r *PostgreSQLAuthorRepository) Create(ctx context.Context, author *models.Author) error {
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO authors (name) VALUES ($1) RETURNING id`,
		author.Name).Scan(&author.ID)
	if err != nil {
		return fmt.Errorf("ошибка при создании автора: %w", err)
	}
	return nil
}
