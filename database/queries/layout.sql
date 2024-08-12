-- name: ListLayouts :many
SELECT * FROM layout 
ORDER BY created_at desc
LIMIT $1 OFFSET $2;

-- name: GetLayoutByID :one
SELECT * FROM layout 
WHERE id = $1
LIMIT 1;

-- name: GetLayoutByRequestID :many
SELECT * FROM layout
WHERE request_id = $1
ORDER BY created_at desc;

-- name: GetOriginalLayoutByDesignID :one
SELECT * FROM layout 
WHERE design_id = $1 AND is_original = true
LIMIT 1;

-- name: GetLayoutComponentsByLayoutID :many
SELECT * FROM layout_components 
WHERE layout_id = $1
ORDER BY created_at desc;


-- name: GetLayoutElementsByLayoutID :many
SELECT * FROM layout_elements
WHERE layout_id = $1
ORDER BY created_at desc;

-- name: CreateLayout :one
INSERT INTO layout (width, height, design_id, request_id, is_original, image_url, stages) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: CreateLayoutComponent :one
INSERT INTO layout_components (
  layout_id,
  design_id, 
  width, 
  height, 
  color, 
  type, 
  xi, 
  xii, 
  yi, 
  yii, 
  bbox_xi, 
  bbox_xii, 
  bbox_yi, 
  bbox_yii
) VALUES (
  $1,         -- design_id
  $2,       -- width
  $3,       -- height
  $4,     -- color
  $5,   -- type (assuming COMPONENT_TYPE allows 'IMAGE')
  $6,        -- xi
  $7,        -- xii
  $8,        -- yi
  $9,        -- yii
  $10,        -- bbox_xi
  $11,        -- bbox_xii
  $12,        -- bbox_yi
  $13,         -- bbox_yii
  $14
)
RETURNING *;


-- name: CreateElement :one
INSERT INTO layout_elements (
  layout_id,
  layer_id,
  asset_id,
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
  inner_xi ,
  inner_xii,
  inner_yi ,
  inner_yii,
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
  $17,
  $18,
  $19,
  $20,
  $21,
  $22,
  $23
)
RETURNING *;

-- name: UpdateLayoutImagByID :exec
UPDATE layout 
SET 
  image_url = $2
WHERE id = $1;


-- name: DeleteLayoutByID :exec
DELETE FROM layout WHERE id = $1;
