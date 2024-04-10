// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: login_user.sql.sql

package database

import (
	"context"
)

const loginUser = `-- name: LoginUser :one
SELECT id, user_name, email, password, created_at, last_login FROM users WHERE email = $1 AND password = $2 LIMIT 1
`

type LoginUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) LoginUser(ctx context.Context, arg LoginUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, loginUser, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.LastLogin,
	)
	return i, err
}
