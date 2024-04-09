package auth

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

func TestJWT(t *testing.T) {
	exp := time.Now().Add(time.Hour * 24).Truncate(time.Second).UTC()

	token := Encode(secret)(payload)
	decodedToken := Decode(token)

	assert.Equal(t, payload, decodedToken.Payload)
	assert.Equal(t, exp, decodedToken.Exp)
}

func TestVerifyToken(t *testing.T) {
	token := Encode(secret)(payload)

	assert.True(t, Verify(token))
	assert.False(t, Verify(token[:len(token)-1]+"a"))
}
