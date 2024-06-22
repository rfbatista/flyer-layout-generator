package mapper

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"

	"github.com/jackc/pgx/v5/pgtype"
)

func LayoutComponentFromDomain(c entities.DesignComponent) database.LayoutComponent {
	return database.LayoutComponent{
		DesignID: c.DesignID,
		Width:    pgtype.Int4{Int32: c.OuterContainer.Width(), Valid: true},
		Height:   pgtype.Int4{Int32: c.OuterContainer.Height(), Valid: true},
		Type:     pgtype.Text{String: c.Type, Valid: true},
		Color:    pgtype.Text{String: c.Color, Valid: true},
		Xi:       pgtype.Int4{Int32: c.OuterContainer.UpperLeft.X, Valid: true},
		Xii:      pgtype.Int4{Int32: c.OuterContainer.DownRight.X, Valid: true},
		Yi:       pgtype.Int4{Int32: c.OuterContainer.UpperLeft.Y, Valid: true},
		Yii:      pgtype.Int4{Int32: c.OuterContainer.DownRight.Y, Valid: true},
	}
}

func LayoutRegionFromDomain(c entities.GridCell) database.LayoutRegion {
	return database.LayoutRegion{}
}

func LayoutTemplateFromDomain(c entities.Template) database.LayoutTemplate {
	return database.LayoutTemplate{
		Width:  pgtype.Int4{Int32: c.Width, Valid: true},
		Height: pgtype.Int4{Int32: c.Height, Valid: true},
		SlotsX: pgtype.Int4{Int32: c.SlotsX, Valid: true},
		SlotsY: pgtype.Int4{Int32: c.SlotsY, Valid: true},
	}
}
