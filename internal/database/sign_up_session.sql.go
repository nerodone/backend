// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: sign_up_session.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createSessionWithOauth = `-- name: CreateSessionWithOauth :exec
insert into Sessions (id , user_id , access_token , refresh_token , platform, method,Oauth_id, password_login_id, created_at, last_login)
values ( gen_random_uuid(), $1 , $2 , $3 , $4 ,'oauth', $5, NULL, NOW(), NOW())
`

type CreateSessionWithOauthParams struct {
	UserID       uuid.NullUUID  `json:"user_id"`
	AccessToken  sql.NullString `json:"access_token"`
	RefreshToken string         `json:"refresh_token"`
	Platform     Eplatform      `json:"platform"`
	OauthID      uuid.NullUUID  `json:"oauth_id"`
}

func (q *Queries) CreateSessionWithOauth(ctx context.Context, arg CreateSessionWithOauthParams) error {
	_, err := q.db.ExecContext(ctx, createSessionWithOauth,
		arg.UserID,
		arg.AccessToken,
		arg.RefreshToken,
		arg.Platform,
		arg.OauthID,
	)
	return err
}

const createSessionWithPassword = `-- name: CreateSessionWithPassword :exec
insert into Sessions (id , user_id , access_token , refresh_token , platform, method,Oauth_id, password_login_id, created_at, last_login)
values ( gen_random_uuid(), $1 , $2 , $3 , $4 ,'password', NULL, $5, NOW(), NOW())
`

type CreateSessionWithPasswordParams struct {
	UserID          uuid.NullUUID  `json:"user_id"`
	AccessToken     sql.NullString `json:"access_token"`
	RefreshToken    string         `json:"refresh_token"`
	Platform        Eplatform      `json:"platform"`
	PasswordLoginID uuid.NullUUID  `json:"password_login_id"`
}

func (q *Queries) CreateSessionWithPassword(ctx context.Context, arg CreateSessionWithPasswordParams) error {
	_, err := q.db.ExecContext(ctx, createSessionWithPassword,
		arg.UserID,
		arg.AccessToken,
		arg.RefreshToken,
		arg.Platform,
		arg.PasswordLoginID,
	)
	return err
}
