// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: adaptation_batch.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const cancelActiveAdaptationBatches = `-- name: CancelActiveAdaptationBatches :many
UPDATE adaptation_batch
SET
  status = 'canceled',
  updated_at = NOW()
WHERE user_id = $1
AND status <> 'canceled'
AND type = $2
RETURNING id, layout_id, design_id, request_id, user_id, type, removed_duplicates, status, started_at, finished_at, error_at, stopped_at, updated_at, created_at, config, log
`

type CancelActiveAdaptationBatchesParams struct {
	UserID pgtype.Int4 `json:"user_id"`
	Type   NullJobType `json:"type"`
}

func (q *Queries) CancelActiveAdaptationBatches(ctx context.Context, arg CancelActiveAdaptationBatchesParams) ([]AdaptationBatch, error) {
	rows, err := q.db.Query(ctx, cancelActiveAdaptationBatches, arg.UserID, arg.Type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AdaptationBatch
	for rows.Next() {
		var i AdaptationBatch
		if err := rows.Scan(
			&i.ID,
			&i.LayoutID,
			&i.DesignID,
			&i.RequestID,
			&i.UserID,
			&i.Type,
			&i.RemovedDuplicates,
			&i.Status,
			&i.StartedAt,
			&i.FinishedAt,
			&i.ErrorAt,
			&i.StoppedAt,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.Config,
			&i.Log,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const closeActiveAdaptation = `-- name: CloseActiveAdaptation :many
UPDATE adaptation_batch
SET
  status = 'finished',
  updated_at = NOW()
WHERE user_id = $1
AND status = 'finished'
AND type = $2
RETURNING id, layout_id, design_id, request_id, user_id, type, removed_duplicates, status, started_at, finished_at, error_at, stopped_at, updated_at, created_at, config, log
`

type CloseActiveAdaptationParams struct {
	UserID pgtype.Int4 `json:"user_id"`
	Type   NullJobType `json:"type"`
}

func (q *Queries) CloseActiveAdaptation(ctx context.Context, arg CloseActiveAdaptationParams) ([]AdaptationBatch, error) {
	rows, err := q.db.Query(ctx, closeActiveAdaptation, arg.UserID, arg.Type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AdaptationBatch
	for rows.Next() {
		var i AdaptationBatch
		if err := rows.Scan(
			&i.ID,
			&i.LayoutID,
			&i.DesignID,
			&i.RequestID,
			&i.UserID,
			&i.Type,
			&i.RemovedDuplicates,
			&i.Status,
			&i.StartedAt,
			&i.FinishedAt,
			&i.ErrorAt,
			&i.StoppedAt,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.Config,
			&i.Log,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createAdaptationBatch = `-- name: CreateAdaptationBatch :one
