ALTER TABLE client
    ALTER COLUMN clientID SET DEFAULT gen_random_uuid(),
    ALTER COLUMN password SET NOT NULL,
    ALTER COLUMN login SET NOT NULL,
    ALTER COLUMN email SET NOT NULL,
    ADD PRIMARY KEY (clientID);

ALTER TABLE client
    ADD CHECK (password != ''),
    ADD CHECK (login != ''),
    ADD CHECK (email != '');

ALTER TABLE container
    ALTER COLUMN containerID SET DEFAULT gen_random_uuid(),
    ALTER COLUMN ipcontainer SET NOT NULL;
   
ALTER TABLE container
    ADD PRIMARY KEY (containerID);
   
ALTER TABLE clientcontainer 
 	ADD PRIMARY KEY (clientID, containerID),
    ADD FOREIGN KEY (clientID) REFERENCES client (clientID) ON DELETE cascade,
    ADD FOREIGN KEY (containerID) REFERENCES container (containerID) ON DELETE cascade;
   
ALTER TABLE historycontainer
    ALTER COLUMN historyID SET DEFAULT gen_random_uuid(),
    ALTER COLUMN timeping SET NOT NULL,
    ALTER COLUMN statusping SET NOT null;
   
ALTER TABLE historycontainer
    ADD CHECK (statusping != ''),
    ADD FOREIGN KEY (containerID) REFERENCES container (containerID) ON DELETE cascade,
    ADD PRIMARY KEY (historyID);
