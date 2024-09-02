package mapper

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"

	"github.com/jackc/pgx/v5/pgtype"
)

func LayoutComponentFromDomain(c entities.LayoutComponent) database.LayoutComponent {
	return database.LayoutComponent{
		DesignID: c.DesignID,
		Width:    pgtype.Int4{Int32: c.OuterContainer.Width(), Valid: true},
		Height:   pgtype.Int4{Int32: c.OuterContainer.Height(), Valid: true},
		Type: database.NullComponentType{
			ComponentType: entities.ComponentTypeToDatabaseComponentType(entities.StringToComponentType(c.Type)),
			Valid:         true,
		},
		Color: pgtype.Text{String: c.Color, Valid: true},
		Xi:    pgtype.Int4{Int32: c.OuterContainer.UpperLeft.X, Valid: true},
		Xii:   pgtype.Int4{Int32: c.OuterContainer.DownRight.X, Valid: true},
		Yi:    pgtype.Int4{Int32: c.OuterContainer.UpperLeft.Y, Valid: true},
		Yii:   pgtype.Int4{Int32: c.OuterContainer.DownRight.Y, Valid: true},
	}
}
