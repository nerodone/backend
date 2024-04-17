package workspaces

import (
	"backend/database"
	"backend/server"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type updateWorkspaceReq struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	WorkspaceID uuid.UUID      `json:"workspace_id"`
}

// updateWorkspace
//
//	@Summary	update workspace metadata (name, description) , empty fields are ignored
//	@Tags		workspaces
//	@Param		request	body	updateWorkspaceReq	true	" "
//	@Accepts	json
//	@Success	204
//	@failure	400	"invalid request payload||invalid workspace_id"
//	@failure	401	"unauthorized access"
//	@failure	404	"workspace not found"
//	@Failure	500	"internal Server Error"
//	@Router		/workspaces/{workspace_id} [put]
func updateWorkspace(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		workspaceIDString := chi.URLParam(r, "workspace_id")
		workspaceID, err := uuid.Parse(workspaceIDString)
		if err != nil {
			s.RespondWithError(w, 400, "invalid workspace_id", "err", err.Error(), "id", workspaceIDString)
			return
		}
		req := updateWorkspaceReq{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.RespondWithError(w, http.StatusBadRequest, "invalid request payload", "err", err.Error())
			return
		}

		err = s.Db.UpdateWorkspace(r.Context(), database.UpdateWorkspaceParams{
			ID:          workspaceID,
			Description: req.Description,
			Name:        req.Name,
		})
		if err != nil {
			s.RespondWithError(w, http.StatusNotFound, "workspace not found", "err", err.Error())
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
