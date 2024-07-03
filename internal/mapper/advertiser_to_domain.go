package mapper

import (
	"algvisual/database"
	"algvisual/internal/entities"
)

func AdvertiserToDomain(raw database.Advertiser) entities.Advertiser {
	return entities.Advertiser{
		ID:         int32(raw.ID),
		Name:       raw.Name,
		CreatedAt:  &raw.CreatedAt.Time,
		UpdatedAt:  &raw.UpdatedAt.Time,
		DeleteedAt: &raw.UpdatedAt.Time,
	}
}
