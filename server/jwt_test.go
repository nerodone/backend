package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var secret = "HelloWorld"
var payload = Payload{
	UserID:   "1",
	Username: "test",
}

var jwt = Init(secret)

func TestJWT(t *testing.T) {
	exp := time.Now().Add(time.Hour * 24).Truncate(time.Second).UTC()

	token, err := jwt.EncodeToken(payload, false)
	if err != nil {
		t.Errorf("Error Encoding Token: %v", err)
	}
	decodedToken, err := jwt.DecodedToken(token)
	if err != nil {
		t.Errorf("Error Decoding Token: %v", err)
	}

	assert.Equal(t, payload, decodedToken.Payload)
	assert.Equal(t, exp, decodedToken.Exp)
}

func TestVerifyToken(t *testing.T) {
	token, err := jwt.EncodeToken(payload, false)
	if err != nil {
		t.Errorf("Error Encoding Token : %v", err)
	}

	assert.True(t, jwt.VerifyToken(token))
	assert.False(t, jwt.VerifyToken(token[:len(token)-1]+"a"))
}
