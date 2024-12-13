-- name: GetAllRoles :many
SELECT * 
FROM roles;

-- name: GetByIdRoles :one
SELECT * 
FROM roles
WHERE id = $1
LIMIT 1;

-- name: CreateRole :one
INSERT INTO roles (role, permission)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateRoleById :one
UPDATE roles
    set role = $2, 
    permission = $3
WHERE id = $1
RETURNING *;

-- name: DeleteRoleById :one
DELETE FROM roles
WHERE id = $1
RETURNING *;

-- name: SoftDeleteRoleById :one
UPDATE roles
SET deleted_at = now()
WHERE id = $1
RETURNING *;

-- name: RestoreRoleById :one
UPDATE roles
SET deleted_at = NULL
WHERE id = $1
RETURNING *;

-- name: GetManyRoles :many
SELECT *
FROM roles
WHERE
  (sqlc.arg(filter_role) IS NULL OR role = sqlc.arg(filter_role))
ORDER BY
    CASE
        WHEN sqlc.arg(sort_order) = 'asc' THEN
            CASE sqlc.arg(sort_name)
                WHEN 'role' THEN role
                WHEN 'created_at' THEN created_at
                WHEN 'updated_at' THEN updated_at
                END
        END ASC,
    CASE
        WHEN sqlc.arg(sort_order) = 'desc' THEN
            CASE sqlc.arg(sort_name)
                WHEN 'role' THEN role
                WHEN 'created_at' THEN created_at
                WHEN 'updated_at' THEN updated_at
                END
        END DESC
LIMIT sqlc.arg(page_limit) OFFSET (sqlc.arg(page) - 1) * sqlc.arg(page_limit);

-- name: GetTotalRoles :one
SELECT COUNT(*) AS total_items
FROM roles
WHERE
  (sqlc.arg(filter_role) IS NULL OR role = sqlc.arg(filter_role));