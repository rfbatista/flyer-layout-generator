-- name: CreateRendererJob :one
INSERT INTO renderer_jobs (layout_id, adaptation_id)
VALUES 
  ($1, $2)
RETURNING id;


-- name: ListRendererJobs :many
SELECT * 
FROM renderer_jobs
WHERE adaptation_id = $1 
AND (status <> sqlc.narg(status) OR NOT @filter_by_status)
;

-- name: UpdateRendererJob :one
UPDATE renderer_jobs
SET 
    status = CASE WHEN @status_do_update::boolean
        THEN sqlc.narg(status) ELSE status END,
    image_id = CASE WHEN @image_do_update::boolean
        THEN sqlc.narg(image_id) ELSE image_id END,
    started_at = CASE WHEN @started_at_do_update::boolean
        THEN sqlc.narg(started_at) ELSE started_at END,
    finished_at = CASE WHEN @finished_at_do_update::boolean
        THEN sqlc.narg(finished_at) ELSE finished_at END,
    error_at = CASE WHEN @error_at_do_update::boolean
        THEN sqlc.narg(error_at) ELSE error_at END,
    log = CASE WHEN @log_do_update::boolean
        THEN sqlc.narg(log) ELSE log END,
    updated_at = NOW()
WHERE id = 1
RETURNING *;
