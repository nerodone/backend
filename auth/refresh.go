package auth

import (
	"backend/database"
	_ "backend/docs"
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

// refresh
//
//	@Summary	send refresh token to get  new access token
//	@Accept		json
//	@Produce	json
//	@Param		request	body		UserRefreshRequest	true	" "
//	@Success	201		{object}	UserRefreshResponse
//	@Failure	500		"Internal Server Error"
//	@Router		/auth/refresh [post]
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
			return
		}

		sessionId, err := uuid.Parse(requestPayload.SessionId)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Error Parsing sessionId", err.Error())
			return
		}
		session, err := s.Db.GetSessionById(s.Ctx, sessionId)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Error Getting Session", err.Error())
			return
		}

		newPayload := server.Payload{
			UserID:   payload.Payload.UserID,
			Username: payload.Payload.Username,
		}
		accessToken, err := s.JWT.EncodeToken(newPayload, false)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Error Encoding Access Token", err.Error())
			return
		}

		_, err = s.Db.UpdateAccessToken(s.Ctx, database.UpdateAccessTokenParams{ID: session.ID, AccessToken: accessToken})
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Error updating access token", err.Error())
			return
		}

		resp := &UserRefreshResponse{
			AccessToken: accessToken,
		}

		s.RespondWithJson(w, http.StatusCreated, resp)
	}
}
