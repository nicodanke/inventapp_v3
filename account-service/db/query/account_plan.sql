-- name: CreateAccountPlan :one
INSERT INTO account_plan (
    plan_id, account_id, price
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateAccountPlan :one
UPDATE account_plan
SET
    ended_at = COALESCE(sqlc.narg(ended_at), ended_at)
WHERE
    account_id = sqlc.arg(account_id) AND ended_at IS NULL
RETURNING *;
