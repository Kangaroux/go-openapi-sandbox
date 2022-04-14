CREATE TABLE IF NOT EXISTS users (
    id          serial PRIMARY KEY,
    created_at  date NOT NULL DEFAULT NOW(),
    updated_at  date NOT NULL,
    email       varchar(255) UNIQUE,
    username    varchar(16) UNIQUE,
    password    varchar(255)
);