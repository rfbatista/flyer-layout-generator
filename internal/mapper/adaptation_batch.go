package mapper

import (
	"algvisual/database"
	"algvisual/internal/entities"
)

func AdaptationBatchToDomain(raw database.AdaptationBatch) entities.AdaptationBatch {
	adap := entities.AdaptationBatch{
		ID:        raw.ID,
		LayoutID:  raw.LayoutID.Int32,
		RequestID: raw.RequestID.Int32,
		UserID:    int64(raw.UserID.Int32),
		Log:       raw.Log.String,
		Status:    AdaptationBatchStatusFromDatabase(raw.Status.AdaptationBatchStatus),
	}
	if raw.StartedAt.Valid {
		adap.StartedAt = raw.StartedAt.Time
	}
	if raw.CreatedAt.Valid {
		adap.CreatedAt = raw.CreatedAt.Time
	}
	if raw.StoppedAt.Valid {
		adap.StoppedAt = raw.StoppedAt.Time
	}
	if raw.FinishedAt.Valid {
		adap.FinishedAt = raw.FinishedAt.Time
	}
	if raw.ErrorAt.Valid {
		adap.ErrorAt = raw.ErrorAt.Time
	}
	return adap
}
