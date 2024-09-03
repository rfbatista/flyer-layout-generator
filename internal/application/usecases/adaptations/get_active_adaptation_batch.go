package adaptations

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/shared"
	"context"

	"go.uber.org/multierr"
)

type GetActiveAdaptationBatchUseCase struct {
	repo *repositories.JobRepository
}

func NewGetActiveAdaptationBatchUseCase(
	repo *repositories.JobRepository,
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
	Data *entities.Job `json:"data"`
}

func (g GetActiveAdaptationBatchUseCase) Execute(
	ctx context.Context,
	req GetActiveAdaptationBatchInput,
) (*GetActiveAdaptationBatchOutput, error) {
	adap, err := g.repo.GetByUser(
		ctx,
		req.Session.UserID,
		repositories.JobRepositoryGetByUserParams{
			Type:                    entities.JobTypeAdaptation,
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
		adap, err = g.repo.Update(ctx, *adap, repositories.JobRepositoryUpdateParams{
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
