package auth

import (
	"backend/database"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func NullableID(id uuid.UUID) uuid.NullUUID {
	return uuid.NullUUID{UUID: id, Valid: true}
}

// TODO:
func validPassword(password string) bool {
	return len(password) >= 8
}

// TODO:
func validEmail(email string) bool {
	return true
}

// TODO:
func validUsername(userName string) bool {
	return true
}

func validPlatform(platform string) bool {
	platforms := []database.Eplatform{"neovim", "web", "desktop", "mobile", "cli", "vscode"}
	for _, p := range platforms {
		if string(p) == platform {
			return true
		}
	}
	return false
}

var (
	ErrDuplicateEmail    = errors.New("email already exists")
	ErrDuplicateUserName = errors.New("username alerady exists")
	ErrDatabaseIssue     = errors.New("databaae issue")
)

type SignupError struct {
	originalErr error
	responseErr error
}

func matchErr(err error) SignupError {
	signupErr := SignupError{originalErr: err}
	switch err.Error() {
	case "pq: duplicate key value violates unique constraint \"users_email_key\"":
		signupErr.responseErr = ErrDuplicateEmail

	case "pq: duplicate key value violates unique constraint \"users_user_name_key\"":
		signupErr.responseErr = ErrDuplicateUserName
	default:
		signupErr.responseErr = ErrDatabaseIssue
	}
	return signupErr
}
