-- name: CreateSessionWithPassword :one
insert into Sessions (id , user_id , access_token , refresh_token , platform, created_at, last_login)
values ( gen_random_uuid(), $1 , $2 , $3 , $4 , NOW(), NOW())
returning id;


