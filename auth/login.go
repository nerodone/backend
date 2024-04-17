package auth

import (
	"backend/database"
	_ "backend/docs"
	"backend/server"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	ErrInvalidEmail    = "invalid email"
	ErrInvalidPassword = "invalid password"
	ErrInvalidRequest  = "invalid request payload"
)

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Platform string `json:"platform" enums:"web,cli,desktop,neovim,vscode"`
}

func (req *loginReq) validateRequest() bool {
	return validEmail(req.Email)
}

type loginRes struct {
	AccessToken  string    `json:"access_token" example:"token"`
	RefreshToken string    `json:"refresh_token" example:"token"`
	ID           string    `json:"id" example:"61fe5f30-2dfa-4991-97e0-f491027bbc41"`
	UserName     string    `json:"user_name" example:"myUserName"`
	Email        string    `json:"email" example:"user@test.com"`
	CreatedAt    time.Time `json:"created_at"`
	LastLogin    time.Time `json:"last_login"`
	Platform     string    `json:"platform" enums:"web,cli,desktop,neovim,vscode"`
	SessionID    string    `json:"session_id" example:"5d60d0ea-afe4-44d2-8455-cdc624becf07"`
}

// Login
//
//	@Summary	Authenticate user and return access and refresh tokens
//	@Tags		auterrInvalidEmaierrInvalidEmaillh
//	@Accept		json
//	@Produce	json
//	@Param		request	body			loginReq	true	" "
//	@Success	200		{object}		loginRes
//	@failure	401		"invalid email"	||	"invalid password"
//	@failure	400		"invalid request payload"
//	@Failure	500		"Internal Server Error"
//	@Router		/auth/login [post]
func login(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqPayload := &loginReq{}

		if err := json.NewDecoder(r.Body).Decode(reqPayload); err != nil || !reqPayload.validateRequest() {
			s.RespondWithError(w, http.StatusBadRequest, ErrInvalidRequest, "err", err.Error())
			return
		}

		userDB, err := s.Db.LoginUser(s.Ctx, reqPayload.Email)
		if err != nil {
			s.RespondWithError(w, http.StatusUnauthorized, ErrInvalidEmail, "err", err.Error())
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(reqPayload.Password)); err != nil {
			s.RespondWithError(w, http.StatusUnauthorized, ErrInvalidPassword, "err", err.Error())
			return
		}
		refreshToken, err := s.JWT.EncodeToken(server.Payload{UserID: userDB.ID.String(), Username: userDB.UserName}, true)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", "err_encoding_token", err.Error())
			return
		}
		accessToken, err := s.JWT.EncodeToken(server.Payload{UserID: userDB.ID.String(), Username: userDB.UserName}, false)
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", "err_encoding_token", err.Error())
			return
		}
		session, err := s.Db.CreateSessionWithPassword(s.Ctx, database.CreateSessionWithPasswordParams{
			UserID:       userDB.ID,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			Platform:     database.Eplatform(reqPayload.Platform),
		})
		if err != nil {
			s.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", "db_err_creating_session", err.Error())
			return
		}

		responsePayload := loginRes{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ID:           userDB.ID.String(),
			SessionID:    session.ID.String(),
			UserName:     userDB.UserName,
			Email:        userDB.Email,
			CreatedAt:    userDB.CreatedAt,
			LastLogin:    userDB.LastLogin,
			Platform:     reqPayload.Platform,
		}
		s.RespondWithJson(w, http.StatusOK, responsePayload)
	}
}
