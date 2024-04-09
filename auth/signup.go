package auth

import (
	"backend/server"
	"net/http"
)

type UserSignupReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserSignupResp struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

func signup(_ *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Auth Signup"))
		if err != nil {
			return
		}
	}
}
