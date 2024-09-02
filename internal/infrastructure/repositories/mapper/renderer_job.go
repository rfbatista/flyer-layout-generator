package mapper

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
)

func RendererJobToDomain(raw database.RendererJob) entities.RenderJob {
	ent := entities.RenderJob{
		ID:           raw.ID,
		LayoutID:     raw.LayoutID.Int32,
		AdaptationID: raw.AdaptationID.Int32,
	}
	if raw.StartedAt.Valid {
		ent.StartedAt = raw.StartedAt.Time
	}
	if raw.CreatedAt.Valid {
		ent.CreatedAt = raw.CreatedAt.Time
	}
	if raw.FinishedAt.Valid {
		ent.FinishedAt = raw.FinishedAt.Time
	}
	if raw.ErrorAt.Valid {
		ent.ErrorAt = raw.ErrorAt.Time
	}
	return ent
}

func RenderJobStatusToDatabase(d entities.RenderJobStatus) database.RendererJobStatus {
	switch d {
	case entities.RenderJobStatusError:
		return database.RendererJobStatusError
	case entities.RenderJobStatusFinished:
		return database.RendererJobStatusFinished
	case entities.RenderJobStatusPending:
		return database.RendererJobStatusPending
	case entities.RenderJobStatusStarted:
		return database.RendererJobStatusStarted
	}
	return database.RendererJobStatusUnknown
}
