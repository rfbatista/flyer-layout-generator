package mapper

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"encoding/json"
)

func LayoutJobStatusToDomain(status database.LayoutJobStatus) entities.LayoutJobStatus {
	switch status {
	case database.LayoutJobStatusPending:
		return entities.LayoutJobStatusPending
	case database.LayoutJobStatusStarted:
		return entities.LayoutJobStatusStarted
	case database.LayoutJobStatusFinished:
		return entities.LayoutJobStatusFinished
	case database.LayoutJobStatusError:
		return entities.LayoutJobStatusError
	default:
		return entities.LayoutJobStatusUnknown
	}
}

func LayoutJobStatusToDatabase(status entities.LayoutJobStatus) database.LayoutJobStatus {
	switch status {
	case entities.LayoutJobStatusPending:
		return database.LayoutJobStatusPending
	case entities.LayoutJobStatusStarted:
		return database.LayoutJobStatusStarted
	case entities.LayoutJobStatusFinished:
		return database.LayoutJobStatusFinished
	case entities.LayoutJobStatusError:
		return database.LayoutJobStatusError
	default:
		return database.LayoutJobStatusUnknown
	}
}

func LayoutJobToDomain(job database.LayoutJob) entities.LayoutJob {
	var config entities.LayoutJobConfig
	json.Unmarshal([]byte(job.Config.String), &config)
	return entities.LayoutJob{
		ID:              job.ID,
		BasedOnLayoutID: job.BasedOnLayoutID.Int32,
		TemplateID:      job.TemplateID.Int32,
		UserID:          job.UserID.Int32,
		AdaptationID:    job.AdaptationBatchID.Int32,
		Status:          LayoutJobStatusToDomain(job.Status),
		StartedAt:       job.StartedAt.Time,
		FinishedAt:      job.FinishedAt.Time,
		ErrorAt:         job.ErrorAt.Time,
		CreatedAt:       job.CreatedAt.Time,
		Log:             job.Log.String,
		Config:          config,
	}
}
