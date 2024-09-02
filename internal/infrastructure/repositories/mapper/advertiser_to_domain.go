package mapper

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
)

func AdvertiserToDomain(raw database.Advertiser) entities.Advertiser {
	return entities.Advertiser{
		ID:         raw.ID,
		Name:       raw.Name,
		CreatedAt:  &raw.CreatedAt.Time,
		UpdatedAt:  &raw.UpdatedAt.Time,
		DeleteedAt: &raw.UpdatedAt.Time,
	}
}
