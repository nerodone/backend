-- +goose Up
CREATE TABLE PasswordLogin (
  id uuid PRIMARY KEY,
  user_id uuid REFERENCES Users(id),
  email varchar(127) NOT NULL,
  password varchar(255),
  last_login timestamp NOT NULL
);

CREATE TABLE Oauth (
  id uuid PRIMARY KEY,
  user_id uuid REFERENCES Users(id) ON DELETE CASCADE,
  provider varchar(255) NOT NULL,
  avatar varchar(255),
  email varchar(255) NOT NULL,
  username varchar(255) NOT NULL
  );

CREATE TYPE EPlatform AS ENUM ('neovim', 'web', 'desktop', 'mobile', 'cli','vscode');

CREATE TYPE EMethod AS ENUM ('password', 'oauth');
CREATE TABLE Sessions (
id uuid PRIMARY KEY,
  user_id uuid REFERENCES Users(id),
  access_token varchar(255),
  refresh_token varchar(255) NOT NULL,
  platform EPlatform NOT NULL,

  method EMethod NOT NULL,
  Oauth_id uuid REFERENCES Oauth(id) ON DELETE CASCADE,
  password_login_id uuid REFERENCES PasswordLogin(id) ON DELETE CASCADE,
  created_at timestamp NOT NULL,
  last_login timestamp NOT NULL
);

-- +goose Down
DROP TABLE Sessions;
DROP TABLE Oauth;
DROP TABLE PasswordLogin;
DROP TYPE IF EXISTS EPlatform;
DROP TYPE IF EXISTS EMethod;
