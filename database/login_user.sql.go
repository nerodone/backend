// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: login_user.sql

package database

import (
	"context"
)

const loginUser = `-- name: LoginUser :one
SELECT id, user_name, email, password, created_at, last_login FROM users WHERE email = $1 LIMIT 1
`

func (q *Queries) LoginUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, loginUser, email)
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
