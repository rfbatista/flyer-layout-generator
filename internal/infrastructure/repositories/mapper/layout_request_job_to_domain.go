package mapper

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"encoding/json"
	"fmt"
)

func LayoutRequestJobToDomain(raw database.LayoutRequestsJob) entities.LayoutRequestJob {
	l := entities.LayoutRequestJob{
		ID:         int32(raw.ID),
		RequestID:  raw.RequestID.Int32,
		DesignID:   raw.DesignID.Int32,
		LayoutID:   raw.LayoutID.Int32,
		TemplateID: raw.TemplateID.Int32,
		Status:     raw.Status.String,
		Log:        raw.Log.String,
		ImageURL:   raw.ImageUrl.String,
	}
	if raw.Config.String != "" {
		var c entities.LayoutRequestConfig
		err := json.Unmarshal([]byte(raw.Config.String), &c)
		if err != nil {
			fmt.Println("falha ao realizar parser da config")
		} else {
			l.Config = &c
		}
	}
	if raw.StartedAt.Valid {
		l.StartedAt = &raw.StartedAt.Time
	}
	if raw.CreatedAt.Valid {
		l.CreatedAt = &raw.CreatedAt.Time
	}
	if raw.StoppedAt.Valid {
		l.StoppedAt = &raw.StoppedAt.Time
	}
	if raw.FinishedAt.Valid {
		l.FinishedAt = &raw.FinishedAt.Time
	}
	if raw.ErrorAt.Valid {
		l.ErrorAt = &raw.ErrorAt.Time
	}
	return l
}
