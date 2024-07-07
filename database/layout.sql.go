// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: layout.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createElement = `-- name: CreateElement :one
INSERT INTO layout_elements (
  layout_id,
  layer_id,
  asset_id,
  design_id,
  name,
  text,
  xi,
  xii,
  yi,
  yii,
  width,
  height,
  is_group,
  group_id,
  level,
  kind,
  component_id,
  image_url,
  inner_xi ,
  inner_xii,
  inner_yi ,
  inner_yii,
  image_extension
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8,
  $9,
  $10,
  $11,
  $12,
  $13,
  $14,
  $15,
  $16,
  $17,
  $18,
  $19,
  $20,
  $21,
  $22,
  $23
)
RETURNING id, design_id, layout_id, component_id, asset_id, name, layer_id, text, xi, xii, yi, yii, inner_xi, inner_xii, inner_yi, inner_yii, width, height, is_group, group_id, level, kind, image_url, image_extension, created_at, updated_at
`

type CreateElementParams struct {
	LayoutID       int32       `json:"layout_id"`
	LayerID        pgtype.Text `json:"layer_id"`
	AssetID        int32       `json:"asset_id"`
	DesignID       int32       `json:"design_id"`
	Name           pgtype.Text `json:"name"`
	Text           pgtype.Text `json:"text"`
	Xi             pgtype.Int4 `json:"xi"`
	Xii            pgtype.Int4 `json:"xii"`
	Yi             pgtype.Int4 `json:"yi"`
	Yii            pgtype.Int4 `json:"yii"`
	Width          pgtype.Int4 `json:"width"`
	Height         pgtype.Int4 `json:"height"`
	IsGroup        pgtype.Bool `json:"is_group"`
	GroupID        pgtype.Int4 `json:"group_id"`
	Level          pgtype.Int4 `json:"level"`
	Kind           pgtype.Text `json:"kind"`
	ComponentID    pgtype.Int4 `json:"component_id"`
	ImageUrl       pgtype.Text `json:"image_url"`
	InnerXi        pgtype.Int4 `json:"inner_xi"`
	InnerXii       pgtype.Int4 `json:"inner_xii"`
	InnerYi        pgtype.Int4 `json:"inner_yi"`
	InnerYii       pgtype.Int4 `json:"inner_yii"`
	ImageExtension pgtype.Text `json:"image_extension"`
}

func (q *Queries) CreateElement(ctx context.Context, arg CreateElementParams) (LayoutElement, error) {
	row := q.db.QueryRow(ctx, createElement,
		arg.LayoutID,
		arg.LayerID,
		arg.AssetID,
		arg.DesignID,
		arg.Name,
		arg.Text,
		arg.Xi,
		arg.Xii,
		arg.Yi,
		arg.Yii,
		arg.Width,
		arg.Height,
		arg.IsGroup,
		arg.GroupID,
		arg.Level,
		arg.Kind,
		arg.ComponentID,
		arg.ImageUrl,
		arg.InnerXi,
		arg.InnerXii,
		arg.InnerYi,
		arg.InnerYii,
		arg.ImageExtension,
	)
	var i LayoutElement
	err := row.Scan(
		&i.ID,
		&i.DesignID,
		&i.LayoutID,
		&i.ComponentID,
		&i.AssetID,
		&i.Name,
		&i.LayerID,
		&i.Text,
		&i.Xi,
		&i.Xii,
		&i.Yi,
		&i.Yii,
		&i.InnerXi,
		&i.InnerXii,
		&i.InnerYi,
		&i.InnerYii,
		&i.Width,
		&i.Height,
		&i.IsGroup,
		&i.GroupID,
		&i.Level,
		&i.Kind,
		&i.ImageUrl,
		&i.ImageExtension,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createLayout = `-- name: CreateLayout :one
