package adaptations

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/shared"
	"context"

	"go.uber.org/multierr"
)

type GetActiveAdaptationBatchUseCase struct {
	repo *repositories.AdaptationBatchRepository
}

func NewGetActiveAdaptationBatchUseCase(
	repo *repositories.AdaptationBatchRepository,
) (*GetActiveAdaptationBatchUseCase, error) {
	if repo == nil {
		return nil, shared.NewInternalError("missing adaptation repository")
	}
	return &GetActiveAdaptationBatchUseCase{repo: repo}, nil
}

type GetActiveAdaptationBatchInput struct {
	Session entities.UserSession
}

type GetActiveAdaptationBatchOutput struct {
	Data *entities.AdaptationBatch `json:"data"`
}

func (g GetActiveAdaptationBatchUseCase) Execute(
	ctx context.Context,
	req GetActiveAdaptationBatchInput,
) (*GetActiveAdaptationBatchOutput, error) {
	adap, err := g.repo.GetByUser(
		ctx,
		req.Session.UserID,
		repositories.AdaptationBatchRepositoryGetByUserParams{
			FilterByPending:         true,
			FilterByRenderingImages: true,
			FilterByStarted:         true,
		},
	)
	if err != nil {
		return nil, err
	}
	sum, err := g.repo.GetSummary(ctx, int32(adap.ID))
	if err != nil {
		return nil, err
	}
	if sum != nil {
		adap.Summary = *sum
	}
	if adap.Summary.Total != 0 && adap.Summary.Total == adap.Summary.Done {
		adap.Status = entities.AdaptationBatchStatusFinished
		adap, err = g.repo.Update(ctx, *adap, repositories.AdaptationBatchRepositoryUpdateParams{
			UpdateStatus: true,
		})
		if err != nil {
			return nil, multierr.Append(err, shared.NewInternalError("failed to update adaptation"))
		}
		adap.Summary = *sum
	}
	return &GetActiveAdaptationBatchOutput{
		Data: adap,
	}, nil
}
