package auth

import (
	"log"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

var TokenAuth *jwtauth.JWTAuth

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

func Encode(secret string) func(payload Payload) string {
	return func(payload Payload) string {
		_ = jwtauth.ExpireIn(time.Hour * 24)

		TokenAuth = jwtauth.New("HS256", []byte(secret), nil)

		payloadMap := map[string]interface{}{
			"user_id":  payload.UserID,
			"username": payload.Username,
		}

		jwtauth.SetExpiryIn(payloadMap, time.Hour*24)
		_, tokenString, err := TokenAuth.Encode(payloadMap)
		if err != nil {
			log.Fatal("Error encoding token", err)
		}

		return tokenString
	}
}

func Decode(token string) DecodedToken {
	decodedToken, err := TokenAuth.Decode(token)
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

func Verify(token string) bool {
	_, err := jwtauth.VerifyToken(TokenAuth, token)
	return err == nil
}
