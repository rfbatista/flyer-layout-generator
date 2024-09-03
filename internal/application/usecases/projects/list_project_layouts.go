package projects

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/shared"
	"context"

	"go.uber.org/multierr"
)

type ListProjectLayoutsUseCase struct {
	repo repositories.LayoutRepository
}

func NewListProjectLayoutsUseCase(
	repo repositories.LayoutRepository,
) (*ListProjectLayoutsUseCase, error) {
	return &ListProjectLayoutsUseCase{
		repo: repo,
	}, nil
}

type ListProjectLayoutsInput struct {
	ProjectID int32
}

type ListProjectLayoutsOutput struct {
	Data []entities.Layout `json:"data,omitempty"`
}

func (u ListProjectLayoutsUseCase) Execute(
	ctx context.Context,
	req ListProjectLayoutsInput,
) (*ListProjectLayoutsOutput, error) {
	layouts, err := u.repo.ListLayoutsByProjectID(ctx, req.ProjectID)
	if err != nil {
		return nil, multierr.Append(err, shared.NewInternalError("failed to list project layouts"))
	}
	return &ListProjectLayoutsOutput{
		Data: layouts,
	}, nil
}
