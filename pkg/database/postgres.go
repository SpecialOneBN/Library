package database

import (
	"Library/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.GetDSN())
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ping не удался: %w", err)
	}

	return db, nil
}
