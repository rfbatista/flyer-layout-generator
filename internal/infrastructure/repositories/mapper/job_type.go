package mapper

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
)

func JobTypeToDatabase(t entities.JobType) database.JobType {
	switch t {
	case entities.JobTypeAdaptation:
		return database.JobTypeAdaptation
	case entities.JobTypeReplication:
		return database.JobTypeReplication
	}
	return database.JobTypeUnknown
}
