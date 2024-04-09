-- +goose Up
CREATE TABLE Users (
   id uuid PRIMARY KEY,
  user_name varchar(255) UNIQUE NOT NULL ,
  email varchar(255) UNIQUE NOT NULL,
  created_at timestamp NOT NULL,
  last_login timestamp NOT NULL
);

-- +goose Down
DROP TABLE Users ;
