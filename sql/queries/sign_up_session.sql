-- name: CreateSessionWithPassword :one
insert into Sessions (id , user_id , access_token , refresh_token , platform, method,Oauth_id, password_login_id, created_at, last_login)
values ( gen_random_uuid(), $1 , $2 , $3 , $4 ,'password', NULL, $5, NOW(), NOW())
returning id;

-- name: CreateSessionWithOauth :one
insert into Sessions (id , user_id , access_token , refresh_token , platform, method,Oauth_id, password_login_id, created_at, last_login)
values ( gen_random_uuid(), $1 , $2 , $3 , $4 ,'oauth', $5, NULL, NOW(), NOW())
returning id;

