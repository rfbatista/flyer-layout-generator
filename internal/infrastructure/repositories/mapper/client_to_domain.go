package mapper

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
)

func ClientToDomain(raw database.Client) entities.Client {
	return entities.Client{
		ID:         int64(raw.ID),
		Name:       raw.Name,
		CreatedAt:  &raw.CreatedAt.Time,
		UpdatedAt:  &raw.UpdatedAt.Time,
		DeleteedAt: &raw.DeletedAt.Time,
	}
}
