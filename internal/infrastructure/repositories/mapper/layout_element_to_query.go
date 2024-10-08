package mapper

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"

	"github.com/jackc/pgx/v5/pgtype"
)

func LayoutElementToCreateElement(
	i entities.LayoutElement,
	LayoutID, designID, componentID int32,
) database.CreateElementParams {
	dbelem := DesignElementToDb(i)
	return database.CreateElementParams{
		DesignID:       designID,
		LayoutID:       LayoutID,
		AssetID:        i.AssetID,
		ComponentID:    pgtype.Int4{Int32: componentID, Valid: true},
		LayerID:        dbelem.LayerID,
		Name:           dbelem.Name,
		Text:           dbelem.Text,
		Xi:             dbelem.Xi,
		Yi:             dbelem.Yi,
		Xii:            dbelem.Xii,
		Yii:            dbelem.Yii,
		InnerXi:        dbelem.InnerXi,
		InnerXii:       dbelem.InnerXii,
		InnerYi:        dbelem.InnerYi,
		InnerYii:       dbelem.InnerYii,
		Kind:           dbelem.Kind,
		IsGroup:        dbelem.IsGroup,
		GroupID:        dbelem.GroupID,
		Level:          dbelem.Level,
		ImageUrl:       dbelem.ImageUrl,
		Width:          dbelem.Width,
		Height:         dbelem.Height,
		ImageExtension: dbelem.ImageExtension,
	}
}
