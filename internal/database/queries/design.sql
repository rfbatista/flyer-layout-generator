
-- name: UpdateDesignByID :one
UPDATE design
SET
-- You can use sqlc.arg() and @ to identify named parameters
    name = CASE WHEN @name_do_update::boolean
        THEN sqlc.narg(name) ELSE name END,

    image_url = CASE WHEN @image_url_do_update::boolean
        THEN sqlc.narg(image_url) ELSE name END,

    width = CASE WHEN @width_do_update::boolean
        THEN sqlc.narg(width) ELSE width END,

    height = CASE WHEN @height_do_update::boolean
        THEN sqlc.narg(height) ELSE height END

WHERE
    id = @design_id
RETURNING *;

-- name: SetDesignAsProccessed :one
UPDATE design
SET
    is_proccessed = true
WHERE
    id = @design_id
RETURNING *;


-- name: Createdesign :one
INSERT INTO design (
  name,
  file_url
) VALUES (
  $1,
  $2
)
RETURNING *;

