package workspaces

import (
	"backend/database"
	"backend/server"
	"fmt"
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
//	@Success		200	{array}	database.Workspace
//	@failure		400	"invalid request payload"
//	@Failure		500	"internal Server Error"
//	@Router			/workspaces [get]
func listWorkspaces(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Error %v", err)
			return
		}

		userID, err := uuid.Parse(claims["user_id"].(string))
		if err != nil {
			s.RespondWithError(w, http.StatusBadRequest, "invalid request payload", "err", err.Error())
			return
		}

		workspaces, err := s.Db.GetAllWorkspaces(r.Context(), userID)
		if workspaces == nil {
			workspaces = []database.Workspace{}
		}
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "internal server error", "err", err.Error())
			return
		}
		s.RespondWithJson(w, http.StatusOK, workspaces)
	}
}
