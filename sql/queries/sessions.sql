-- name: AuthenicateUser :one
SELECT EXISTS (
  SELECT 1
  FROM sessions 
  WHERE user_id = sqlc.arg(userID)::uuid
  AND access_token = sqlc.arg(accessToken)::string
);
