-- +goose Up
-- Add password column to client table
ALTER TABLE client
ADD COLUMN password TEXT;

-- +goose Down
-- Remove the password column
ALTER TABLE client
DROP COLUMN IF EXISTS password;
