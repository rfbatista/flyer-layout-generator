-- name: CreateReplicationBatch :one
INSERT INTO adaptation_batch (
    layout_id, design_id, request_id, status, user_id,
    started_at, finished_at, error_at, stopped_at, updated_at, config, log
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING id;
-- name: ListReplicationBatch :many
SELECT * FROM adaptation_batch LIMIT $1 OFFSET $2;

-- name: GetReplicationBatchByID :one
SELECT * FROM adaptation_batch WHERE id = $1;

-- name: GetReplicationBatchByUser :many
SELECT * 
FROM adaptation_batch 
WHERE user_id = $1 
AND (
status = 'pending' 
OR status = 'started' 
OR status = 'pending' 
OR status = 'finished' 
)
;

-- name: UpdateReplicationBatch :one
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
    id = sqlc.narg(adaptation_id)
RETURNING *;

-- name: CancelActiveReplicationBatches :many
UPDATE adaptation_batch
SET
  status = 'canceled',
  updated_at = NOW()
WHERE user_id = $1 AND status <> 'canceled'
RETURNING *;

-- name: CloseActiveReplication :many
UPDATE adaptation_batch
SET
  status = 'finished',
  updated_at = NOW()
WHERE user_id = $1 AND status = 'finished'
RETURNING *;

-- name: GetReplicationSummary :one
SELECT coalesce(count(*), 0)::int as total,
coalesce(sum(
    case status
    when 'finished' then 1
    when 'error' then 1
    else 0 end), 0)::int as done
FROM layout_jobs
WHERE adaptation_batch_id = $1;
