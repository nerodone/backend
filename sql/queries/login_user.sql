-- name: LoginUser :one
SELECT * FROM users WHERE email = $1 AND password = $2 LIMIT 1;