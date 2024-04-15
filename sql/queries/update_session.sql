-- name: UpdateAccessToken :one
UPDATE SESSIONS
SET access_token = $2, last_login = NOW()
WHERE id = $1
RETURNING access_token;
