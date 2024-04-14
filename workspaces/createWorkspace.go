package workspaces

import (
	"backend/database"
	"backend/server"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
)

type createWorkspaceReq struct {
	Name        string `json:"workspace_name"`
	Description string `json:"description"`
}

func createWorkspace(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := createWorkspaceReq{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.RespondWithError(w, http.StatusBadRequest, "invalid request", "err", err.Error())
			return
		}
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

		workspaceParams := database.CreateWorkspaceParams{
			Name:        req.Name,
			Owner:       userID,
			Description: sql.NullString{Valid: req.Description != "", String: req.Description},
		}

		workspace, err := s.Db.CreateWorkspace(r.Context(), workspaceParams)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "internal server error", "err", err.Error())
			return
		}
		s.RespondWithJson(w, http.StatusCreated, workspace)
	}
}