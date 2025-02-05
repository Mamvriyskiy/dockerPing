-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL limitations';

ALTER TABLE client
    ALTER COLUMN password SET NOT NULL,
    ALTER COLUMN login SET NOT NULL,
    ALTER COLUMN email SET NOT NULL;

ALTER TABLE client
    ADD CHECK (password != ''),
    ADD CHECK (login != ''),
    ADD CHECK (email != ''),
    ADD PRIMARY KEY (clientID);

ALTER TABLE container
    ALTER COLUMN ipcontainer SET NOT NULL,
    ALTER COLUMN namecontainer SET NOT null;
   
ALTER TABLE container
    ADD CHECK (namecontainer != ''),
    ADD PRIMARY KEY (containerID);
   
ALTER TABLE clientcontainer 
 	ADD PRIMARY KEY (clientID, containerID),
    ADD FOREIGN KEY (clientID) REFERENCES client (clientID) ON DELETE cascade,
    ADD FOREIGN KEY (containerID) REFERENCES container (containerID) ON DELETE cascade;
   
ALTER TABLE historycontainer
    ALTER COLUMN timeping SET NOT NULL,
    ALTER COLUMN statusping SET NOT null;
   
ALTER TABLE historycontainer
    ADD CHECK (statusping != ''),
    ADD FOREIGN KEY (containerID) REFERENCES container (containerID) ON DELETE cascade,
    ADD PRIMARY KEY (historyID);
   

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
