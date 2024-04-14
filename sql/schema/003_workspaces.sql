-- +goose Up
CREATE TABLE workspaces (
    id uuid PRIMARY KEY,
    owner uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE ,
    name varchar(255) NOT NULL,
    description text,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

CREATE TABLE projects (
    id uuid PRIMARY KEY,
    name varchar(255) NOT NULL,
    description text,
    workspace uuid NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

CREATE TABLE priority (
    id uuid PRIMARY KEY,
    name varchar(255) NOT NULL,
    workspace uuid NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
    color varchar(255) NOT NULL
);

CREATE TABLE tasks (
    id uuid PRIMARY KEY,
    name varchar(255) NOT NULL,
    project_id uuid NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    status varchar(255) NOT NULL,
    description text,
    due date,
    url varchar(255),
    created_by uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at timestamp NOT NULL,
    priority uuid REFERENCES priority(id) ON DELETE CASCADE ,
    previous uuid REFERENCES tasks(id) ON DELETE CASCADE,
    next uuid REFERENCES tasks(id) ON DELETE CASCADE
);

CREATE TABLE subtasks (
    id uuid PRIMARY KEY,
    name varchar(255) NOT NULL,
    task_id uuid NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
    status varchar(255) NOT NULL,
    description text,
    due date ,
    url varchar(255),
    created_by uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at timestamp NOT NULL,
    priority uuid REFERENCES priority(id) ON DELETE CASCADE,
    previous uuid REFERENCES subtasks(id) ON DELETE CASCADE,
    next uuid REFERENCES subtasks(id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE subtasks;
DROP TABLE tasks;
DROP TABLE priority;
DROP TABLE projects;
DROP TABLE workspaces;
