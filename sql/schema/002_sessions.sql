-- +goose Up
CREATE TYPE EPlatform AS ENUM ('neovim', 'web', 'desktop', 'mobile', 'cli','vscode');
CREATE TABLE Sessions (
   id uuid PRIMARY KEY,
  user_id uuid NOT NULL REFERENCES Users(id) ON DELETE CASCADE NOT NULL,
  access_token varchar(255)NOT NULL,
  refresh_token varchar(255) NOT NULL,
  platform EPlatform NOT NULL,
  created_at timestamp NOT NULL,
  last_login timestamp NOT NULL
);
-- +goose Down
DROP TABLE Sessions;
DROP TYPE IF EXISTS EPlatform;
