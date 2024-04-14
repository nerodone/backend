package workspaces

import (
	"backend/database"
	"time"

	"github.com/google/uuid"
)

type workspace struct {
	ID          uuid.UUID `json:"id"`
	Owner       uuid.UUID `json:"owner"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func WorkspaceFromDB(w *database.Workspace) workspace {
	desc := ""
	if w.Description.Valid {
		desc = w.Description.String
	}
	return workspace{
		ID:          w.ID,
		Owner:       w.Owner,
		Name:        w.Name,
		Description: desc,
		CreatedAt:   w.CreatedAt,
		UpdatedAt:   w.UpdatedAt,
	}
}
