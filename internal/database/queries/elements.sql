-- name: GetPhotoshopElements :many
SELECT * FROM photoshop_element 
WHERE photoshop_id = $1;

-- name: GetElements :many
SELECT * FROM photoshop_element 
WHERE photoshop_id = $1;


-- name: GetphotoshopElementsByIDlist :many
select * from photoshop_element 
where id = any (sqlc.slice(ids));
