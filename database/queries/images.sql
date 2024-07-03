-- name: ListImagesGenerated :many
SELECT * FROM images
OFFSET $1 LIMIT $2;
