CREATE TABLE IF NOT EXISTS role (
    role_id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
    );

INSERT INTO role (name) VALUES ('admin');
INSERT INTO role (name) VALUES ('receptionist');
INSERT INTO role (name) VALUES ('cashier');

CREATE TABLE IF NOT EXISTS "user" (
    user_id SERIAL PRIMARY KEY,
    role_id INT NOT NULL REFERENCES role(role_id) ON DELETE CASCADE,
    username VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE
    );

CREATE TABLE IF NOT EXISTS client_category (
    client_category_id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE
    );

INSERT INTO client_category (name) VALUES ('students');
INSERT INTO client_category (name) VALUES ('faculty');
INSERT INTO client_category (name) VALUES ('staff');

CREATE TABLE IF NOT EXISTS client (
    client_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(10) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    client_category_id INT NOT NULL REFERENCES client_category(client_category_id) ON DELETE CASCADE,
    balance FLOAT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE
    );