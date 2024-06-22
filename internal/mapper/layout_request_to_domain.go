package mapper

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"time"
)

func LayoutRequestToDomain(raw database.LayoutRequest) entities.LayoutRequest {
	var createdAt *time.Time
	var startedAt *time.Time
	var stoppedAt *time.Time
	var errorAt *time.Time
	if raw.CreatedAt.Valid {
		createdAt = &raw.CreatedAt.Time
	}
	return entities.LayoutRequest{
		ID:        int32(raw.ID),
		DesignID:  raw.DesignID.Int32,
		CreatedAt: createdAt,
		StartedAt: startedAt,
		StoppedAt: stoppedAt,
		ErrorAt:   errorAt,
	}
}
