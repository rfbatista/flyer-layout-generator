package layoutgenerator

import (
	"algvisual/internal/application/usecases/designassets"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/shared"
	"context"

	"go.uber.org/zap"
)

type RemoveSimilarLayoutsFromJobUseCase struct {
	db   *database.Queries
	repo repositories.LayoutRepository
	das  *designassets.DesignAssetService
	log  *zap.Logger
}

func NewRemoveSimilarLayoutsFromJobUseCase(
	db *database.Queries,
	repo repositories.LayoutRepository,
	das *designassets.DesignAssetService,
	log *zap.Logger,
) (*RemoveSimilarLayoutsFromJobUseCase, error) {
	return &RemoveSimilarLayoutsFromJobUseCase{repo: repo, db: db, das: das, log: log}, nil
}

type RemoveSimilarLayoutsFromJobInput struct {
	JobID2 int64
}

type RemoveSimilarLayoutsFromJobOutput struct{}

func (u RemoveSimilarLayoutsFromJobUseCase) Execute(
	ctx context.Context,
	req RemoveSimilarLayoutsFromJobInput,
) (*RemoveSimilarLayoutsFromJobOutput, error) {
	layouts, err := u.repo.ListLayoutsByJob(ctx, req.JobID2)
	if err != nil {
		return nil, shared.NewInternalError("failed to list layouts by jobs")
	}
	var fullLayouts []entities.Layout
	for _, l := range layouts {
		var data GetLayoutByIDOutput
		data, err = GetLayoutByIDUseCase(ctx, u.db, GetLayoutByIDRequest{
			LayoutID: int32(l.ID),
		}, u.das)
		if err != nil {
			return nil, shared.NewInternalError("failed to get layou by id")
		}
		fullLayouts = append(fullLayouts, data.Layout)
	}
	alreadyDeleted := make(map[int32]bool)
	for _, base := range fullLayouts {
		for _, compareTo := range fullLayouts {
			if compareTo.ID == base.ID {
				continue
			}
			if alreadyDeleted[base.ID] {
				continue
			}
			isSimilar := IsSimilar(base, compareTo)
			if isSimilar {
				err = u.repo.SoftDeleteLayout(ctx, compareTo)
				if err != nil {
					u.log.Warn("failed to soft delete layout")
				} else {
					alreadyDeleted[compareTo.ID] = true
				}
			}
		}
	}
	return &RemoveSimilarLayoutsFromJobOutput{}, nil
}
