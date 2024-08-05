-- name: GetClientByID :one
SELECT *
FROM clients
WHERE id = $1
LIMIT 1;


-- name: ListClients :many
SELECT *
FROM clients
LIMIT $1 OFFSET $2;

-- name: CreateClient :one
INSERT INTO clients (name, company_id)
VALUES ($1, $2)
RETURNING id;
