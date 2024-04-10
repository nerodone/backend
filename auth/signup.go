package auth

import (
	"backend/internal/database"
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

func validateRequest(req *UserSignupRequest) bool {
	return validUsername(req.UserName) && validPlatform(req.Platform) && validPassword(req.Password) && validEmail(req.Email)
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
		if err := json.NewDecoder(r.Body).Decode(reqPayload); err != nil || !validateRequest(reqPayload) {
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

		accessToken := s.JWT.EncodeToken(server.Payload{UserID: user.ID.String(), Username: user.UserName}, false)
		refreshToken := s.JWT.EncodeToken(server.Payload{UserID: user.ID.String(), Username: user.UserName}, true)

		SessionID, err := s.Db.CreateSessionWithPassword(s.Ctx, database.CreateSessionWithPasswordParams{
			UserID:          user.ID,
			Platform:        database.Eplatform(reqPayload.Platform),
			PasswordLoginID: NullableID(passwdID),
			AccessToken:     accessToken,
			RefreshToken:    refreshToken,
		})

		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
			return
		}
		resp := &UserSignupResponse{
			UserName:     user.UserName,
			Email:        user.Email,
			LastLogin:    time.Now().Truncate(time.Second).UTC().String(),
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
		_ = resp
		_ = SessionID

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
			return
		}
	}
}
