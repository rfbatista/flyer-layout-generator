-- name: GetPhotoshop :one
SELECT * FROM photoshop
WHERE id = $1 LIMIT 1;


-- name: SetComponent :exec
UPDATE photoshop_element 
SET component_id = $2, component_type = $3
WHERE id = $1;


-- name: CreateElement :one
INSERT INTO photoshop_element (
  layer_id,
  photoshop_id,
  name,
  text,
  xi,
  xii,
  yi,
  yii,
  width,
  height,
  is_group,
  group_id,
  level,
  kind,
  component_id,
  component_type,
  image_url
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8,
  $9,
  $10,
  $11,
  $12,
  $13,
  $14,
  $15,
  $16,
  $17
)
RETURNING *;


-- name: CreatePhotoshop :one
INSERT INTO photoshop (
  name,
  image_url,
  file_url
) VALUES (
  $1,
  $2,
  $3
)
RETURNING *;
