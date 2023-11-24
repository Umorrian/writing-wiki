package model

import (
	"database/sql"
	"time"
)

type Article struct {
	ID          int64        `db:"id"`
	Title       string       `db:"title"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
	CurrentText *ArticleText `db:"current_text"`
}

type ArticleText struct {
	ID        int64        `db:"id"`
	ArticleID int64        `db:"article_id"`
	Content   string       `db:"content"`
	Version   int64        `db:"version"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type User struct {
	ID        int64
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
