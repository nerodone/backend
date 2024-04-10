package server

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	secret  = "HelloWorld"
	payload = Payload{
		UserID:   "1",
		Username: "test",
	}
)

var jwt = Init(secret)

func TestJWT(t *testing.T) {
	exp := time.Now().Add(time.Hour * 24).Truncate(time.Second).UTC()

	token := jwt.EncodeToken(payload, false)
	decodedToken := jwt.DecodedToken(token)

	assert.Equal(t, payload, decodedToken.Payload)
	assert.Equal(t, exp, decodedToken.Exp)
}

func TestVerifyToken(t *testing.T) {
	token := jwt.EncodeToken(payload, false)

	assert.True(t, jwt.VerifyToken(token))
	assert.False(t, jwt.VerifyToken(token[:len(token)-1]+"a"))
}
