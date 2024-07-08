package mapper

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"encoding/json"
	"time"
)

func LayoutRequestToDomain(raw database.LayoutRequest) entities.LayoutRequest {
	var createdAt *time.Time
	var startedAt *time.Time
	var stoppedAt *time.Time
	var errorAt *time.Time
	var config entities.LayoutRequestConfig
	if raw.CreatedAt.Valid {
		createdAt = &raw.CreatedAt.Time
	}
	if raw.Config.Valid {
		json.Unmarshal([]byte(raw.Config.String), &config)
	}
	return entities.LayoutRequest{
		ID:        int32(raw.ID),
		DesignID:  raw.DesignID.Int32,
		CreatedAt: createdAt,
		Total:     raw.Total.Int32,
		Done:      raw.Done,
		Config:    config,
		StartedAt: startedAt,
		StoppedAt: stoppedAt,
		ErrorAt:   errorAt,
	}
}
