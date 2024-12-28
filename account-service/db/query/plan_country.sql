-- name: GetPlanCountry :one
SELECT * FROM plan_country
WHERE plan_id = $1 AND country = $2 LIMIT 1;
