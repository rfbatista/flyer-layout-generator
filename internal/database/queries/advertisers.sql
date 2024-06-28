-- name: GetAdvertiserByID :one
SELECT *
FROM advertisers
WHERE id = $1
LIMIT 1;

-- name: ListAdvertisers :many
SELECT *
FROM advertisers
LIMIT $1 OFFSET $2;
