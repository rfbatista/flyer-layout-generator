// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: advertisers.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAdvertiser = `-- name: CreateAdvertiser :one
INSERT INTO advertisers (name, company_id)
VALUES ($1, $2)
RETURNING id
`

type CreateAdvertiserParams struct {
	Name      string      `json:"name"`
	CompanyID pgtype.Int4 `json:"company_id"`
}

func (q *Queries) CreateAdvertiser(ctx context.Context, arg CreateAdvertiserParams) (int64, error) {
	row := q.db.QueryRow(ctx, createAdvertiser, arg.Name, arg.CompanyID)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getAdvertiserByID = `-- name: GetAdvertiserByID :one
SELECT id, name, created_at, updated_at, deleted_at, company_id
FROM advertisers
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAdvertiserByID(ctx context.Context, id int64) (Advertiser, error) {
	row := q.db.QueryRow(ctx, getAdvertiserByID, id)
	var i Advertiser
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.CompanyID,
	)
	return i, err
}

const listAdvertisers = `-- name: ListAdvertisers :many
SELECT id, name, created_at, updated_at, deleted_at, company_id
FROM advertisers
WHERE company_id = $3
LIMIT $1 OFFSET $2
`

type ListAdvertisersParams struct {
	Limit     int32       `json:"limit"`
	Offset    int32       `json:"offset"`
	CompanyID pgtype.Int4 `json:"company_id"`
}

func (q *Queries) ListAdvertisers(ctx context.Context, arg ListAdvertisersParams) ([]Advertiser, error) {
	rows, err := q.db.Query(ctx, listAdvertisers, arg.Limit, arg.Offset, arg.CompanyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Advertiser
	for rows.Next() {
		var i Advertiser
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.CompanyID,
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
