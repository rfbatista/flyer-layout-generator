package adaptations

import (
	"algvisual/internal/application/usecases/layoutgenerator"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/shared"
	"context"

	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type GetActiveAdaptationBatchUseCase struct {
	log    *zap.Logger
	repo   *repositories.JobRepository
	remove *layoutgenerator.RemoveSimilarLayoutsFromJobUseCase
}

func NewGetActiveAdaptationBatchUseCase(
	log *zap.Logger,
	repo *repositories.JobRepository,
	remove *layoutgenerator.RemoveSimilarLayoutsFromJobUseCase,
) (*GetActiveAdaptationBatchUseCase, error) {
	if repo == nil {
		return nil, shared.NewInternalError("missing adaptation repository")
	}
	return &GetActiveAdaptationBatchUseCase{repo: repo, log: log, remove: remove}, nil
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
		g.log.Error("failed to get jobs", zap.Error(err))
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
		if !adap.RemovedSimilars {
			_, err = g.remove.Execute(ctx, layoutgenerator.RemoveSimilarLayoutsFromJobInput{
				JobID2: adap.ID,
			})
			if err != nil {
				return nil, multierr.Append(
					err,
					shared.NewInternalError("failed to remove duplications"),
				)
			}
			_, err = g.repo.Update(ctx, *adap, repositories.JobRepositoryUpdateParams{
				UpdateCleanedDuplicates: true,
			})
			if err != nil {
				return nil, multierr.Append(
					err,
					shared.NewInternalError("failed to update batch after removed duplicates"),
				)
			}
		}
		adap.Summary = *sum
	}
	return &GetActiveAdaptationBatchOutput{
		Data: adap,
	}, nil
}
