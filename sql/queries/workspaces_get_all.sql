
-- name: GetAllWorkspaces :many
--GetAllWorkspaces returns all workspaces that the user is the owner of or collaborates on.
SELECT *
FROM workspaces w
WHERE w.owner = sqlc.arg(user_id)::uuid
OR EXISTS (
  SELECT 1
  FROM collaborate c
  WHERE c.user = sqlc.arg(user_id)::uuid
  AND c.workspace = w.id
);
