package auth

import (
	"backend/database"
	"backend/server"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type UserRefreshRequest struct {
	SessionId    string `json:"session_id"`
	RefreshToken string `json:"refresh_token"`
}

type UserRefreshResponse struct {
	AccessToken string `json:"access_token"`
}

func validateRefreshRequest(s *server.Server, req *UserRefreshRequest) bool {
	refreshToken := req.RefreshToken
	return s.JWT.VerifyToken(refreshToken)
}

func refresh(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestPayload := &UserRefreshRequest{}
		err := json.NewDecoder(r.Body).Decode(requestPayload)

		if err != nil || !validateRefreshRequest(s, requestPayload) {
			s.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		payload, err := s.JWT.DecodedToken(requestPayload.RefreshToken)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Error Decoding Token", err.Error())
		}

		sessionId, err := uuid.Parse(requestPayload.SessionId)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Error Parsing sessionId", err.Error())
		}
		session, err := s.Db.GetSessionById(s.Ctx, sessionId)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Error Getting Session", err.Error())
		}

		newPayload := server.Payload{
			UserID:   payload.Payload.UserID,
			Username: payload.Payload.Username,
		}
		accessToken, err := s.JWT.EncodeToken(newPayload, false)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Error Encoding Access Token", err.Error())
		}

		_, err = s.Db.UpdateAccessToken(s.Ctx, database.UpdateAccessTokenParams{ID: session.ID, AccessToken: accessToken})
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Error updating access token", err.Error())
		}

		resp := &UserRefreshResponse{
			AccessToken: accessToken,
		}

		s.RespondWithJson(w, http.StatusCreated, resp)
	}
}
