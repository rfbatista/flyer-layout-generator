-- name: GetPhotoshop :one
SELECT * FROM photoshop
WHERE id = $1 LIMIT 1;
