-- +goose Up
CREATE TABLE auth (
  session_id uuid PRIMARY KEY,
  refresh_token varchar(255),
  user_id uuid  REFERENCES users(id)
);

-- +goose Down
DROP TABLE auth;
