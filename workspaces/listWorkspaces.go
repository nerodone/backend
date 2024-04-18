package workspaces

import (
	"backend/database"
	"backend/server"
	"backend/types"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
)

// listWorkspaces
//
//	@Description	List all worksapces that the user is either the owner of or collaborates on
//	@Summary		List all worksapces that the user is either the owner of or collaborates on
//	@Tags			workspaces
//	@Produce		json
//	@Success		200	{array}	types.Workspace
//	@failure		400	"invalid token"
//	@failure		401	"unauthorized access"
//	@Failure		500	"internal Server Error"
//	@Router			/workspaces [get]
func listWorkspaces(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			s.RespondWithError(w, http.StatusBadRequest, "invalid token", "err", err.Error())
			return
		}

		userID, err := uuid.Parse(claims["user_id"].(string))
		if err != nil {
			s.RespondWithError(w, http.StatusBadRequest, "invalid token", "err", err.Error())
			return
		}

		WorkspaceJoinProjectRows, err := s.Db.GetAllWorkspacesWithProjects(r.Context(), userID)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "internal server error", "err", err.Error())
			return
		}
		WorkspacesWithProjects := MatchProjectsToWorkspaces(WorkspaceJoinProjectRows)
		AllWorkspacesDB, err := s.Db.GetAllWorkspaces(r.Context(), userID)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "internal server error", "err", err.Error())
			return
		}

		AllWorkspaces := []types.Workspace{}
		for _, w := range AllWorkspacesDB {
			AllWorkspaces = append(AllWorkspaces, types.WorkspacefromDB(&w))
		}
		if len(WorkspacesWithProjects) < len(AllWorkspaces) {
			AllWorkspaces = MergeWorkspaces(AllWorkspaces, WorkspacesWithProjects)
		}
		s.RespondWithJson(w, http.StatusOK, AllWorkspaces)
	}
}

func MatchProjectsToWorkspaces(rows []database.GetAllWorkspacesWithProjectsRow) []types.Workspace {
	workspaces := []types.Workspace{}
	if len(rows) == 0 {
		return workspaces
	}

	currentWorkspace := types.WorkspacefromDB(&rows[0].Workspace)
	currentWorkspace.Projects = append(currentWorkspace.Projects, types.ProjectFromDB(rows[0].Project))
	for idx, row := range rows {

		if idx == 0 {
			continue
		}

		if rows[idx-1].Workspace.ID != row.Workspace.ID {
			workspaces = append(workspaces, currentWorkspace)
			currentWorkspace = types.WorkspacefromDB(&row.Workspace)
		}
		currentWorkspace.Projects = append(currentWorkspace.Projects, types.ProjectFromDB(row.Project))
	}
	workspaces = append(workspaces, currentWorkspace)
	return workspaces
}

func MergeWorkspaces(wAll []types.Workspace, wWithProjects []types.Workspace) []types.Workspace {
	for _, wp := range wWithProjects {
		for idx, w := range wAll {
			if wp.ID == w.ID {
				wAll[idx].Projects = wp.Projects
				break
			}
		}
	}
	return wAll
}
