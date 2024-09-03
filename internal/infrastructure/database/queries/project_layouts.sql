-- name: SaveProjectLayout :one
INSERT INTO project_layouts (layout_id, design_id, project_id, updated_at)
VALUES ($1, $2, $3, NOW())
RETURNING *;

-- name: ListProjectLayoutsByProject :many
SELECT layout.*
FROM layout
INNER JOIN project_layouts AS pl ON pl.layout_id = layout.id
WHERE project_id = $1;

-- name: DeleteProjectLayout :one
DELETE FROM project_layouts
WHERE id = 1
RETURNING *;

