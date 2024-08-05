package mapper

import (
	"algvisual/database"
	"algvisual/internal/entities"
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
