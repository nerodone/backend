package workspaces

import (
	"backend/database"
	"backend/server"
	"backend/types"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// listWorkspaces
//
//	@Description	List all worksapces that the user is either the owner of or collaborates on
//	@Summary		List all worksapces that the user is either the owner of or collaborates on
//	@Tags			workspaces
//	@Produce		json
//	@Accepts		json
//	@Success		200	{object}	types.Workspace
//	@failure		401	"unauthorized access"
//	//	@failure		400	"Failed to get workspace"
//	@Failure		500	"internal Server Error"
//	@Router			/workspaces/{workspace_id} [get]
func getWorkspaceByID(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		workspace_id := chi.URLParam(r, "workspace_id")
		id, err := uuid.Parse(workspace_id)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "internal server error", "err", err.Error())
			return
		}

		WokrpsaceJoinProjectsRows, err := s.Db.GetWorkspaceWithProjectsByID(r.Context(), id)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Failed to get workspace", "err", err.Error())
			return
		}

		if WokrpsaceJoinProjectsRows == nil {
			s.RespondWithJson(w, 200, []string{})
			return
		}
		workspace := NestProjectInWokrspace(WokrpsaceJoinProjectsRows)
		s.RespondWithJson(w, 200, workspace)
	}
}

func NestProjectInWokrspace(rows []database.GetWorkspaceWithProjectsByIDRow) types.Workspace {
	workspace := rows[0].Workspace
	projects := []types.Project{}
	for _, row := range rows {
		projects = append(projects, types.ProjectFromDB(row.Project))
	}
	WorkspaceWithProjects, remainingProjects := types.MapProjectsInWorkpsace(types.WorkspacefromDB(&workspace), projects...)
	if len(remainingProjects) != 0 {
		log.Panicf("remaining projects projectName : %v", remainingProjects[0].Name)
	}
	return WorkspaceWithProjects
}
