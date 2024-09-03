package mapper

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
)

func AdaptationBatchStatusToDatabase(
	a entities.AdaptationBatchStatus,
) database.AdaptationBatchStatus {
	switch a {
	case entities.AdaptationBatchStatusCanceled:
		return database.AdaptationBatchStatusCanceled
	case entities.AdaptationBatchStatusClosed:
		return database.AdaptationBatchStatusClosed
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
	case database.AdaptationBatchStatusClosed:
		return entities.AdaptationBatchStatusClosed
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
