-- Set Environment Variables
\set userdb `echo "$APP_DB_USER"`
\set passdb `echo "$APP_DB_PASS"`
\set dbname `echo "$APP_DB_NAME"`

-- Create Database
CREATE DATABASE :"dbname";

-- Create User and Grant Privileges
CREATE USER :"userdb" WITH ENCRYPTED PASSWORD :'passdb';
GRANT ALL PRIVILEGES ON DATABASE :"dbname" TO :"userdb";
\connect :"dbname" :"userdb"

-- Create Extension and Install into Database
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create User Table
CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    username VARCHAR(250),
    firstname VARCHAR(250),
    lastname VARCHAR(250),
    email VARCHAR(250),
    phone VARCHAR(250)
);

-- Insert User Data 
INSERT INTO "user" (username, firstname, lastname, email, phone) VALUES 
    ('username1', 'John', 'Doe', 'john@mail.com', '1234567890'),
    ('username2', 'Jane', 'Doe', 'jane@mail.com', '0987654321'),
    ('username3', 'Mary', 'Jane', 'mary@mail.com', ''),
    ('username4', 'John', 'Smith', 'smith@ac.com', '22222222222'),
    ('username5', 'Thomas', 'Anderson', 'ander@ww.com', '1112222333');

