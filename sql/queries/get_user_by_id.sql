-- name: GetUserByID :one
SELECT * FROM Users
WHERE id = $1 limit 1;
