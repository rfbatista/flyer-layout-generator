-- name: CreateProject :one
INSERT INTO projects (
  name,
  client_id,
  advertiser_id
) VALUES (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: GetProjectByID :one
SELECT *
FROM projects
WHERE id = $1
LIMIT 1;


-- name: ListProjects :many
SELECT *
FROM projects
LIMIT $1 OFFSET $2;
