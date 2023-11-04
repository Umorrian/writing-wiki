-- +goose Up
CREATE TABLE articles (
     id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
     name text NOT NULL
);

-- +goose Down
DROP TABLE articles;