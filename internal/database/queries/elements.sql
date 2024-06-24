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
