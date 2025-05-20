package repository

import (
	"Library/internal/models"
	"Library/internal/repository"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type PostgreSQLUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &PostgreSQLUserRepository{db: db}
}

func (r *PostgreSQLUserRepository) GetByID(ctx context.Context, id int64) (models.User, error) {
	var user models.User
	err := r.db.QueryRowContext(ctx, `SELECT id, name, email FROM users WHERE id = $1`, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, nil
		}
		return models.User{}, fmt.Errorf("ошибка при получении пользователя по id: %w", err)
	}

	return user, nil
}

func (r *PostgreSQLUserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, name, email FROM users`)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса пользователей: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, fmt.Errorf("ошибка чтения пользователя: %w", err)
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *PostgreSQLUserRepository) GetAllUsersWithBooksSubqueries(ctx context.Context) ([]models.User, error) {
	start := time.Now()

	users, err := r.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []models.User

	for i := range users {
		rows, err := r.db.QueryContext(ctx, `
			SELECT b.id, b.name, a.id, a.name
			FROM rented_books rb
			JOIN books b ON rb.book_id = b.id
			JOIN authors a ON b.author_id = a.id
			WHERE rb.user_id = $1
		`, users[i].ID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var book models.Book
			var author models.Author
			if err := rows.Scan(&book.ID, &book.Name, &author.ID, &author.Name); err != nil {
				return nil, err
			}
			book.Author = &author
			users[i].RentedBooks = append(users[i].RentedBooks, book)
		}

		if len(users[i].RentedBooks) > 0 {
			result = append(result, users[i])
		}
	}

	log.Printf("GetAllUsersWithBooksSubqueries took %s", time.Since(start))
	return result, nil
}

func (r *PostgreSQLUserRepository) GetAllUsersWithBooksJoin(ctx context.Context) ([]models.User, error) {
	start := time.Now()

	rows, err := r.db.QueryContext(ctx, `
		SELECT u.id, u.name, u.email,
		       b.id, b.name, b.author_id,
		       a.id, a.name
		FROM users u
		LEFT JOIN rented_books rb ON u.id = rb.user_id
		LEFT JOIN books b ON rb.book_id = b.id
		LEFT JOIN authors a ON b.author_id = a.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usersMap := make(map[int64]*models.User)

	for rows.Next() {
		var userID int64
		var userName, userEmail string
		var bookID, authorID sql.NullInt64
		var bookName, authorName sql.NullString

		if err := rows.Scan(&userID, &userName, &userEmail, &bookID, &bookName, &authorID, &authorName); err != nil {
			return nil, err
		}

		user, ok := usersMap[userID]
		if !ok {
			user = &models.User{
				ID:    userID,
				Name:  userName,
				Email: userEmail,
			}
			usersMap[userID] = user
		}

		if bookID.Valid {
			book := models.Book{
				ID:   bookID.Int64,
				Name: bookName.String,
				Author: &models.Author{
					ID:   authorID.Int64,
					Name: authorName.String,
				},
			}
			user.RentedBooks = append(user.RentedBooks, book)
		}
	}

	var users []models.User
	for _, user := range usersMap {
		users = append(users, *user)
	}

	log.Printf("GetAllUsersWithBooksJoin took %s", time.Since(start))
	return users, nil
}
