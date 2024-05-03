-- name: GetPhotoshop :one
SELECT * FROM photoshop
WHERE id = $1 LIMIT 1;

-- name: GetPhotoshopComponentByID :one
SELECT * FROM photoshop_components
WHERE photoshop_id = $1 LIMIT 1;

-- name: GetPhotoshopBackgroundComponent :one
SELECT * FROM photoshop_components
WHERE photoshop_id = $1 AND type = 'background' LIMIT 1;

-- name: ListPhotoshop :many
SELECT * FROM photoshop
OFFSET $1 LIMIT $2;

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
  image_url,
  image_extension
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
  file_url,
  width,
  height,
  image_extension
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: ListPhotoshopElements :many
SELECT * FROM photoshop_element 
WHERE photoshop_id = $1
LIMIT $2 OFFSET $3;
