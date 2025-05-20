-- +goose Up
-- +goose StatementBegin
CREATE TABLE rented_books (
    user_id INTEGER NOT NULL REFERENCES users (id),
    book_id INTEGER NOT NULL UNIQUE REFERENCES books(id),
    rented_at TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (user_id, book_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rented_books;
-- +goose StatementEnd
