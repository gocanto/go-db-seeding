-- schema_init.up.sql

BEGIN;

-- Create the users table
CREATE TABLE users
(
    id       SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    handle   VARCHAR(255) NOT NULL
);

COMMIT;
