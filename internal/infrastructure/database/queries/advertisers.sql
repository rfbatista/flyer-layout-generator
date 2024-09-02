-- name: GetAdvertiserByID :one
SELECT *
FROM advertisers
WHERE id = $1
LIMIT 1;

-- name: ListAdvertisers :many
SELECT *
FROM advertisers
WHERE company_id = $3
LIMIT $1 OFFSET $2;

-- name: CreateAdvertiser :one
INSERT INTO advertisers (name, company_id)
VALUES ($1, $2)
RETURNING id;
