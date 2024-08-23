package mapper

import (
	"algvisual/database"
	"algvisual/internal/entities"
)

func AdaptationBatchStatusToDatabase(
	a entities.AdaptationBatchStatus,
) database.AdaptationBatchStatus {
	switch a {
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
