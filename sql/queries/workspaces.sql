-- name: CreateWorkspace :one
insert into workspaces (id, name, owner, updated_at, created_at, description) values (
  gen_random_uuid(),$1, $2, NOW(), NOW(), $3
)
returning *;


--GetAllWorkspacesWithProjects returns all workspaces that contain projects where the user is the owner of or collaborates on the workspace.
-- name: GetAllWorkspacesWithProjects :many
SELECT sqlc.embed(workspaces), sqlc.embed(projects)
FROM workspaces
JOIN projects ON projects.workspace = workspaces.id
WHERE workspaces.owner = sqlc.arg(user_id)::uuid
OR EXISTS (
  SELECT 1
  FROM collaborate c
  WHERE c.user_id = sqlc.arg(user_id)::uuid
  AND c.workspace_id = workspaces.id
);
-- name: GetAllWorkspaces :many
--GetAllWorkspaces doesnt return projects assosiated to each workspace
SELECT *
FROM workspaces w
WHERE w.owner = $1
OR EXISTS (
  SELECT 1
  FROM collaborate c
  WHERE c.user_id = $1
  AND c.workspace_id = w.id
);


-- name: GetWorkspaceWithProjectsByID :many
SELECT sqlc.embed(workspaces), sqlc.embed(projects) 
FROM workspaces JOIN projects ON projects.workspace_id = workspaces.id 
WHERE workspaces.id = $1;


-- name: GetWorkspaceByID :one
-- GetWorkspaceByID doesnt reuturn projects asssiated with the project
SELECT * FROM workspaces WHERE id = $1;




-- name: UpdateWorkspace :exec
--UpdateWorkspace can only update the workspace name and description
UPDATE workspaces SET name = $3, description = $2, updated_at = NOW() WHERE id = $1;


-- name: DeleteWorkspace :exec
DELETE FROM workspaces WHERE id = $1;

