package types

import (
	"backend/database"
	_ "backend/docs"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ID          uuid.UUID `json:"id"`
	WorkspaceID uuid.UUID `json:"workspace_id"`
	// Tasks []Task `json:"tasks"`
}

func ProjectFromDB(pDB database.Project) Project {
	desc := ""
	if pDB.Description.Valid {
		desc = pDB.Description.String
	}
	return Project{
		CreatedAt:   pDB.CreatedAt,
		UpdatedAt:   pDB.UpdatedAt,
		Name:        pDB.Name,
		ID:          pDB.ID,
		Description: desc,
		WorkspaceID: pDB.Workspace,
	}
}
