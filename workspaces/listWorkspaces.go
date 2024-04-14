package workspaces

import (
	"backend/database"
	"backend/server"
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
)

type listWorkspacesResponse struct {
	Workspaces []database.Workspace `json:"workspaces"`
}

func listWorkspaces(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Error %v", err)
			return
		}

		userID, err := uuid.Parse(claims["user_id"].(string))
		if err != nil {
			s.RespondWithError(w, http.StatusBadRequest, "invalid request", "err", err.Error())
			return
		}

		workspaces, err := s.Db.GetAllWorkspaces(r.Context(), userID)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "internal server error", "err", err.Error())
			return
		}
		s.RespondWithJson(w, http.StatusOK, workspaces)
	}
}
