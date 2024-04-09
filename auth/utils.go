package auth

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func responsWithJson(w http.ResponseWriter, stausCode int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(stausCode)
	_, err = w.Write(response)
	if err != nil {
		return err
	}
	return nil
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	_ = responsWithJson(w, status, map[string]string{"error": message})
}
