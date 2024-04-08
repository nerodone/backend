package main

import "net/http"

type UserSignupReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserSignupResp struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

func handlerUserSignUp(w http.ResponseWriter, r *http.Request) {
}