INSERT INTO layout (width, height, design_id, is_original, image_url, stages) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, design_id, is_original, image_url, width, height, data, stages, created_at, updated_at, deleted_at
`

type CreateLayoutParams struct {
	Width      pgtype.Int4 `json:"width"`
	Height     pgtype.Int4 `json:"height"`
	DesignID   pgtype.Int4 `json:"design_id"`
	IsOriginal pgtype.Bool `json:"is_original"`
	ImageUrl   pgtype.Text `json:"image_url"`
	Stages     []string    `json:"stages"`
}

func (q *Queries) CreateLayout(ctx context.Context, arg CreateLayoutParams) (Layout, error) {
	row := q.db.QueryRow(ctx, createLayout,
		arg.Width,
		arg.Height,
		arg.DesignID,
		arg.IsOriginal,
		arg.ImageUrl,
		arg.Stages,
	)
	var i Layout
	err := row.Scan(
		&i.ID,
		&i.DesignID,
		&i.IsOriginal,
		&i.ImageUrl,
		&i.Width,
		&i.Height,
		&i.Data,
		&i.Stages,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const createLayoutComponent = `-- name: CreateLayoutComponent :one
INSERT INTO layout_components (
  layout_id,
  design_id, 
  width, 
  height, 
  color, 
  type, 
  xi, 
  xii, 
  yi, 
  yii, 
  bbox_xi, 
  bbox_xii, 
  bbox_yi, 
  bbox_yii
) VALUES (
  $1,         -- design_id
  $2,       -- width
  $3,       -- height
  $4,     -- color
  $5,   -- type (assuming COMPONENT_TYPE allows 'IMAGE')
  $6,        -- xi
  $7,        -- xii
  $8,        -- yi
  $9,        -- yii
  $10,        -- bbox_xi
  $11,        -- bbox_xii
  $12,        -- bbox_yi
  $13,         -- bbox_yii
  $14
)
RETURNING id, layout_id, design_id, width, height, is_original, color, type, xi, xii, yi, yii, bbox_xi, bbox_xii, bbox_yi, bbox_yii, priority, inner_xi, inner_xii, inner_yi, inner_yii, created_at
`

type CreateLayoutComponentParams struct {
	LayoutID int32             `json:"layout_id"`
	DesignID int32             `json:"design_id"`
	Width    pgtype.Int4       `json:"width"`
	Height   pgtype.Int4       `json:"height"`
	Color    pgtype.Text       `json:"color"`
	Type     NullComponentType `json:"type"`
	Xi       pgtype.Int4       `json:"xi"`
	Xii      pgtype.Int4       `json:"xii"`
	Yi       pgtype.Int4       `json:"yi"`
	Yii      pgtype.Int4       `json:"yii"`
	BboxXi   pgtype.Int4       `json:"bbox_xi"`
	BboxXii  pgtype.Int4       `json:"bbox_xii"`
	BboxYi   pgtype.Int4       `json:"bbox_yi"`
	BboxYii  pgtype.Int4       `json:"bbox_yii"`
}

func (q *Queries) CreateLayoutComponent(ctx context.Context, arg CreateLayoutComponentParams) (LayoutComponent, error) {
	row := q.db.QueryRow(ctx, createLayoutComponent,
		arg.LayoutID,
		arg.DesignID,
		arg.Width,
		arg.Height,
		arg.Color,
		arg.Type,
		arg.Xi,
		arg.Xii,
		arg.Yi,
		arg.Yii,
		arg.BboxXi,
		arg.BboxXii,
		arg.BboxYi,
		arg.BboxYii,
	)
	var i LayoutComponent
	err := row.Scan(
		&i.ID,
		&i.LayoutID,
		&i.DesignID,
		&i.Width,
		&i.Height,
		&i.IsOriginal,
		&i.Color,
		&i.Type,
		&i.Xi,
		&i.Xii,
		&i.Yi,
		&i.Yii,
		&i.BboxXi,
		&i.BboxXii,
		&i.BboxYi,
		&i.BboxYii,
		&i.Priority,
		&i.InnerXi,
		&i.InnerXii,
		&i.InnerYi,
		&i.InnerYii,
		&i.CreatedAt,
	)
	return i, err
}

const getLayoutByID = `-- name: GetLayoutByID :one
SELECT id, design_id, is_original, image_url, width, height, data, stages, created_at, updated_at, deleted_at FROM layout 
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetLayoutByID(ctx context.Context, id int64) (Layout, error) {
	row := q.db.QueryRow(ctx, getLayoutByID, id)
	var i Layout
	err := row.Scan(
		&i.ID,
		&i.DesignID,
		&i.IsOriginal,
		&i.ImageUrl,
		&i.Width,
		&i.Height,
		&i.Data,
		&i.Stages,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getLayoutComponentsByLayoutID = `-- name: GetLayoutComponentsByLayoutID :many
SELECT id, layout_id, design_id, width, height, is_original, color, type, xi, xii, yi, yii, bbox_xi, bbox_xii, bbox_yi, bbox_yii, priority, inner_xi, inner_xii, inner_yi, inner_yii, created_at FROM layout_components 
WHERE layout_id = $1
ORDER BY created_at desc
`

func (q *Queries) GetLayoutComponentsByLayoutID(ctx context.Context, layoutID int32) ([]LayoutComponent, error) {
	rows, err := q.db.Query(ctx, getLayoutComponentsByLayoutID, layoutID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LayoutComponent
	for rows.Next() {
		var i LayoutComponent
		if err := rows.Scan(
			&i.ID,
			&i.LayoutID,
			&i.DesignID,
			&i.Width,
			&i.Height,
			&i.IsOriginal,
			&i.Color,
			&i.Type,
			&i.Xi,
			&i.Xii,
			&i.Yi,
			&i.Yii,
			&i.BboxXi,
			&i.BboxXii,
			&i.BboxYi,
			&i.BboxYii,
			&i.Priority,
			&i.InnerXi,
			&i.InnerXii,
			&i.InnerYi,
			&i.InnerYii,
			&i.CreatedAt,
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

const getLayoutElementsByLayoutID = `-- name: GetLayoutElementsByLayoutID :many
SELECT id, design_id, layout_id, component_id, asset_id, name, layer_id, text, xi, xii, yi, yii, inner_xi, inner_xii, inner_yi, inner_yii, width, height, is_group, group_id, level, kind, image_url, image_extension, created_at, updated_at FROM layout_elements
WHERE layout_id = $1
ORDER BY created_at desc
`

func (q *Queries) GetLayoutElementsByLayoutID(ctx context.Context, layoutID int32) ([]LayoutElement, error) {
	rows, err := q.db.Query(ctx, getLayoutElementsByLayoutID, layoutID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LayoutElement
	for rows.Next() {
		var i LayoutElement
		if err := rows.Scan(
			&i.ID,
			&i.DesignID,
			&i.LayoutID,
			&i.ComponentID,
			&i.AssetID,
			&i.Name,
			&i.LayerID,
			&i.Text,
			&i.Xi,
			&i.Xii,
			&i.Yi,
			&i.Yii,
			&i.InnerXi,
			&i.InnerXii,
			&i.InnerYi,
			&i.InnerYii,
			&i.Width,
			&i.Height,
			&i.IsGroup,
			&i.GroupID,
			&i.Level,
			&i.Kind,
			&i.ImageUrl,
			&i.ImageExtension,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getOriginalLayoutByDesignID = `-- name: GetOriginalLayoutByDesignID :one
SELECT id, design_id, is_original, image_url, width, height, data, stages, created_at, updated_at, deleted_at FROM layout 
WHERE design_id = $1 AND is_original = true
LIMIT 1
`

func (q *Queries) GetOriginalLayoutByDesignID(ctx context.Context, designID pgtype.Int4) (Layout, error) {
	row := q.db.QueryRow(ctx, getOriginalLayoutByDesignID, designID)
	var i Layout
	err := row.Scan(
		&i.ID,
		&i.DesignID,
		&i.IsOriginal,
		&i.ImageUrl,
		&i.Width,
		&i.Height,
		&i.Data,
		&i.Stages,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listLayouts = `-- name: ListLayouts :many
SELECT id, design_id, is_original, image_url, width, height, data, stages, created_at, updated_at, deleted_at FROM layout 
ORDER BY created_at desc
LIMIT $1 OFFSET $2
`

type ListLayoutsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListLayouts(ctx context.Context, arg ListLayoutsParams) ([]Layout, error) {
	rows, err := q.db.Query(ctx, listLayouts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Layout
	for rows.Next() {
		var i Layout
		if err := rows.Scan(
			&i.ID,
			&i.DesignID,
			&i.IsOriginal,
			&i.ImageUrl,
			&i.Width,
			&i.Height,
			&i.Data,
			&i.Stages,
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
