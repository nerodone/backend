-- name: AuthenicateUser :one
SELECT EXISTS (
  SELECT 1
  FROM sessions 
  WHERE user_id = $1
  AND access_token = $2
);
