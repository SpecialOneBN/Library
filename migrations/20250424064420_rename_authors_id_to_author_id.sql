-- +goose Up
-- +goose StatementBegin
ALTER TABLE books RENAME COLUMN authors_id TO author_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE books RENAME COLUMN author_id TO authors_id;
-- +goose StatementEnd
