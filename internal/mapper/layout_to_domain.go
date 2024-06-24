package mapper

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
)

func LayoutToDomain(raw database.Layout) entities.Layout {
	return entities.Layout{
		ID:       int32(raw.ID),
		Width:    raw.Width.Int32,
		Height:   raw.Height.Int32,
		ImageURL: raw.ImageUrl.String,
		DesignID: raw.DesignID.Int32,
	}
}
