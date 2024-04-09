package server

import (
	"log"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

type Payload struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type DecodedToken struct {
	Payload Payload
	Aud     []string
	Exp     time.Time
	Issuer  string
	Jti     string
	Nbf     time.Time
	Sub     string
}

type JwtProvider struct {
	TokenAuth *jwtauth.JWTAuth
}

func Init(secret string) JwtProvider {
	_ = jwtauth.ExpireIn(time.Hour * 24)
	tokenAuth := jwtauth.New("HS256", []byte(secret), nil)
	return JwtProvider{TokenAuth: tokenAuth}
}

func (jwt JwtProvider) EncodeToken(payload Payload) string {
	payloadMap := map[string]interface{}{
		"user_id":  payload.UserID,
		"username": payload.Username,
	}

	jwtauth.SetExpiryIn(payloadMap, time.Hour*24)
	_, tokenString, err := jwt.TokenAuth.Encode(payloadMap)
	if err != nil {
		log.Fatal("Error encoding token", err)
	}

	return tokenString
}

func (jwt JwtProvider) DecodedToken(token string) DecodedToken {
	decodedToken, err := jwt.TokenAuth.Decode(token)
	if err != nil {
		log.Fatal("Error decoding token", err)
	}

	claims := decodedToken.PrivateClaims()

	payload := Payload{
		UserID:   claims["user_id"].(string),
		Username: claims["username"].(string),
	}

	return DecodedToken{
		Payload: payload,
		Aud:     decodedToken.Audience(),
		Exp:     decodedToken.Expiration(),
		Issuer:  decodedToken.Issuer(),
		Jti:     decodedToken.JwtID(),
		Nbf:     decodedToken.NotBefore(),
		Sub:     decodedToken.Subject(),
	}
}

func (jwt JwtProvider) VerifyToken(token string) bool {
	_, err := jwtauth.VerifyToken(jwt.TokenAuth, token)
	return err == nil
}
