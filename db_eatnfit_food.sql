CREATE DATABASE db_eatnfit_food;
USE db_eatnfit_food;

CREATE TABLE tb_food (
    food_id VARCHAR(36) PRIMARY KEY NOT NULL,
    food_portion INT NULL DEFAULT 0,
    food_name VARCHAR(100) NOT NULL,
    food_calories FLOAT NULL DEFAULT 0,
    food_fat FLOAT NULL DEFAULT 0,
    food_carbs FLOAT NULL DEFAULT 0,
    food_protein FLOAT NULL DEFAULT 0,
    food_price INT NOT NULL,
    food_desc TEXT NULL,
    food_status INT NOT NULL
);

CREATE TABLE tb_packet (
    packet_id VARCHAR(36) PRIMARY KEY NOT NULL,
    packet_name VARCHAR(100) NOT NULL,
    packet_price INT NOT NULL,
    packet_desc TEXT NOT NULL,
    packet_status INT NOT NULL
);

CREATE TABLE tb_packet_and_food (
    pm_id VARCHAR(36) PRIMARY KEY NOT NULL,
    packet_id VARCHAR(36) NOT NULL,
    food_id VARCHAR(36) NOT NULL,
    pm_status INT NOT NULL
);

CREATE TABLE tb_transaction (
    trans_id VARCHAR(36) PRIMARY KEY NOT NULL,
    trans_date DATE NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    packet_id VARCHAR(36) NOT NULL,
    portion INT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    start_time TIMESTAMP NOT NULL,
    address TEXT NOT NULL,
    payment_id VARCHAR(36) NOT NULL,
    transaction_status INT NOT NULL DEFAULT 1
);

CREATE TABLE tb_transaction_status (
    status_id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    status_name VARCHAR(100) NOT NULL
);

CREATE TABLE tb_payment (
    payment_id VARCHAR(36) PRIMARY KEY NOT NULL,
    payment_name VARCHAR(100) NOT NULL,
    payment_status INT NOT NULL
);
