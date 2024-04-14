// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: workspaces.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createWorkspace = `-- name: CreateWorkspace :one
insert into workspaces (id, name, owner, updated_at, created_at, description) values (
  gen_random_uuid(),$1, $2, NOW(), NOW(), $3
)
returning id, owner, name, description, created_at, updated_at
`

type CreateWorkspaceParams struct {
	Name        string         `json:"name"`
	Owner       uuid.UUID      `json:"owner"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) CreateWorkspace(ctx context.Context, arg CreateWorkspaceParams) (Workspace, error) {
	row := q.db.QueryRowContext(ctx, createWorkspace, arg.Name, arg.Owner, arg.Description)
	var i Workspace
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllWorkspaces = `-- name: GetAllWorkspaces :many
SELECT id, owner, name, description, created_at, updated_at
FROM workspaces w
WHERE w.owner = $1::uuid
OR EXISTS (
  SELECT 1
  FROM collaborate c
  WHERE c.user = $1::uuid
  AND c.workspace = w.id
)
`

// GetAllWorkspaces returns all workspaces that the user is the owner of or collaborates on.
func (q *Queries) GetAllWorkspaces(ctx context.Context, userID uuid.UUID) ([]Workspace, error) {
	rows, err := q.db.QueryContext(ctx, getAllWorkspaces, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Workspace
	for rows.Next() {
		var i Workspace
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getWorkspaceByID = `-- name: GetWorkspaceByID :one
SELECT id, owner, name, description, created_at, updated_at FROM workspaces WHERE id = $1
`

func (q *Queries) GetWorkspaceByID(ctx context.Context, id uuid.UUID) (Workspace, error) {
	row := q.db.QueryRowContext(ctx, getWorkspaceByID, id)
	var i Workspace
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
