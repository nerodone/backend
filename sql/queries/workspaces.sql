-- name: CreateWorkspace :one
insert into workspaces (id, name, owner, updated_at, created_at, description) values (
  gen_random_uuid(),$1, $2, NOW(), NOW(), $3
)
returning *;

-- name: GetAllWorkspaces :many
--GetAllWorkspaces returns all workspaces that the user is the owner of or collaborates on.
SELECT *
FROM workspaces w
WHERE w.owner = sqlc.arg(user_id)::uuid
OR EXISTS (
  SELECT 1
  FROM collaborate c
  WHERE c.user_id = sqlc.arg(user_id)::uuid
  AND c.workspace_id = w.id
);

-- name: GetWorkspaceByID :one
SELECT * FROM workspaces WHERE id = $1;      


-- name: UpdateWorkspace :exec
--UpdateWorkspace can only update the workspace name and description
UPDATE workspaces SET name = $3, description = $2, updated_at = NOW() WHERE id = $1;


-- name: DeleteWorkspace :exec
DELETE FROM workspaces WHERE id = $1;

