-- +goose Up
-- Add user_id column to client table with foreign key constraint
ALTER TABLE client
ADD COLUMN user_id UUID;

ALTER TABLE client
ADD CONSTRAINT fk_client_user FOREIGN KEY(user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL;

-- +goose Down
-- Remove the foreign key constraint and user_id column
ALTER TABLE client
DROP CONSTRAINT IF EXISTS fk_client_user;

ALTER TABLE client
DROP COLUMN IF EXISTS user_id;
