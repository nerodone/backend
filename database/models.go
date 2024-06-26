// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Eplatform string

const (
	EplatformNeovim  Eplatform = "neovim"
	EplatformWeb     Eplatform = "web"
	EplatformDesktop Eplatform = "desktop"
	EplatformMobile  Eplatform = "mobile"
	EplatformCli     Eplatform = "cli"
	EplatformVscode  Eplatform = "vscode"
)

func (e *Eplatform) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Eplatform(s)
	case string:
		*e = Eplatform(s)
	default:
		return fmt.Errorf("unsupported scan type for Eplatform: %T", src)
	}
	return nil
}

type NullEplatform struct {
	Eplatform Eplatform `json:"eplatform"`
	Valid     bool      `json:"valid"` // Valid is true if Eplatform is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullEplatform) Scan(value interface{}) error {
	if value == nil {
		ns.Eplatform, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Eplatform.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullEplatform) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Eplatform), nil
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Platform     Eplatform `json:"platform"`
	CreatedAt    time.Time `json:"created_at"`
	LastLogin    time.Time `json:"last_login"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	LastLogin time.Time `json:"last_login"`
}
