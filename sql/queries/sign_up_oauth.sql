-- name: SignupWithOauth :one
insert into Oauth (id , user_id , provider , avatar , email, username)
values ( gen_random_uuid(), $1 , $2 , $3 , $4, $5)
returning id, username, avatar ;

