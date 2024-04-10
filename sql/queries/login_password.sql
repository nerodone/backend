-- name: LoginUser :one
SELECT id FROM passwordLogin WHERE email = @email AND password = @password LIMIT 1;
