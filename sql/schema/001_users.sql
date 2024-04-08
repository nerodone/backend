-- +goose Up
CREATE TABLE users (
  id uuid PRIMARY KEY,
  created_at timestamp NOT NULL,
  user_name varchar(255) UNIQUE NOT NULL ,
  api_key varchar(64) UNIQUE NOT NULL DEFAULT encode(sha256(random()::text::bytea), 'hex')
);

-- +goose Down
DROP TABLE users ;
