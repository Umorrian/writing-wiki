-- +goose Up
ALTER TABLE articles ADD content text NOT NULL;

-- +goose Down
ALTER TABLE articles DROP COLUMN content;