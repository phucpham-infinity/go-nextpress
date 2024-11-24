-- name: GetAllRoles :many
SELECT * 
FROM roles;

-- name: GetByIdRoles :one
SELECT * 
FROM roles
WHERE id = $1
LIMIT 1;

-- name: CreateOneRoles :one
INSERT INTO roles (role, permission)
VALUES ($1, $2)
RETURNING *;

