-- name: CreateUser :one
insert into Users (id , user_name , email , created_at , last_login)
values ( gen_random_uuid(), $1 , $2 , NOW() , NOW() )
    returning *;

