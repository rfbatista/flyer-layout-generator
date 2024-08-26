package usecase

import (
	"algvisual/internal/adaptations/repositories"
	"algvisual/internal/entities"
	"context"
)

type GetActiveAdaptationBatchInput struct {
	Session entities.UserSession
}

type GetActiveAdaptationBatchOutput struct {
	Data *entities.AdaptationBatch
}

func GetActiveAdaptationBatchUseCase(
	ctx context.Context,
	req GetActiveAdaptationBatchInput,
	repo *repositories.AdaptationBatchRepository,
) (*GetActiveAdaptationBatchOutput, error) {
	adap, err := repo.GetByUser(
		ctx,
		req.Session.UserID,
		repositories.AdaptationBatchRepositoryGetByUserParams{
			DoByStatus: true,
			Status: []entities.AdaptationBatchStatus{
				entities.AdaptationBatchStatusPending,
				entities.AdaptationBatchStatusStarted,
				entities.AdaptationBatchStatusRenderingImages,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	return &GetActiveAdaptationBatchOutput{
		Data: adap,
	}, nil
}
