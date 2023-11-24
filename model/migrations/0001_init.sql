-- +goose Up
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME
);

CREATE TABLE articles (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    title TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME
);

CREATE TABLE article_texts (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    article_id INTEGER NOT NULL,
    content TEXT NOT NULL DEFAULT '',
    version INTEGER NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME,

    FOREIGN KEY (article_id) REFERENCES articles(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX article_version_uniq ON article_texts(article_id, version);

-- +goose Down
DROP TABLE article_texts;
DROP TABLE articles;
DROP TABLE users;