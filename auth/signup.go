package auth

import (
	"backend/database"
	"backend/server"
	"encoding/json"
	"net/http"
	"time"
)

type UserSignupRequest struct {
	UserName string `json:"user_name"`
	Platform string `json:"platform"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (req *UserSignupRequest) validateSignupRequest() bool {
	return validUsername(req.UserName) && validPlatform(req.Platform) && validPassword(req.Password) && validEmail(req.Email)
}

type UserSignupResponse struct {
	ID           string    `json:"id"`
	UserName     string    `json:"user_name"`
	Email        string    `json:"email"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	SessionID    string    `json:"session_id"`
	LastLogin    time.Time `json:"last_login"`
	CreatedAt    time.Time `json:"created_at"`
}

func signup(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqPayload := &UserSignupRequest{}
		if err := json.NewDecoder(r.Body).Decode(reqPayload); err != nil || !reqPayload.validateSignupRequest() {
			s.RespondWithError(w, http.StatusBadRequest, "Invalid request payload", "err", err.Error())
			return
		}

		hashedPass, err := hashPassword(reqPayload.Password)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", "err", err.Error())
			return
		}

		user, err := s.Db.CreateUser(s.Ctx, database.CreateUserParams{
			UserName: reqPayload.UserName,
			Email:    reqPayload.Email,
			Password: hashedPass,
		})
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", "err", "hihk")
			return
		}

		accessToken, err := s.JWT.EncodeToken(server.Payload{UserID: user.ID.String(), Username: user.UserName}, false)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", "err", err.Error())
			return
		}

		refreshToken, err := s.JWT.EncodeToken(server.Payload{UserID: user.ID.String(), Username: user.UserName}, true)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", "err", err.Error())
			return
		}

		Session, err := s.Db.CreateSessionWithPassword(s.Ctx, database.CreateSessionWithPasswordParams{
			UserID:       user.ID,
			Platform:     database.Eplatform(reqPayload.Platform),
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", "err", err.Error())
			return
		}
		resp := &UserSignupResponse{
			ID:           user.ID.String(),
			UserName:     user.UserName,
			Email:        user.Email,
			LastLogin:    time.Now().Truncate(time.Second).UTC(),
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			SessionID:    Session.ID.String(),
			CreatedAt:    Session.CreatedAt.Truncate(time.Second).UTC(),
		}
		s.ResponsWithJson(w, http.StatusCreated, resp)
	}
}
