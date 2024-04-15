package workspaces

import (
	"backend/server"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

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
			s.RespondWithError(w, http.StatusInternalServerError, "internal server error", "err", err.Error())
			return
		}
		s.RespondWithJson(w, http.StatusOK, workspace)
	}
}
