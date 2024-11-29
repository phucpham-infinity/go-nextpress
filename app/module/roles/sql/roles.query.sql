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

--name: UpdateOneRoles :one
UPDATE roles
SET role = $2, permission = $3
WHERE id = $1
RETURNING *;

--name: DeleteOneRoles :one
DELETE FROM roles
WHERE id = $1
RETURNING *;

-- name: GetManyRoles :many
SELECT *
FROM roles
WHERE
    (:text_search IS NULL OR role ILIKE CONCAT('%', :text_search, '%'))
  AND (:filter_role IS NULL OR role = :filter_role)
ORDER BY
    CASE
        WHEN :sort_order = 'asc' THEN
            CASE :sort_column
                WHEN 'role' THEN role
                WHEN 'created_at' THEN created_at
                WHEN 'updated_at' THEN updated_at
                END
        END ASC,
    CASE
        WHEN :sort_order = 'desc' THEN
            CASE :sort_column
                WHEN 'role' THEN role
                WHEN 'created_at' THEN created_at
                WHEN 'updated_at' THEN updated_at
                END
        END DESC
LIMIT :page_limit OFFSET (:page - 1) * :page_limit;
