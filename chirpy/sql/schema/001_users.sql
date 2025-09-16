-- +goose Up
CREATE TABLE users(
    id uuid primary key,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    email TEXT NOT NULL UNIQUE
);

ALTER TABLE users 
ADD COLUMN hashed_password TEXT NOT NULL DEFAULT 'unset';

-- +goose Down
DROP TABLE users;
