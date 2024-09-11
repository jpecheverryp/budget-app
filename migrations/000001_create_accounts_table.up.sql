CREATE TABLE IF NOT EXISTS account (
    account_name TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp
);
