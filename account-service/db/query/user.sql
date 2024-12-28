-- name: CreateUser :one
INSERT INTO "user" (
    account_id, name, lastname, email, username, password, phone, is_admin, role_id, active
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"
WHERE account_id = $1 AND id = $2 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM "user"
WHERE username = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM "user"
WHERE account_id = $1
ORDER BY name, lastname
LIMIT $2
OFFSET $3;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE account_id = $1 AND id = $2;
