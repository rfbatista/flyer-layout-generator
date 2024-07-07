-- name: GetdesignElements :many
SELECT * FROM layout_elements 
WHERE design_id = $1;

-- name: GetElements :many
SELECT * FROM layout_elements 
WHERE design_id = $1;


-- name: GetElementsByLayoutID :many
SELECT * FROM layout_elements 
WHERE layout_id = $1;


-- name: GetdesignElementsByIDlist :many
select * from layout_elements 
where id = any (sqlc.slice(ids));


-- name: GetDesignElementsByComponentID :many
select * from layout_elements 
where component_id = $1;


-- name: GetLayoutElementByID :one
SELECT * FROM layout_elements 
WHERE id = $1
LIMIT 1;


-- name: UpdateLayoutElementPosition :one
UPDATE layout_elements
SET 
    xi              = $2,
    xii             = $3,
    yi              = $4,
    yii             = $5,
    inner_xi              = $6,
    inner_xii             = $7,
    inner_yi              = $8,
    inner_yii             = $9
WHERE
    id = $1
RETURNING *;


-- name: UpdateLayoutElementSize :one
UPDATE layout_elements
SET 
    xi              = $2,
    xii             = $3,
    yi              = $4,
    yii             = $5,
    inner_xi              = $6,
    inner_xii             = $7,
    inner_yi              = $8,
    inner_yii             = $9,
    width                 = $10,
    height                = $11
WHERE
    id = $1
RETURNING *;
