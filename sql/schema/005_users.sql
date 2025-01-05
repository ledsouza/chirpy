-- +goose Up
ALTER TABLE users
ADD COLUMN is_chirpy_red boolean DEFAULT false;

-- +goose Down
ALTER TABLE users
DELETE COLUMN is_chirpy_red;