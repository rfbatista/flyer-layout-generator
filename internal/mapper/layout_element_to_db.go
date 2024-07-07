package mapper

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func DesignElementToDb(input entities.LayoutElement) database.LayoutElement {
	return database.LayoutElement{
		ID:             input.ID,
		DesignID:       input.DesignID,
		ComponentID:    pgtype.Int4{Int32: input.ComponentID, Valid: true},
		AssetID:        input.AssetID,
		Name:           pgtype.Text{String: input.Name, Valid: true},
		LayerID:        pgtype.Text{String: input.LayerID, Valid: true},
		Text:           pgtype.Text{String: input.Text, Valid: true},
		Xi:             pgtype.Int4{Int32: input.OuterContainer.UpperLeft.X, Valid: true},
		Xii:            pgtype.Int4{Int32: input.OuterContainer.DownRight.X, Valid: true},
		Yi:             pgtype.Int4{Int32: input.OuterContainer.UpperLeft.Y, Valid: true},
		Yii:            pgtype.Int4{Int32: input.OuterContainer.DownRight.Y, Valid: true},
		InnerXi:        pgtype.Int4{Int32: input.InnerContainer.UpperLeft.X, Valid: true},
		InnerXii:       pgtype.Int4{Int32: input.InnerContainer.DownRight.X, Valid: true},
		InnerYi:        pgtype.Int4{Int32: input.InnerContainer.UpperLeft.Y, Valid: true},
		InnerYii:       pgtype.Int4{Int32: input.InnerContainer.DownRight.Y, Valid: true},
		Width:          pgtype.Int4{Int32: input.OuterContainer.Width(), Valid: true},
		Height:         pgtype.Int4{Int32: input.OuterContainer.Height(), Valid: true},
		IsGroup:        pgtype.Bool{Bool: input.IsGroup, Valid: true},
		GroupID:        pgtype.Int4{Int32: input.GroupId, Valid: true},
		Level:          pgtype.Int4{Int32: input.Level, Valid: true},
		Kind:           pgtype.Text{String: input.Kind, Valid: true},
		ImageUrl:       pgtype.Text{String: input.ImageURL, Valid: true},
		ImageExtension: pgtype.Text{String: input.ImageExtension, Valid: true},
		CreatedAt:      pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt:      pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
}
