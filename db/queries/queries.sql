-- name: ListBooks :many
SELECT * FROM books;

-- name: AddBooks :exec
INSERT INTO books (id, name)  VALUES ($1, $2);