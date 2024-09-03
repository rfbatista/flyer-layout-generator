package replications

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/shared"
	"context"

	"go.uber.org/multierr"
)

type GetActiveReplicationUseCase struct {
	repo *repositories.JobRepository
}

func NewGetActiveReplicationUseCase(
	repo *repositories.JobRepository,
) (*GetActiveReplicationUseCase, error) {
	if repo == nil {
		return nil, shared.NewInternalError("missing adaptation repository")
	}
	return &GetActiveReplicationUseCase{repo: repo}, nil
}

type GetActiveReplicationInput struct {
	Session entities.UserSession
}

type GetActiveReplicationOutput struct {
	Data *entities.Job `json:"data"`
}

func (u GetActiveReplicationUseCase) Execute(
	ctx context.Context,
	req GetActiveReplicationInput,
) (*GetActiveReplicationOutput, error) {
	adap, err := u.repo.GetByUser(
		ctx,
		req.Session.UserID,
		repositories.JobRepositoryGetByUserParams{
			Type:                    entities.JobTypeReplication,
			FilterByPending:         true,
			FilterByRenderingImages: true,
			FilterByStarted:         true,
		},
	)
	if err != nil {
		return nil, err
	}
	sum, err := u.repo.GetSummary(ctx, int32(adap.ID))
	if err != nil {
		return nil, err
	}
	if sum != nil {
		adap.Summary = *sum
	}
	if adap.Summary.Total != 0 && adap.Summary.Total == adap.Summary.Done {
		adap.Status = entities.AdaptationBatchStatusFinished
		adap, err = u.repo.Update(ctx, *adap, repositories.JobRepositoryUpdateParams{
			UpdateStatus: true,
		})
		if err != nil {
			return nil, multierr.Append(err, shared.NewInternalError("failed to update adaptation"))
		}
		adap.Summary = *sum
	}
	return &GetActiveReplicationOutput{
		Data: adap,
	}, nil
}
