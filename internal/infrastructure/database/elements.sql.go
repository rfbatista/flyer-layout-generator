// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: elements.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getDesignElementsByComponentID = `-- name: GetDesignElementsByComponentID :many
select id, design_id, layout_id, component_id, asset_id, name, layer_id, text, xi, xii, yi, yii, inner_xi, inner_xii, inner_yi, inner_yii, width, height, is_group, group_id, level, kind, image_url, image_extension, created_at, updated_at from layout_elements 
where component_id = $1
`

func (q *Queries) GetDesignElementsByComponentID(ctx context.Context, componentID pgtype.Int4) ([]LayoutElement, error) {
	rows, err := q.db.Query(ctx, getDesignElementsByComponentID, componentID)
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

const getElements = `-- name: GetElements :many
SELECT id, design_id, layout_id, component_id, asset_id, name, layer_id, text, xi, xii, yi, yii, inner_xi, inner_xii, inner_yi, inner_yii, width, height, is_group, group_id, level, kind, image_url, image_extension, created_at, updated_at FROM layout_elements 
WHERE design_id = $1
`

func (q *Queries) GetElements(ctx context.Context, designID int32) ([]LayoutElement, error) {
	rows, err := q.db.Query(ctx, getElements, designID)
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

const getElementsByLayoutID = `-- name: GetElementsByLayoutID :many
SELECT id, design_id, layout_id, component_id, asset_id, name, layer_id, text, xi, xii, yi, yii, inner_xi, inner_xii, inner_yi, inner_yii, width, height, is_group, group_id, level, kind, image_url, image_extension, created_at, updated_at FROM layout_elements 
WHERE layout_id = $1
`

func (q *Queries) GetElementsByLayoutID(ctx context.Context, layoutID int32) ([]LayoutElement, error) {
	rows, err := q.db.Query(ctx, getElementsByLayoutID, layoutID)
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

const getLayoutElementByID = `-- name: GetLayoutElementByID :one
SELECT id, design_id, layout_id, component_id, asset_id, name, layer_id, text, xi, xii, yi, yii, inner_xi, inner_xii, inner_yi, inner_yii, width, height, is_group, group_id, level, kind, image_url, image_extension, created_at, updated_at FROM layout_elements 
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetLayoutElementByID(ctx context.Context, id int32) (LayoutElement, error) {
	row := q.db.QueryRow(ctx, getLayoutElementByID, id)
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

const getdesignElements = `-- name: GetdesignElements :many
SELECT id, design_id, layout_id, component_id, asset_id, name, layer_id, text, xi, xii, yi, yii, inner_xi, inner_xii, inner_yi, inner_yii, width, height, is_group, group_id, level, kind, image_url, image_extension, created_at, updated_at FROM layout_elements 
WHERE design_id = $1
`

func (q *Queries) GetdesignElements(ctx context.Context, designID int32) ([]LayoutElement, error) {
	rows, err := q.db.Query(ctx, getdesignElements, designID)
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

const getdesignElementsByIDlist = `-- name: GetdesignElementsByIDlist :many
select id, design_id, layout_id, component_id, asset_id, name, layer_id, text, xi, xii, yi, yii, inner_xi, inner_xii, inner_yi, inner_yii, width, height, is_group, group_id, level, kind, image_url, image_extension, created_at, updated_at from layout_elements 
where id = any ($1)
`

func (q *Queries) GetdesignElementsByIDlist(ctx context.Context, ids []int32) ([]LayoutElement, error) {
	rows, err := q.db.Query(ctx, getdesignElementsByIDlist, ids)
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

const updateLayoutElementPosition = `-- name: UpdateLayoutElementPosition :one
UPDATE layout_elements
SET 
    xi              = $2,
    xii             = $3,
    yi              = $4,
    yii             = $5,
    inner_xi              = $6,
    inner_xii             = $7,
    inner_yi              = $8,
    inner_yii             = $9
WHERE
    id = $1
RETURNING id, design_id, layout_id, component_id, asset_id, name, layer_id, text, xi, xii, yi, yii, inner_xi, inner_xii, inner_yi, inner_yii, width, height, is_group, group_id, level, kind, image_url, image_extension, created_at, updated_at
`

type UpdateLayoutElementPositionParams struct {
	ID       int32       `json:"id"`
	Xi       pgtype.Int4 `json:"xi"`
	Xii      pgtype.Int4 `json:"xii"`
	Yi       pgtype.Int4 `json:"yi"`
	Yii      pgtype.Int4 `json:"yii"`
	InnerXi  pgtype.Int4 `json:"inner_xi"`
	InnerXii pgtype.Int4 `json:"inner_xii"`
	InnerYi  pgtype.Int4 `json:"inner_yi"`
	InnerYii pgtype.Int4 `json:"inner_yii"`
}

func (q *Queries) UpdateLayoutElementPosition(ctx context.Context, arg UpdateLayoutElementPositionParams) (LayoutElement, error) {
	row := q.db.QueryRow(ctx, updateLayoutElementPosition,
		arg.ID,
		arg.Xi,
		arg.Xii,
		arg.Yi,
		arg.Yii,
		arg.InnerXi,
		arg.InnerXii,
		arg.InnerYi,
		arg.InnerYii,
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

const updateLayoutElementSize = `-- name: UpdateLayoutElementSize :one
UPDATE layout_elements
SET 
    xi              = $2,
    xii             = $3,
    yi              = $4,
    yii             = $5,
    inner_xi              = $6,
    inner_xii             = $7,
    inner_yi              = $8,
    inner_yii             = $9,
    width                 = $10,
    height                = $11
WHERE
    id = $1
RETURNING id, design_id, layout_id, component_id, asset_id, name, layer_id, text, xi, xii, yi, yii, inner_xi, inner_xii, inner_yi, inner_yii, width, height, is_group, group_id, level, kind, image_url, image_extension, created_at, updated_at
`

type UpdateLayoutElementSizeParams struct {
	ID       int32       `json:"id"`
	Xi       pgtype.Int4 `json:"xi"`
	Xii      pgtype.Int4 `json:"xii"`
	Yi       pgtype.Int4 `json:"yi"`
	Yii      pgtype.Int4 `json:"yii"`
	InnerXi  pgtype.Int4 `json:"inner_xi"`
	InnerXii pgtype.Int4 `json:"inner_xii"`
	InnerYi  pgtype.Int4 `json:"inner_yi"`
	InnerYii pgtype.Int4 `json:"inner_yii"`
	Width    pgtype.Int4 `json:"width"`
	Height   pgtype.Int4 `json:"height"`
}

func (q *Queries) UpdateLayoutElementSize(ctx context.Context, arg UpdateLayoutElementSizeParams) (LayoutElement, error) {
	row := q.db.QueryRow(ctx, updateLayoutElementSize,
		arg.ID,
		arg.Xi,
		arg.Xii,
		arg.Yi,
		arg.Yii,
		arg.InnerXi,
		arg.InnerXii,
		arg.InnerYi,
		arg.InnerYii,
		arg.Width,
		arg.Height,
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
