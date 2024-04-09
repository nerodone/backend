-- +goose Up
CREATE TABLE PasswordLogin (
  id uuid PRIMARY KEY,
  user_id uuid REFERENCES Users(id),
  email varchar(127),
  password varchar(255),
  last_login timestamp
);

CREATE TABLE Oauth (
  id uuid PRIMARY KEY,
  user_id uuid REFERENCES Users(id) ON DELETE CASCADE,
  provider varchar(255),
  avatar varchar(255),
  email varchar(255),
  username varchar(255)
  );

CREATE TYPE EPlatform AS ENUM ('neovim', 'web', 'desktop', 'mobile', 'cli','vscode');

CREATE TABLE Sessions (
id uuid PRIMARY KEY,
  user_id uuid REFERENCES Users(id),
  access_token varchar(255),
  refresh_token varchar(255),
  platform EPlatform,
  method varchar(255),
  Oauth_id uuid REFERENCES Oauth(id) ON DELETE CASCADE, --fk
  password_login_id uuid REFERENCES password_login(id) ON DELETE CASCADE,

  created_at timestamp,
  last_login timestamp
);

-- +goose Down
DROP TABLE Sessions;
DROP TABLE Oauth;
DROP TABLE PasswordLogin;
