-- +goose Up
CREATE TYPE EPlatform AS ENUM ('neovim', 'web', 'desktop', 'mobile', 'cli','vscode');
CREATE TABLE Sessions (
id uuid PRIMARY KEY,
  user_id uuid REFERENCES Users(id),
  access_token varchar(255),
  refresh_token varchar(255),
  platform EPlatform,
  method varchar(255),
  Oauth_id uuid, --fk
  password_login_id uuid, --fk

  created_at timestamp,
  last_login timestamp
);

-- +goose Down
DROP TABLE Sessions;
