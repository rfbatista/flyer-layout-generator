-- name: GetClientByID :one
SELECT *
FROM clients
WHERE id = $1
LIMIT 1;


-- name: ListClients :many
SELECT *
FROM clients
LIMIT $1 OFFSET $2;
