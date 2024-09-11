CREATE TABLE IF NOT EXISTS account (
    account_name TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT current_timestamp,
    updated_at TEXT NOT NULL DEFAULT current_timestamp
);
