package auth

import (
	"backend/internal/database"
	"backend/server"
	"encoding/json"
	"net/http"
)

type UserSignupRequest struct {
	UserName string `json:"user_name"`
	Platform string `json:"platform"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserSignupResponse struct {
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	LastLogin    string `json:"last_login"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func signup(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqPayload := &UserSignupRequest{}
		if err := json.NewDecoder(r.Body).Decode(reqPayload); err != nil {
			s.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		hashedPass, err := hashPassword(reqPayload.Password)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
			return
		}

		user, err := s.Db.CreateUser(s.Ctx, database.CreateUserParams{UserName: reqPayload.UserName, Email: reqPayload.Email})
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", "err", err.Error())
			return
		}

		passwdID, err := s.Db.SignupWithPassword(s.Ctx, database.SignupWithPasswordParams{
			UserID: user.ID, Email: user.Email, Password: hashedPass,
		})
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", "err", err.Error())
			return
		}
		// TODO: access/refresh_tokens
		SessionID, err := s.Db.CreateSessionWithPassword(s.Ctx, database.CreateSessionWithPasswordParams{
			UserID:          user.ID,
			Platform:        database.Eplatform(reqPayload.Platform),
			PasswordLoginID: NullableID(passwdID),
		})
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
			return
		}
		resp := &UserSignupResponse{}
		_ = resp
		_ = SessionID
	}
}
