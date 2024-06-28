// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: clients.sql

package database

import (
	"context"
)

const getClientByID = `-- name: GetClientByID :one
SELECT id, name, created_at, updated_at, deleted_at
FROM clients
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetClientByID(ctx context.Context, id int64) (Client, error) {
	row := q.db.QueryRow(ctx, getClientByID, id)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listClients = `-- name: ListClients :many
SELECT id, name, created_at, updated_at, deleted_at
FROM clients
LIMIT $1 OFFSET $2
`

type ListClientsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListClients(ctx context.Context, arg ListClientsParams) ([]Client, error) {
	rows, err := q.db.Query(ctx, listClients, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Client
	for rows.Next() {
		var i Client
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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
