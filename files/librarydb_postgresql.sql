
-- Comando para crear base de datos
-- CREATE DATABASE librarydb;


-- Comando para crear tabla books
DROP TABLE IF EXISTS books;

create table books (
	id SERIAL primary key,
	title varchar(250) not null,
	author varchar(250) not null,
	yearPublication int not null,
	summary varchar(500)
);