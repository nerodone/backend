package auth

import (
	"backend/server"
	"encoding/json"
	"net/http"
)

type UserSignupReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserSignupResp struct {
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	LastLogin    string `json:"last_login"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func signup(_ *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqPayload := &UserSignupReq{}
		if err := json.NewDecoder(r.Body).Decode(reqPayload); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		_, _ = hashPassword(reqPayload.Password)

		if _, err := w.Write([]byte("Auth Signup")); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}
}
