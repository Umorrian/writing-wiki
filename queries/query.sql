-- name: SelectArticle :one
SELECT * FROM articles
WHERE id = ? LIMIT 1;

-- name: SelectArticleByName :one
SELECT * FROM articles
WHERE name = ? LIMIT 1;

-- name: SelectAllArticles :many
SELECT * FROM articles;

-- name: InsertArticle :one
INSERT INTO articles (
name, content
) VALUES (?, ?)
RETURNING *;

-- name: UpdateArticle :exec
UPDATE articles
set name = ?,
    content = ?
WHERE id = ?;

-- name: DeleteArticle :exec
DELETE FROM articles
WHERE id = ?;