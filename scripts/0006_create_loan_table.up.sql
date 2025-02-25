CREATE TABLE IF NOT EXISTS loan
(
    id          INT(64) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id     INT(64) UNSIGNED NOT NULL,
    book_id     INT(64) UNSIGNED NOT NULL,
    loan_date   DATE NOT NULL,
    return_date DATE,
    FOREIGN KEY (user_id) REFERENCES user (id),
    FOREIGN KEY (book_id) REFERENCES book (id)
);