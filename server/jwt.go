package server

import (
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

func (jwt JwtProvider) EncodeToken(payload Payload, isRefreshToken bool) (string, error) {
	payloadMap := map[string]interface{}{
		"user_id":  payload.UserID,
		"username": payload.Username,
	}

	var duration time.Duration

	if isRefreshToken {
		duration = time.Hour * 24 * 90
	} else {
		duration = time.Hour * 24
	}

	jwtauth.SetExpiryIn(payloadMap, duration)
	_, tokenString, err := jwt.TokenAuth.Encode(payloadMap)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jwt JwtProvider) DecodedToken(token string) (DecodedToken, error) {
	decodedToken, err := jwt.TokenAuth.Decode(token)
	if err != nil {
		return DecodedToken{}, err
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
	}, nil
}

func (jwt JwtProvider) VerifyToken(token string) bool {
	_, err := jwtauth.VerifyToken(jwt.TokenAuth, token)
	return err == nil
}
