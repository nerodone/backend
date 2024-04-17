package workspaces

import (
	"backend/server"
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
//	@Success		200	{object}	workspace
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
		workspace, err := s.Db.GetWorkspaceByID(r.Context(), id)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Failed to get workspace", "err", err.Error())
			return
		}
		s.RespondWithJson(w, http.StatusOK, WorkspaceFromDB(&workspace))
	}
}
