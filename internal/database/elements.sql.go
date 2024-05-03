// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: elements.sql

package database

import (
	"context"
)

const getElements = `-- name: GetElements :many
SELECT id, photoshop_id, name, layer_id, text, xi, xii, yi, yii, width, height, is_group, group_id, level, kind, component_id, image_url, image_extension, created_at, updated_at FROM photoshop_element 
WHERE photoshop_id = $1
`

func (q *Queries) GetElements(ctx context.Context, photoshopID int32) ([]PhotoshopElement, error) {
	rows, err := q.db.Query(ctx, getElements, photoshopID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PhotoshopElement
	for rows.Next() {
		var i PhotoshopElement
		if err := rows.Scan(
			&i.ID,
			&i.PhotoshopID,
			&i.Name,
			&i.LayerID,
			&i.Text,
			&i.Xi,
			&i.Xii,
			&i.Yi,
			&i.Yii,
			&i.Width,
			&i.Height,
			&i.IsGroup,
			&i.GroupID,
			&i.Level,
			&i.Kind,
			&i.ComponentID,
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

const getPhotoshopElements = `-- name: GetPhotoshopElements :many
SELECT id, photoshop_id, name, layer_id, text, xi, xii, yi, yii, width, height, is_group, group_id, level, kind, component_id, image_url, image_extension, created_at, updated_at FROM photoshop_element 
WHERE photoshop_id = $1
`

func (q *Queries) GetPhotoshopElements(ctx context.Context, photoshopID int32) ([]PhotoshopElement, error) {
	rows, err := q.db.Query(ctx, getPhotoshopElements, photoshopID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PhotoshopElement
	for rows.Next() {
		var i PhotoshopElement
		if err := rows.Scan(
			&i.ID,
			&i.PhotoshopID,
			&i.Name,
			&i.LayerID,
			&i.Text,
			&i.Xi,
			&i.Xii,
			&i.Yi,
			&i.Yii,
			&i.Width,
			&i.Height,
			&i.IsGroup,
			&i.GroupID,
			&i.Level,
			&i.Kind,
			&i.ComponentID,
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

const getphotoshopElementsByIDlist = `-- name: GetphotoshopElementsByIDlist :many
select id, photoshop_id, name, layer_id, text, xi, xii, yi, yii, width, height, is_group, group_id, level, kind, component_id, image_url, image_extension, created_at, updated_at from photoshop_element 
where id = any ($1)
`

func (q *Queries) GetphotoshopElementsByIDlist(ctx context.Context, ids []int32) ([]PhotoshopElement, error) {
	rows, err := q.db.Query(ctx, getphotoshopElementsByIDlist, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PhotoshopElement
	for rows.Next() {
		var i PhotoshopElement
		if err := rows.Scan(
			&i.ID,
			&i.PhotoshopID,
			&i.Name,
			&i.LayerID,
			&i.Text,
			&i.Xi,
			&i.Xii,
			&i.Yi,
			&i.Yii,
			&i.Width,
			&i.Height,
			&i.IsGroup,
			&i.GroupID,
			&i.Level,
			&i.Kind,
			&i.ComponentID,
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
