ALTER TABLE users
ADD username VARCHAR(255) NOT NULL;

ALTER TABLE users
ADD CONSTRAINT UNIQUE unique_username (username);