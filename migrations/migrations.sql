-- migrations/migrations.sql
CREATE DATABASE IF NOT EXISTS pos_system;
USE pos_system;

-- Tabel Produk
CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    category VARCHAR(100),
    stock_quantity INT NOT NULL
);

-- Tabel Pengguna
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL
);

-- Tabel Transaksi
CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    total_amount DECIMAL(10, 2),
    payment_method VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
