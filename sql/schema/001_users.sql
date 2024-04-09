-- +goose Up
CREATE TABLE Users (
   id uuid PRIMARY KEY,
  user_name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  created_at timestamp NOT NULL,
  last_login timestamp NOT NULL
);

-- +goose Down
DROP TABLE users ;
