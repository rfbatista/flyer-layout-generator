package mapper

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"

	"github.com/jackc/pgx/v5/pgtype"
)

func LayoutComponentFromDomain(c entities.DesignComponent) database.LayoutComponent {
	return database.LayoutComponent{
		DesignID: c.DesignID,
		Width:    pgtype.Int4{Int32: c.Width, Valid: true},
		Height:   pgtype.Int4{Int32: c.Height, Valid: true},
		Xi:       pgtype.Int4{Int32: c.Xi, Valid: true},
		Xii:      pgtype.Int4{Int32: c.Xii, Valid: true},
		Yi:       pgtype.Int4{Int32: c.Yi, Valid: true},
		Yii:      pgtype.Int4{Int32: c.Yii, Valid: true},
	}
}

func LayoutRegionFromDomain(c entities.Cell) database.LayoutRegion {
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
