-- name: Getplan :one
SELECT * FROM plan
WHERE id = $1 LIMIT 1;

-- name: ListPlans :many
SELECT * FROM plan
ORDER BY name
LIMIT $1
OFFSET $2;
