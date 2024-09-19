-- Budget App Schema

-- Accounts
CREATE TABLE IF NOT EXISTS account (
    id INTEGER PRIMARY KEY NOT NULL,
    account_name TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp,
    user_id INTEGER NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
);

-- Sessions
CREATE TABLE IF NOT EXISTS sessions(
	token TEXT PRIMARY KEY,
	data BLOB NOT NULL,
	expiry REAL NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions(expiry);

-- Users
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY NOT NULL,
    username TEXT NOT NULL,
    email TEXT NOT NULL, 
    hashed_password TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp,
    CONSTRAINT email_unique UNIQUE (email)
);
