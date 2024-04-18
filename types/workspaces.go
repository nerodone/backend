package types

import (
	"backend/database"
	"time"

	_ "backend/docs"

	"github.com/google/uuid"
)

type Workspace struct {
	ID          uuid.UUID `json:"id"`
	Owner       uuid.UUID `json:"owner"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Projects    []Project `json:"projects"`
}

func WorkspacefromDB(w *database.Workspace) Workspace {
	desc := ""
	if w.Description.Valid {
		desc = w.Description.String
	}
	return Workspace{
		ID:          w.ID,
		Owner:       w.Owner,
		Name:        w.Name,
		Description: desc,
		CreatedAt:   w.CreatedAt,
		UpdatedAt:   w.UpdatedAt,
	}
}

// MapProjectsInWorkpsace checks if any of the projects ps belong to workspace w
// if so it adds to them to w and deletes them from
func MapProjectsInWorkpsace(w Workspace, ps ...Project) (Workspace, []Project) {
	w.Projects = []Project{}
	remainingProjects := []Project{}
	for _, p := range ps {
		if p.WorkspaceID == w.ID {
			w.Projects = append(w.Projects, p)
		} else {
			remainingProjects = append(remainingProjects, p)
		}
	}
	return w, remainingProjects
}
