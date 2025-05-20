-- +goose Up
-- +goose StatementBegin
CREATE TABLE books(
    id SERIAL PRIMARY KEY,
    name VARCHAR(250),
    authors_id INTEGER NOT NULL REFERENCES authors (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS books;
-- +goose StatementEnd
