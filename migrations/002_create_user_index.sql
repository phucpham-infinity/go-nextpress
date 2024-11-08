-- +goose Up
CREATE UNIQUE INDEX idx_uuid ON users (uuid);
CREATE UNIQUE INDEX idx_username ON users (username);
CREATE UNIQUE INDEX idx_email ON users (email);

-- +goose Down
DROP INDEX IF EXISTS idx_uuid;
DROP INDEX IF EXISTS idx_username;
DROP INDEX IF EXISTS idx_email;