-- name: GetPhotoshop :one
SELECT * FROM photoshop
WHERE id = $1 LIMIT 1;

-- name: GetPhotoshopBackgroundComponent :one
SELECT * FROM photoshop_components
WHERE photoshop_id = $1 AND type = 'background' LIMIT 1;

-- name: ListPhotoshop :many
SELECT * FROM photoshop
OFFSET $1 LIMIT $2;

-- name: CreateComponent :one
INSERT INTO photoshop_components (
  photoshop_id,
  width,
  height,
  type
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;


-- name: UpdateManyPhotoshopElement :many
UPDATE photoshop_element
SET 
-- You can use sqlc.arg() and @ to identify named parameters
    component_id = CASE WHEN @component_id_do_update::boolean
        THEN @component_id::int ELSE component_id END,

    name = CASE WHEN @name_do_update::boolean
        THEN @name::text ELSE name END
WHERE
    id IN (sqlc.slice(ids)) and photoshop_id = @photoshop_id
RETURNING *;

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
  $16
)
RETURNING *;

-- name: CreatePhotoshop :one
INSERT INTO photoshop (
  name,
  image_url,
  file_url,
  width,
  height
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;

-- name: ListPhotoshopElements :many
SELECT * FROM photoshop_element 
WHERE photoshop_id = $1
LIMIT $2 OFFSET $3;

-- name: GetPhotoshopElements :many
SELECT * FROM photoshop_element 
WHERE photoshop_id = $1;

