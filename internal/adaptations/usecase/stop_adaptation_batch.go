package usecase

import (
	"algvisual/internal/adaptations/errors"
	"algvisual/internal/adaptations/repositories"
	"algvisual/internal/entities"
	"algvisual/internal/shared"
	"context"
)

type StopAdaptationBatchInput struct {
	Session entities.UserSession
}

type StopAdaptationBatchOutput struct {
	Data []entities.AdaptationBatch
}

func StopAdaptationBatchUseCase(
	ctx context.Context,
	req StopAdaptationBatchInput,
	repo *repositories.AdaptationBatchRepository,
) (*StopAdaptationBatchOutput, error) {
	session := req.Session
	batches, err := repo.CancelActiveAdaptations(ctx, session.UserID)
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
