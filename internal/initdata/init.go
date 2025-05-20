package initdata

import (
	"database/sql"
	"github.com/brianvoe/gofakeit/v6"
	"log"
)

func InitDatabase(db *sql.DB) {
	gofakeit.Seed(0)

	if isTableEmpty(db, "users") {
		log.Println("Таблица users пуста. Идет заполнение...")
		for i := 0; i < 50; i++ {
			_, err := db.Exec(`INSERT INTO users (name, email) VALUES ($1, $2)`, gofakeit.Name(), gofakeit.Email())
			if err != nil {
				log.Fatalf("Ошибка при вставке пользователя: %v", err)
			}
		}
	}

	if isTableEmpty(db, "authors") {
		log.Println("Таблица authors пуста. Идет заполнение...")
		for i := 0; i < 10; i++ {
			_, err := db.Exec(`INSERT INTO authors (name) VALUES ($1)`, gofakeit.Name())
			if err != nil {
				log.Fatalf("Ошибка при вставке автора: %v", err)
			}
		}
	}

	if isTableEmpty(db, "books") {
		log.Println("Таблица books пуста. Заполняем...")
		rows, err := db.Query(`SELECT id FROM authors`)
		if err != nil {
			log.Fatalf("Ошибка при получении авторов: %v", err)
		}
		defer rows.Close()

		var authorIDs []int
		for rows.Next() {
			var id int
			_ = rows.Scan(&id)
			authorIDs = append(authorIDs, id)
		}

		for i := 0; i < 100; i++ {
			authorID := authorIDs[gofakeit.Number(0, len(authorIDs)-1)]
			_, err := db.Exec(`INSERT INTO books  (name, author_id) VALUES ($1, $2)`, gofakeit.BookTitle(), authorID)
			if err != nil {
				log.Fatalf("Ошибка при вставке книги: %v", err)
			}
		}
	}

}

func isTableEmpty(db *sql.DB, table string) bool {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM ` + table).Scan(&count)
	if err != nil {
		log.Fatalf("ОШибка при проверке таблицы %s: %v", table, err)
	}
	return count == 0
}
