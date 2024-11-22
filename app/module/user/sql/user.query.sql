-- name: GetManyUsers :many
SELECT id, username, email, activation_key, status, created_at, updated_at, registered_at, deleted_at 
FROM users;

-- name: CreateUser :one
INSERT INTO users (username, password, email, activation_key)
VALUES ($1, $2, $3, $4) 
RETURNING id, username, email, activation_key, status, created_at, updated_at, registered_at, deleted_at;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: ActivateUser :one
UPDATE users SET status = 'active', activation_key = NULL, updated_at = now(), registered_at = now() 
WHERE email = $1 
RETURNING id, username, email, activation_key, status, created_at, updated_at, registered_at, deleted_at;