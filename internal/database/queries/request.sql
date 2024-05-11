-- name: StartRequestProcess :one
INSERT INTO request (name, started_at)
VALUES ($1, $2)
RETURNING *;

-- name: FinishRequestProcess :one
UPDATE request
SET
    finished_at = $2,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: StarRequestStep :one
INSERT INTO request_steps
( name, request_id, started_at)
VALUES ($1,$2,$3)
RETURNING *;

-- name: UpdateRequestStep :one
UPDATE request_steps
SET
    log = CASE WHEN @do_add_log::boolean
                    THEN sqlc.narg(log) ELSE log END,
    error_at = CASE WHEN @do_add_error_at::boolean
                   THEN sqlc.narg(error_at) ELSE error_at END,
    finished_at = CASE WHEN @do_add_finished_at::boolean
                        THEN sqlc.narg(finished_at) ELSE finished_at END,
    updated_at = now()
WHERE id = @request_step_id
RETURNING *;
