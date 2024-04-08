-- +goose Up
CREATE TABLE Users (
   id uuid PRIMARY KEY,
  user_name varchar(255),
  email varchar(255),
  created_at timestamp,
  last_login timestamp
);

-- +goose Down
DROP TABLE users ;
