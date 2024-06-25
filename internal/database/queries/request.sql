-- name: CreateLayoutRequest :one
INSERT INTO layout_requests (design_id, layout_id, config)
VALUES ($1, $2, $3)
RETURNING *;

-- name: StartLayoutRequest :one
UPDATE layout_requests_jobs
SET
    started_at = now(),
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: FinishLayoutRequest :one
UPDATE layout_requests_jobs
SET
    finished_at = $2,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: CreateLayoutRequestJob :one
INSERT INTO layout_requests_jobs
(request_id, template_id, design_id, config)
VALUES ($1,$2, $3, $4)
RETURNING *;

-- name: UpdateLayoutRequest :one
UPDATE layout_requests
SET
    log = CASE WHEN @do_add_log::boolean
                    THEN sqlc.narg(log) ELSE log END,
    updated_at = now()
WHERE id = @layout_request_id
RETURNING *;

-- name: UpdateLayoutRequestJob :one
UPDATE layout_requests_jobs
SET
    log = CASE WHEN @do_add_log::boolean
                    THEN sqlc.narg(log) ELSE log END,
    error_at = CASE WHEN @do_add_error_at::boolean
                   THEN sqlc.narg(error_at) ELSE error_at END,
    finished_at = CASE WHEN @do_add_finished_at::boolean
                        THEN sqlc.narg(finished_at) ELSE finished_at END,
    started_at = CASE WHEN @do_add_started_at::boolean
                        THEN sqlc.narg(started_at) ELSE started_at END,
    stopped_at = CASE WHEN @do_add_stopped_at::boolean
                        THEN sqlc.narg(stopped_at) ELSE stopped_at END,
    status = CASE WHEN @do_add_status::boolean
                        THEN sqlc.narg(status) ELSE status END,
    image_url = CASE WHEN @do_add_image_url::boolean
                        THEN sqlc.narg(image_url) ELSE image_url END,
    layout_id = CASE WHEN @do_add_layout_id::boolean
                        THEN sqlc.narg(layout_id) ELSE layout_id END,
    updated_at = now()
WHERE id = @layout_request_job_id
RETURNING *;

-- name: ListLayoutRequestJobsNotStarted :many
SELECT *
FROM layout_requests_jobs
WHERE started_at is NULL
ORDER BY created_at DESC
LIMIT $1
;

-- name: ListLayoutRequestJobs :many
SELECT *
FROM layout_requests_jobs
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetLayoutRequestByID :one
SELECT *
FROM layout_requests
WHERE id = $1
LIMIT 1;

-- name: ListLayoutRequests :many
SELECT *
FROM layout_requests
ORDER BY created_at DESC
LIMIT $1 OFFSET $2
;

-- name: GetRequestJobsByRequestID :many
SELECT *
FROM layout_requests_jobs AS lrj 
WHERE lrj.request_id = $1
;

