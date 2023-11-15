-- Create a new database
CREATE DATABASE IF NOT EXISTS gouserdb;

-- Use the newly created database
USE gouserdb;

-- Create Addresses table
CREATE TABLE IF NOT EXISTS addresses (
    id INT PRIMARY KEY AUTO_INCREMENT,
    street VARCHAR(80) NOT NULL,
    number INTEGER NOT NULL,
    complement VARCHAR(80),
    postal_code VARCHAR(80),
    city VARCHAR(80) NOT NULL,
    state VARCHAR(80) NOT NULL,
    country VARCHAR(80) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Users table
CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(80) NOT NULL,
    age INTEGER NOT NULL,
    email VARCHAR(120) NOT NULL UNIQUE,
    password CHAR(64) NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
    address_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    CONSTRAINT fk_addresses 
        FOREIGN KEY (address_id) 
            REFERENCES addresses(id)
);
