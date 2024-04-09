package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var fakeSecret = "HelloWorld"
var fakePayload = Payload{
	UserID:   "1",
	Username: "test",
}
var fakeToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMSIsInVzZXJuYW1lIjoidGVzdCJ9.n-R354fEZAn0_SqllaJZ4IEtQd_HI3XUYa8BOw2OEog"

func TestEncode(t *testing.T) {
	encoder := Encode(fakeSecret)
	token := encoder(fakePayload)
	assert.Equal(t, fakeToken, token)
}

func TestDecode(t *testing.T) {
	payload, _ := Decode(fakeToken)
	assert.Equal(t, fakePayload, payload)
}
