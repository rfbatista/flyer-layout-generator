package layoutgenerator

import (
	"algvisual/internal/application/errors"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/shared"
	"context"

	"go.uber.org/multierr"
)

type GetLayoutFromAdaptationUseCase struct {
	repo repositories.LayoutRepository
}

func NewGetLayoutFromAdaptationUseCase(
	repo repositories.LayoutRepository,
) (*GetLayoutFromAdaptationUseCase, error) {
	return &GetLayoutFromAdaptationUseCase{repo: repo}, nil
}

type GetLayoutFromAdaptationInput struct {
	AdaptationID int64 `json:"adaptation_id,omitempty" param:"adaptation_id"`
}

type GetLayoutFromAdaptationOutput struct {
	Data []entities.Layout `json:"data,omitempty"`
}

func (u GetLayoutFromAdaptationUseCase) Execute(
	ctx context.Context,
	req GetLayoutFromAdaptationInput,
) (*GetLayoutFromAdaptationOutput, error) {
	listOfLayouts, err := u.repo.ListLayoutsByAdaptation(ctx, req.AdaptationID)
	if err != nil {
		return nil, multierr.Append(
			err,
			shared.NewError(errors.NO_ADAPTATION_FOUND, "couldnt find a the list of layouts", ""),
		)
	}
	return &GetLayoutFromAdaptationOutput{
		Data: listOfLayouts,
	}, nil
}
