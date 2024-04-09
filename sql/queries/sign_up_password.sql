-- name: SignupWithPassword :exec
insert into PasswordLogin (id , user_id , email , password , last_login)
values ( gen_random_uuid(), $1 , $2 , $3 , NOW() );


