package auth

import (
	"backend/internal/database"
	"backend/server"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type loginReq struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	SessionId uuid.UUID `json:"session_id"`
}

func (req *loginReq) validateRequest() bool {
	return validEmail(req.Email)
}

type loginRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func login(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqPayload := &UserSignupRequest{}

		if err := json.NewDecoder(r.Body).Decode(reqPayload); err != nil || !reqPayload.validateSignupRequest() {
			s.RespondWithError(w, http.StatusBadRequest, "Invalid request payload", "err", err.Error())
			return
		}

		hashedPassword, err := hashPassword(reqPayload.Password)
		if err != nil {
			s.RespondWithError(w, http.StatusBadRequest, "password too long", "err", err.Error())
			return
		}

		passwordID, err := s.Db.LoginUser(s.Ctx, database.LoginUserParams{
			Password: hashedPassword,
			Email:    reqPayload.Email,
		})
		if err != nil {
			s.RespondWithError(w, http.StatusUnauthorized, "Invalid email or password", "err", err.Error())
			return
		}
		_ = passwordID
	}
}
