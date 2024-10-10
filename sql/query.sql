-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1;

-- name: CreateAuthor :one
INSERT INTO authors (name, email)
VALUES ($1, $2) RETURNING *;

-- name: GetAllAuthors :many
SELECT * FROM authors;

-- name: CreatePost :one
INSERT INTO posts (title, content, author_id)
VALUES ($1, $2, $3) RETURNING *;

-- name: DeletePost :exec
DELETE from posts
WHERE id = $1;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;

-- name: GetPost :one
SELECT * FROM posts WHERE id = $1 LIMIT 1;

-- name: GetPosts :many
SELECT * FROM posts;

-- name: UpdatePost :exec
UPDATE posts SET title = $2, content = $3, updated_at = NOW()
WHERE id = $1;

-- name: CreateComment :one
INSERT INTO comments (content, author_id, post_id)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetCommentsForAuthor :many
SELECT * FROM comments
WHERE author_id = $1;

-- name: GetCommentsForPost :many
SELECT * FROM comments
WHERE post_id = $1;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;