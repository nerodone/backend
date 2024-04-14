-- name: CreateWorkspace :one
insert into workspaces (id, name, owner, updated_at, created_at, description) values (
  gen_random_uuid(),$1, $2, NOW(), NOW(), $3
)
returning *;
