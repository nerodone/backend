package auth

import (
	"log"
	"os"

	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

var tokenAuth *jwtauth.JWTAuth

type Payload struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

func encode(payload Payload) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	secret := os.Getenv("JWT_SECRET")
	tokenAuth = jwtauth.New("HS256", []byte(secret), nil)

	payloadMap := map[string]interface{}{
		"user_id":  payload.UserID,
		"username": payload.Username,
	}
	_, tokenString, err := tokenAuth.Encode(payloadMap)
	if err != nil {
		log.Fatal("Error encoding token", err)
	}

	return tokenString
}

func decode(token string) (Payload, jwt.Token) {
	decodedToken, err := tokenAuth.Decode(token)

	if err != nil {
		log.Fatal("Error decoding token", err)
	}

	claims := decodedToken.PrivateClaims()

	payload := Payload{
		UserID:   claims["user_id"].(string),
		Username: claims["username"].(string),
	}

	return payload, decodedToken
}
