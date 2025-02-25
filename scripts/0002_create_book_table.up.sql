-- create_book_table.up.sql
CREATE TABLE IF NOT EXISTS book
(
    id           INT(64) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title         VARCHAR(255) NOT NULL,
    author        VARCHAR(255) NOT NULL,
    isbn         VARCHAR(255) NOT NULL,
    is_available BOOLEAN DEFAULT TRUE
);