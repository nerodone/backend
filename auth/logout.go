package auth

import (
	"backend/server"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type logoutReq struct {
	SessionID uuid.UUID `json:"session_id"`
}

// type logoutRes strcut { }

func logout(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqPayload := &logoutReq{}

		if err := json.NewDecoder(r.Body).Decode(reqPayload); err != nil {
			s.RespondWithError(w, http.StatusBadRequest, "internal server error", "error_decoding_request_payload", err.Error())
		}

		if err := s.Db.LogoutDeleteSession(s.Ctx, reqPayload.SessionID); err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "internal server error", "error_deleting_session", err.Error())
		}

		w.WriteHeader(http.StatusOK)
	}
}
