package adaptations

import (
	"algvisual/internal/application/errors"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/shared"
	"context"
)

type StopAdaptationBatchInput struct {
	Session entities.UserSession
}

type StopAdaptationBatchOutput struct {
	Data []entities.Job
}

func StopAdaptationBatchUseCase(
	ctx context.Context,
	req StopAdaptationBatchInput,
	repo *repositories.JobRepository,
) (*StopAdaptationBatchOutput, error) {
	session := req.Session
	batches, err := repo.CancelActiveAdaptations(ctx, session.UserID, entities.JobTypeAdaptation)
	if err != nil {
		return nil, shared.NewError(
			errors.CANT_CANCEL_ADAPTATIONS,
			"falha ao cancelar adaptações",
			err.Error(),
		)
	}
	return &StopAdaptationBatchOutput{
		Data: batches,
	}, nil
}
