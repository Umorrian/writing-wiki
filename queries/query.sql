-- name: SelectArticleWithCurrentContent :one
SELECT sqlc.embed(a), sqlc.embed(at) FROM articles as a
JOIN article_texts at on a.id = at.article_id
WHERE a.id = ?
ORDER BY at.version DESC
LIMIT 1;

-- name: SelectArticleByNameWithCurrentContent :one
SELECT sqlc.embed(a), sqlc.embed(at) FROM articles as a
JOIN article_texts at on a.id = at.article_id
WHERE a.title = ?
ORDER BY at.version DESC
LIMIT 1;

-- name: SelectAllArticles :many
SELECT * FROM articles;

-- name: InsertArticle :one
INSERT INTO articles (
title, created_at
) VALUES (?, DATETIME())
RETURNING *;

-- name: UpdateArticle :exec
UPDATE articles
set title = ?,
    updated_at = DATETIME()
WHERE id = ?;

-- name: DeleteArticle :exec
DELETE FROM articles
WHERE id = ?;

-- name: SelectArticleTextVersionsByArticleId :many
SELECT * FROM article_texts
where article_id = ?;

-- name: InsertArticleTextVersion :one
INSERT INTO article_texts (
article_id, content, version, created_at
) VALUES (
@article_id,
?,
(SELECT MAX(version)+1
 FROM article_texts as subat
 WHERE subat.id=@article_id
 GROUP BY subat.version),
DATETIME())
RETURNING *;