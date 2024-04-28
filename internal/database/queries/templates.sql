-- name: ListTemplates :many
SELECT sqlc.embed(templates), sqlc.embed(templates_slots), sqlc.embed(templates_distortions)
FROM templates
JOIN templates_slots ON templates_slots.template_id = templates.id
JOIN templates_distortions ON templates_distortions.template_id = templates.id
LIMIT $1 OFFSET $2;


-- name: GetTemplate :one
SELECT sqlc.embed(templates)
FROM templates
WHERE templates.id = $1 LIMIT 1;


-- name: GetTemplateSlots :many
SELECT sqlc.embed(templates_slots)
FROM templates_slots
WHERE template_id = $1;

-- name: GetTemplateDistortion :one
SELECT sqlc.embed(templates_distortions)
FROM templates_distortions
WHERE template_id = $1 LIMIT 1;


-- name: CreateTemplate :one
INSERT INTO templates (
  name,
  type,
  width,
  height
) VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING *;

-- name: CreateTemplateSlot :one
INSERT INTO templates_slots (
  xi,
  yi,
  width,
  height,
  template_id
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;

-- name: CreateTemplateDistortions :one
INSERT INTO templates_distortions (
  x,
  y,
  template_id
) VALUES (
  $1,
  $2,
  $3
)
RETURNING *;
