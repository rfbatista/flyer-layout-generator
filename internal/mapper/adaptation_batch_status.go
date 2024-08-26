package mapper

import (
	"algvisual/database"
	"algvisual/internal/entities"
)

func AdaptationBatchStatusToDatabase(
	a entities.AdaptationBatchStatus,
) database.AdaptationBatchStatus {
	switch a {
	case entities.AdaptationBatchStatusCanceled:
		return database.AdaptationBatchStatusCanceled
	case entities.AdaptationBatchStatusError:
		return database.AdaptationBatchStatusError
	case entities.AdaptationBatchStatusFinished:
		return database.AdaptationBatchStatusFinished
	case entities.AdaptationBatchStatusPending:
		return database.AdaptationBatchStatusPending
	case entities.AdaptationBatchStatusStarted:
		return database.AdaptationBatchStatusStarted
	}
	return database.AdaptationBatchStatusUnknown
}

func AdaptationBatchStatusFromDatabase(
	d database.AdaptationBatchStatus,
) entities.AdaptationBatchStatus {
	switch d {
	case database.AdaptationBatchStatusCanceled:
		return entities.AdaptationBatchStatusCanceled
	case database.AdaptationBatchStatusError:
		return entities.AdaptationBatchStatusError
	case database.AdaptationBatchStatusFinished:
		return entities.AdaptationBatchStatusFinished
	case database.AdaptationBatchStatusPending:
		return entities.AdaptationBatchStatusPending
	case database.AdaptationBatchStatusStarted:
		return entities.AdaptationBatchStatusStarted
	}
	return entities.AdaptationBatchStatusUnknown
}