INSERT INTO adaptation_batch (
    layout_id, design_id, request_id, status, user_id, type,
    started_at, finished_at, error_at, stopped_at, updated_at, config, log
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
RETURNING id
`

type CreateAdaptationBatchParams struct {
	LayoutID   pgtype.Int4               `json:"layout_id"`
	DesignID   pgtype.Int4               `json:"design_id"`
	RequestID  pgtype.Int4               `json:"request_id"`
	Status     NullAdaptationBatchStatus `json:"status"`
	UserID     pgtype.Int4               `json:"user_id"`
	Type       NullJobType               `json:"type"`
	StartedAt  pgtype.Timestamp          `json:"started_at"`
	FinishedAt pgtype.Timestamp          `json:"finished_at"`
	ErrorAt    pgtype.Timestamp          `json:"error_at"`
	StoppedAt  pgtype.Timestamp          `json:"stopped_at"`
	UpdatedAt  pgtype.Timestamp          `json:"updated_at"`
	Config     pgtype.Text               `json:"config"`
	Log        pgtype.Text               `json:"log"`
}

func (q *Queries) CreateAdaptationBatch(ctx context.Context, arg CreateAdaptationBatchParams) (int64, error) {
	row := q.db.QueryRow(ctx, createAdaptationBatch,
		arg.LayoutID,
		arg.DesignID,
		arg.RequestID,
		arg.Status,
		arg.UserID,
		arg.Type,
		arg.StartedAt,
		arg.FinishedAt,
		arg.ErrorAt,
		arg.StoppedAt,
		arg.UpdatedAt,
		arg.Config,
		arg.Log,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getAdaptationBatchByID = `-- name: GetAdaptationBatchByID :one
SELECT id, layout_id, design_id, request_id, user_id, type, removed_duplicates, status, started_at, finished_at, error_at, stopped_at, updated_at, created_at, config, log FROM adaptation_batch WHERE id = $1
`

func (q *Queries) GetAdaptationBatchByID(ctx context.Context, id int64) (AdaptationBatch, error) {
	row := q.db.QueryRow(ctx, getAdaptationBatchByID, id)
	var i AdaptationBatch
	err := row.Scan(
		&i.ID,
		&i.LayoutID,
		&i.DesignID,
		&i.RequestID,
		&i.UserID,
		&i.Type,
		&i.RemovedDuplicates,
		&i.Status,
		&i.StartedAt,
		&i.FinishedAt,
		&i.ErrorAt,
		&i.StoppedAt,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.Config,
		&i.Log,
	)
	return i, err
}

const getAdaptationBatchByUser = `-- name: GetAdaptationBatchByUser :many
SELECT id, layout_id, design_id, request_id, user_id, type, removed_duplicates, status, started_at, finished_at, error_at, stopped_at, updated_at, created_at, config, log 
FROM adaptation_batch 
WHERE user_id = $1 
AND (
status = 'pending' 
OR status = 'started' 
OR status = 'pending' 
OR status = 'finished' 
)
AND type = $2
`

type GetAdaptationBatchByUserParams struct {
	UserID pgtype.Int4 `json:"user_id"`
	Type   NullJobType `json:"type"`
}

func (q *Queries) GetAdaptationBatchByUser(ctx context.Context, arg GetAdaptationBatchByUserParams) ([]AdaptationBatch, error) {
	rows, err := q.db.Query(ctx, getAdaptationBatchByUser, arg.UserID, arg.Type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AdaptationBatch
	for rows.Next() {
		var i AdaptationBatch
		if err := rows.Scan(
			&i.ID,
			&i.LayoutID,
			&i.DesignID,
			&i.RequestID,
			&i.UserID,
			&i.Type,
			&i.RemovedDuplicates,
			&i.Status,
			&i.StartedAt,
			&i.FinishedAt,
			&i.ErrorAt,
			&i.StoppedAt,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.Config,
			&i.Log,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getJobSummary = `-- name: GetJobSummary :one
SELECT coalesce(count(*), 0)::int as total,
coalesce(sum(
    case status
    when 'finished' then 1
    when 'error' then 1
    else 0 end), 0)::int as done
FROM layout_jobs
WHERE adaptation_batch_id = $1
`

type GetJobSummaryRow struct {
	Total int32 `json:"total"`
	Done  int32 `json:"done"`
}

func (q *Queries) GetJobSummary(ctx context.Context, adaptationBatchID pgtype.Int4) (GetJobSummaryRow, error) {
	row := q.db.QueryRow(ctx, getJobSummary, adaptationBatchID)
	var i GetJobSummaryRow
	err := row.Scan(&i.Total, &i.Done)
	return i, err
}

const listAdaptationBatch = `-- name: ListAdaptationBatch :many
SELECT id, layout_id, design_id, request_id, user_id, type, removed_duplicates, status, started_at, finished_at, error_at, stopped_at, updated_at, created_at, config, log FROM adaptation_batch LIMIT $1 OFFSET $2
`

type ListAdaptationBatchParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAdaptationBatch(ctx context.Context, arg ListAdaptationBatchParams) ([]AdaptationBatch, error) {
	rows, err := q.db.Query(ctx, listAdaptationBatch, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AdaptationBatch
	for rows.Next() {
		var i AdaptationBatch
		if err := rows.Scan(
			&i.ID,
			&i.LayoutID,
			&i.DesignID,
			&i.RequestID,
			&i.UserID,
			&i.Type,
			&i.RemovedDuplicates,
			&i.Status,
			&i.StartedAt,
			&i.FinishedAt,
			&i.ErrorAt,
			&i.StoppedAt,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.Config,
			&i.Log,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAdaptationBatch = `-- name: UpdateAdaptationBatch :one
UPDATE adaptation_batch
SET 
    status = CASE WHEN $1::boolean
        THEN $2 ELSE status END,
    started_at = CASE WHEN $3::boolean
        THEN $4 ELSE started_at END,
    finished_at = CASE WHEN $5::boolean
        THEN $6 ELSE finished_at END,
    error_at = CASE WHEN $7::boolean
        THEN $8 ELSE error_at END,
    stopped_at = CASE WHEN $9::boolean
        THEN $10 ELSE stopped_at END,
    removed_duplicates = CASE WHEN $11::boolean
        THEN $12 ELSE removed_duplicates END,
    log = CASE WHEN $9::boolean
        THEN $13 ELSE log END,
    updated_at = NOW()
WHERE
    id = $14
RETURNING id, layout_id, design_id, request_id, user_id, type, removed_duplicates, status, started_at, finished_at, error_at, stopped_at, updated_at, created_at, config, log
`

type UpdateAdaptationBatchParams struct {
	StatusDoUpdate            bool                      `json:"status_do_update"`
	Status                    NullAdaptationBatchStatus `json:"status"`
	StartedAtDoUpdate         bool                      `json:"started_at_do_update"`
	StartedAt                 pgtype.Timestamp          `json:"started_at"`
	FinishedAtDoUpdate        bool                      `json:"finished_at_do_update"`
	FinishedAt                pgtype.Timestamp          `json:"finished_at"`
	ErrorAtDoUpdate           bool                      `json:"error_at_do_update"`
	ErrorAt                   pgtype.Timestamp          `json:"error_at"`
	StoppedAtDoUpdate         bool                      `json:"stopped_at_do_update"`
	StoppedAt                 pgtype.Timestamp          `json:"stopped_at"`
	RemovedDuplicatesDoUpdate bool                      `json:"removed_duplicates_do_update"`
	RemovedDuplicates         pgtype.Bool               `json:"removed_duplicates"`
	Log                       pgtype.Text               `json:"log"`
	AdaptationID              pgtype.Int8               `json:"adaptation_id"`
}

func (q *Queries) UpdateAdaptationBatch(ctx context.Context, arg UpdateAdaptationBatchParams) (AdaptationBatch, error) {
	row := q.db.QueryRow(ctx, updateAdaptationBatch,
		arg.StatusDoUpdate,
		arg.Status,
		arg.StartedAtDoUpdate,
		arg.StartedAt,
		arg.FinishedAtDoUpdate,
		arg.FinishedAt,
		arg.ErrorAtDoUpdate,
		arg.ErrorAt,
		arg.StoppedAtDoUpdate,
		arg.StoppedAt,
		arg.RemovedDuplicatesDoUpdate,
		arg.RemovedDuplicates,
		arg.Log,
		arg.AdaptationID,
	)
	var i AdaptationBatch
	err := row.Scan(
		&i.ID,
		&i.LayoutID,
		&i.DesignID,
		&i.RequestID,
		&i.UserID,
		&i.Type,
		&i.RemovedDuplicates,
		&i.Status,
		&i.StartedAt,
		&i.FinishedAt,
		&i.ErrorAt,
		&i.StoppedAt,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.Config,
		&i.Log,
	)
	return i, err
}
