CREATE TABLE IF NOT EXISTS users (
    username TEXT NOT NULL,
    email TEXT NOT NULL, 
    hashed_password TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp,
    CONSTRAINT email_unique UNIQUE (email)
);
