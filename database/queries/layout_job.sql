-- name: CreateLayoutJob :one
INSERT INTO layout_jobs (
    based_on_layout_id, status, user_id, template_id,
    started_at, finished_at, error_at, updated_at, config, log
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9,$10
)
RETURNING id;

-- name: GetLayoutJobByID :one
SELECT * FROM layout_jobs WHERE id = $1;

-- name: UpdateLayoutJob :one
UPDATE layout_jobs
SET 
    status = CASE WHEN @status_do_update::boolean
        THEN sqlc.narg(status) ELSE status END,
    started_at = CASE WHEN @started_at_do_update::boolean
        THEN sqlc.narg(started_at) ELSE started_at END,
    finished_at = CASE WHEN @finished_at_do_update::boolean
        THEN sqlc.narg(finished_at) ELSE finished_at END,
    error_at = CASE WHEN @error_at_do_update::boolean
        THEN sqlc.narg(error_at) ELSE error_at END,
    log = CASE WHEN @stopped_at_do_update::boolean
        THEN sqlc.narg(log) ELSE log END,
    created_layout_id= CASE WHEN @created_layout_do_update::boolean
        THEN sqlc.narg(created_layout_id) ELSE created_layout_id END,
    updated_at = NOW()
WHERE
    id = sqlc.narg(id)
RETURNING *;

-- name: CancelLayoutJob :many
UPDATE layout_jobs
SET
  status = 'canceled',
  updated_at = NOW()
WHERE status <> 'canceled'
AND (adaptation_batch_id = $1 OR NOT @filter_by_adaptation)
AND (replication_batch_id = $1 OR NOT @filter_by_replication)
RETURNING *;
