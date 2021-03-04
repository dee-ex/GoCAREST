CREATE TABLE IF NOT EXISTS users (
    id          INT PRIMARY KEY AUTO_INCREMENT,
    username    VARCHAR(255) NOT NULL,
    email       VARCHAR(255) NOT NULL,
    password    VARCHAR(255) NOT NULL,
    token       VARCHAR(2048) NOT NULL,
    created_at  DATETIME NOT NULL DEFAULT NOW()
);