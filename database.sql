-- Active: 1700143600758@@127.0.0.1@3306@ddd

show DATABASES;

CREATE TABLE
    customers(
        id VARCHAR(100) PRIMARY KEY UNIQUE NOT NULL,
        name VARCHAR(100) NOT NULL
    ) engine = InnoDB;

DESC customers;