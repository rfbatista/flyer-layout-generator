-- name: CreateAdaptationBatch :one
INSERT INTO adaptation_batch (
    layout_id, design_id, request_id, template_id, status, 
    started_at, finished_at, error_at, stopped_at, updated_at, config, log
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING id;

-- name: ListAdaptationBatch :many
SELECT * FROM adaptation_batch LIMIT $1 OFFSET $2;

-- name: GetAdaptationBatchByID :one
SELECT * FROM adaptation_batch WHERE id = $1;

-- name: GetAdaptationBatchByUser :many
SELECT * 
FROM adaptation_batch 
WHERE user_id = $1 
AND (status <> sqlc.narg(status) OR NOT @filter_by_status)
;

-- name: UpdateAdaptationBatch :one
UPDATE adaptation_batch
SET 
    status = CASE WHEN @status_do_update::boolean
        THEN sqlc.narg(status) ELSE status END,
    started_at = CASE WHEN @started_at_do_update::boolean
        THEN sqlc.narg(started_at) ELSE started_at END,
    finished_at = CASE WHEN @finished_at_do_update::boolean
        THEN sqlc.narg(finished_at) ELSE finished_at END,
    error_at = CASE WHEN @error_at_do_update::boolean
        THEN sqlc.narg(error_at) ELSE error_at END,
    stopped_at = CASE WHEN @stopped_at_do_update::boolean
        THEN sqlc.narg(stopped_at) ELSE stopped_at END,
    log = CASE WHEN @stopped_at_do_update::boolean
        THEN sqlc.narg(log) ELSE log END,
    updated_at = NOW()
WHERE
    id = ANY (sqlc.slice(ids)) and design_id = @design_id
RETURNING *;
