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

func signup(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Auth Signup"))
	}
}
