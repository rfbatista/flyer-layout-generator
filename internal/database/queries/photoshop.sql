-- name: Getdesign :one
SELECT * FROM design
WHERE id = $1 LIMIT 1;

-- name: GetdesignComponentByID :one
SELECT * FROM design_components
WHERE design_id = $1 LIMIT 1;

-- name: GetdesignBackgroundComponent :one
SELECT * FROM design_components
WHERE design_id = $1 AND type = 'background' LIMIT 1;

-- name: Listdesign :many
SELECT * FROM design
OFFSET $1 LIMIT $2;

-- name: CreateElement :one
INSERT INTO design_element (
  layer_id,
  design_id,
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

-- name: ListdesignElements :many
SELECT * FROM design_element 
WHERE design_id = $1
LIMIT $2 OFFSET $3;
