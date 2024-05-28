-- name: ListLayouts :many
SELECT * FROM layout 
ORDER BY created_at desc
LIMIT $1 OFFSET $2;

-- name: GetLayoutComponentsByLayoutID :many
SELECT * FROM layout_components 
WHERE layout_id = $1
ORDER BY created_at desc;

-- name: GetLayoutTemplateByLayoutID :many
SELECT * FROM layout_template 
WHERE layout_id = $1
ORDER BY created_at desc;

-- name: GetLayoutRegionByLayoutID :many
SELECT * FROM layout_region
WHERE layout_id = $1
ORDER BY created_at desc;

-- name: CreateLayout :one
INSERT INTO layout (width, height) VALUES ($1, $2) RETURNING *;

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

-- name: CreateLayoutTemplate :one
INSERT INTO layout_template (
  layout_id,
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

-- name: CreateLayoutRegion :one
INSERT INTO layout_region (
  layout_id,
  xi, 
  xii, 
  yi, 
  yii 
) VALUES (
  $1,  -- url
  $2,                             -- photoshop_id
  $3,                             -- template_id
  $4,           -- created_at
  $5
) 
RETURNING *;
