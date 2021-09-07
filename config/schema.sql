--create user jumpcloud;
--alter user jumpcloud with encrypted password 'password';
--DROP DATABASE IF EXISTS jumpcloud;
--CREATE DATABASE jumpcloud;
--grant all privileges on database jumpcloud to jumpcloud;

CREATE TABLE IF NOT EXISTS hashes (id serial, encodedval varchar(256));
