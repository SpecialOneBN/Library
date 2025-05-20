package database

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose/v3"
)

func RunMigrations(db *sql.DB, migrationsDir string) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("не удалось установить диалект goose: %w", err)
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		return fmt.Errorf("ошибка миграции: %w", err)
	}

	return nil
}
