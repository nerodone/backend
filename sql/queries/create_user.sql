-- name: CreateUser :one
insert into Users (id , user_name , email, password, created_at, last_login)
values ( gen_random_uuid(), $1 , $2 , $3 ,NOW() , NOW() )
    returning *;

