-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL tables';

CREATE TABLE IF NOT EXISTS client (
    clientID UUID,
    password varchar(255),
    login varchar(255),
    email varchar(255)
);

CREATE TABLE IF NOT EXISTS clientcontainer (
    ID UUID,
    clientID UUID,
    containerID UUID
);

CREATE TABLE IF NOT EXISTS container (
    containerID UUID,
    ipcontainer INET,
    namecontainer varchar(255)
);

CREATE TABLE IF NOT EXISTS historycontainer (
    historyID UUID,
    containerID UUID,
   	timeping TIMESTAMPTZ, 
   	statusping varchar(255)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE clientcontainer, client, historycontainer, container;
-- +goose StatementEnd
