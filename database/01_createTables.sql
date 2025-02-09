CREATE TABLE IF NOT EXISTS client (
    clientID UUID,
    password varchar(255),
    login varchar(255),
    email varchar(255)
);

CREATE TABLE IF NOT EXISTS clientcontainer (
    clientID UUID,
    containerID UUID
);

CREATE TABLE IF NOT EXISTS container (
    containerID UUID,
    ipcontainer INET
);

CREATE TABLE IF NOT EXISTS historycontainer (
    historyID UUID,
    containerID UUID,
   	timeping TIMESTAMP, 
   	statusping varchar(255)
);

