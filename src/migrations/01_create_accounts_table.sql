CREATE TABLE IF NOT EXISTS accounts (
    id INT auto_increment PRIMARY KEY,
    user_id INT NOT NULL,
    balance FLOAT NOT NULL,
    created BIGINT UNSIGNED,
    updated BIGINT UNSIGNED,
    active TINYINT DEFAULT 1,
    CONSTRAINT user_unigue UNIQUE (user_id)
);