-- name: CreateProject :one
INSERT INTO projects (
  name,
  client_id,
  company_id,
  advertiser_id
) VALUES (
  $1,
  $2,
  $3,
  $4
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
WHERE company_id = $3
LIMIT $1 OFFSET $2;

-- name: UpdateProjectByID :one
UPDATE projects
SET 
    briefing = CASE WHEN @briefing_do_update::boolean
        THEN @briefing::text ELSE briefing END,
    name = CASE WHEN @name_do_update::boolean
        THEN @name::text ELSE name END,
    use_ai = CASE WHEN @use_ai_do_update::boolean
        THEN @use_ai::bool ELSE use_ai END
WHERE
  id = @id
RETURNING *;
