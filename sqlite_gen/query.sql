-- name: GetArticle :one
SELECT * FROM articles
WHERE id = ? LIMIT 1;

-- name: GetArticleByName :one
SELECT * FROM articles
WHERE name = ? LIMIT 1;

-- name: GetAllArticles :many
SELECT * FROM articles;

-- name: CreateArticle :one
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