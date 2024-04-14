-- +goose Up
create table collaborate (
    id uuid primary key,
    user_id uuid not null references users (id),
    workspace_id uuid not null references workspaces (id),
    edit boolean not null,
    manage boolean not null,
    created_at timestamp not null
);

create table invite (
    id uuid primary key,
    from_id uuid not null references users (id),
    to_id uuid not null references users (id),
    workspace_id uuid not null references workspaces (id),
    created_at timestamp not null,
    manage boolean not null,
    edit boolean not null
);
-- Table Collaborate{
--   id string
--   user string
--   workspace string
--   edit boolean
--   manage boolean
--   created_at date
-- }
-- Table Invite{
--   id string
--   user string
--   invited string
--   workspace string
--   created_at string
-- }
