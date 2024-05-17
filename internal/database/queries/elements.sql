-- name: GetdesignElements :many
SELECT * FROM design_element 
WHERE design_id = $1;

-- name: GetElements :many
SELECT * FROM design_element 
WHERE design_id = $1;


-- name: GetdesignElementsByIDlist :many
select * from design_element 
where id = any (sqlc.slice(ids));


-- name: GetDesignElementsByComponentID :many
select * from design_element 
where component_id = $1;
