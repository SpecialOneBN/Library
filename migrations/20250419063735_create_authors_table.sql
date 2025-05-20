-- +goose Up
-- +goose StatementBegin
CREATE TABLE authors(
     id SERIAL PRIMARY KEY,
     name VARCHAR(200) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS authors;
-- +goose StatementEnd
