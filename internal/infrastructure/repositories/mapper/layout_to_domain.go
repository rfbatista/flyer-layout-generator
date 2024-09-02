package mapper

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
)

func LayoutToDomain(raw database.Layout) entities.Layout {
	return entities.Layout{
		ID:         int32(raw.ID),
		Width:      raw.Width.Int32,
		TemplateID: raw.TemplateID.Int32,
		Height:     raw.Height.Int32,
		ImageURL:   raw.ImageUrl.String,
		Stages:     raw.Stages,
		DesignID:   raw.DesignID.Int32,
	}
}
