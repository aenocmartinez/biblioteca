-- script para crear base de datos librarydb 

drop database if exists librarydb;

create database librarydb;

use librarydb;

create table books (
	id bigint auto_increment primary key,
	title varchar(250) not null,
	author varchar(250) not null,
	yearPublication int not null,
	summary varchar(500)
);