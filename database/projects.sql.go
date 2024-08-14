// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: projects.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProject = `-- name: CreateProject :one
INSERT INTO projects (
  name,
  client_id,
  company_id,
  advertiser_id
) VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING id, client_id, advertiser_id, briefing, use_ai, name, created_at, updated_at, deleted_at, company_id
`

type CreateProjectParams struct {
	Name         string      `json:"name"`
	ClientID     pgtype.Int4 `json:"client_id"`
	CompanyID    pgtype.Int4 `json:"company_id"`
	AdvertiserID pgtype.Int4 `json:"advertiser_id"`
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRow(ctx, createProject,
		arg.Name,
		arg.ClientID,
		arg.CompanyID,
		arg.AdvertiserID,
	)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.ClientID,
		&i.AdvertiserID,
		&i.Briefing,
		&i.UseAi,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.CompanyID,
	)
	return i, err
}

const getProjectByID = `-- name: GetProjectByID :one
SELECT id, client_id, advertiser_id, briefing, use_ai, name, created_at, updated_at, deleted_at, company_id
FROM projects
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetProjectByID(ctx context.Context, id int64) (Project, error) {
	row := q.db.QueryRow(ctx, getProjectByID, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.ClientID,
		&i.AdvertiserID,
		&i.Briefing,
		&i.UseAi,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.CompanyID,
	)
	return i, err
}

const listProjects = `-- name: ListProjects :many
SELECT id, client_id, advertiser_id, briefing, use_ai, name, created_at, updated_at, deleted_at, company_id
FROM projects
WHERE company_id = $3
LIMIT $1 OFFSET $2
`

type ListProjectsParams struct {
	Limit     int32       `json:"limit"`
	Offset    int32       `json:"offset"`
	CompanyID pgtype.Int4 `json:"company_id"`
}

func (q *Queries) ListProjects(ctx context.Context, arg ListProjectsParams) ([]Project, error) {
	rows, err := q.db.Query(ctx, listProjects, arg.Limit, arg.Offset, arg.CompanyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.ClientID,
			&i.AdvertiserID,
			&i.Briefing,
			&i.UseAi,
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

const updateProjectByID = `-- name: UpdateProjectByID :one
UPDATE projects
SET 
    briefing = CASE WHEN $1::boolean
        THEN $2::text ELSE briefing END,
    name = CASE WHEN $3::boolean
        THEN $4::text ELSE name END,
    use_ai = CASE WHEN $5::boolean
        THEN $6::bool ELSE use_ai END
WHERE
  id = $7
RETURNING id, client_id, advertiser_id, briefing, use_ai, name, created_at, updated_at, deleted_at, company_id
`

type UpdateProjectByIDParams struct {
	BriefingDoUpdate bool   `json:"briefing_do_update"`
	Briefing         string `json:"briefing"`
	NameDoUpdate     bool   `json:"name_do_update"`
	Name             string `json:"name"`
	UseAiDoUpdate    bool   `json:"use_ai_do_update"`
	UseAi            bool   `json:"use_ai"`
	ID               int64  `json:"id"`
}

func (q *Queries) UpdateProjectByID(ctx context.Context, arg UpdateProjectByIDParams) (Project, error) {
	row := q.db.QueryRow(ctx, updateProjectByID,
		arg.BriefingDoUpdate,
		arg.Briefing,
		arg.NameDoUpdate,
		arg.Name,
		arg.UseAiDoUpdate,
		arg.UseAi,
		arg.ID,
	)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.ClientID,
		&i.AdvertiserID,
		&i.Briefing,
		&i.UseAi,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.CompanyID,
	)
	return i, err
}
