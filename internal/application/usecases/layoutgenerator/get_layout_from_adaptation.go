package layoutgenerator

import (
	"algvisual/internal/application/errors"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/shared"
	"context"

	"go.uber.org/multierr"
)

type GetLayoutByJobsUseCase struct {
	repo repositories.LayoutRepository
}

func NewGetLayoutByJobUseCase(
	repo repositories.LayoutRepository,
) (*GetLayoutByJobsUseCase, error) {
	return &GetLayoutByJobsUseCase{repo: repo}, nil
}

type GetLayoutByJobsInput struct {
	AdaptationID int64 `json:"adaptation_id,omitempty" param:"adaptation_id"`
}

type GetLayoutByJobsOutput struct {
	Data []entities.Layout `json:"data,omitempty"`
}

func (u GetLayoutByJobsUseCase) Execute(
	ctx context.Context,
	req GetLayoutByJobsInput,
) (*GetLayoutByJobsOutput, error) {
	listOfLayouts, err := u.repo.ListLayoutsByAdaptation(ctx, req.AdaptationID)
	if err != nil {
		return nil, multierr.Append(
			err,
			shared.NewError(errors.NO_ADAPTATION_FOUND, "couldnt find a the list of layouts", ""),
		)
	}
	return &GetLayoutByJobsOutput{
		Data: listOfLayouts,
	}, nil
}
