-- name: GetManyUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (username, password, email, activation_key)
VALUES ($1, $2, $3, $4) 
RETURNING uuid, username, email, activation_key, status, created_at, updated_at, registered_at, deleted_at;