-- name: GetSessionById :one
SELECT * FROM SESSIONS
WHERE id = $1 limit 1;
