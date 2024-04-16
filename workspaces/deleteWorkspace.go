package workspaces

import (
	"backend/server"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// deleteWorkspace
//
//	@Summary	delete workspace
//	@Tags		workspaces
//	@Success	200
//	@failure	400	"invalid workspace_id"
//	@failure	404	"workspace not found"
//	@Router		/workspaces/{workspace_id} [delete]
func deleteWorkspace(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		workspaceIDString := chi.URLParam(r, "workspace_id")

		workspaceID, err := uuid.Parse(workspaceIDString)
		if err != nil {
			s.RespondWithError(w, http.StatusBadRequest, "invalid workspace_id", "err", err.Error())
			return
		}

		if err := s.Db.DeleteWorkspace(r.Context(), workspaceID); err != nil {
			s.RespondWithError(w, 404, "workspace not found", "err", err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
