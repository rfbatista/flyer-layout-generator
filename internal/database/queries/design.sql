
-- name: UpdateDesignByID :one
UPDATE design_element
SET
-- You can use sqlc.arg() and @ to identify named parameters
    name = CASE WHEN @name_do_update::boolean
        THEN sqlc.narg(name) ELSE name END,

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

