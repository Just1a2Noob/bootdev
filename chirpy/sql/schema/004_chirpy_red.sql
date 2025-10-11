-- +goose Up
ALTER TABLE users 
ADD COLUMN is_chirpy_red TEXT NOT NULL DEFAULT 'false';

-- +goose Down
ALTER TABLE users
DROP COLUMN IF EXITS is_chirpy_red;
