package auth

import (
	_ "backend/docs"
	"backend/server"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type logoutReq struct {
	SessionID uuid.UUID `json:"session_id" example:"f5b1c9e2-1b2c-4b5c-8f1d-8f5b1c9e2f1d"`
}

// logout
//
//	@Summary	delete session : revoke access token refresh token
//	@Accept		json
//	@Produce	json
//	@Param		request	body	logoutReq	true	" "
//	@Success	200
//	@failure	400	"invalid request payload"
//	@Failure	500	"internal Server Error"
//	@Router		/auth/logout [post]
func logout(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqPayload := &logoutReq{}

		if err := json.NewDecoder(r.Body).Decode(reqPayload); err != nil {
			s.RespondWithError(w, http.StatusBadRequest, "invalid request payload", "error_decoding_request_payload", err.Error())
		}

		if err := s.Db.LogoutDeleteSession(s.Ctx, reqPayload.SessionID); err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "internal server error", "error_deleting_session", err.Error())
		}

		w.WriteHeader(http.StatusOK)
	}
}
