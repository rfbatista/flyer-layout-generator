package layoutgenerator

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/repositories"
	"context"
)

type DeleteLayoutByIdInput struct {
	LayoutID int64 `json:"layout_id,omitempty" param:"layout_id"`
}

type DeleteLayoutByIdOutput struct {
	Data entities.Layout `json:"data,omitempty"`
}

func DeleteLayoutByIdUseCase(
	ctx context.Context,
	req DeleteLayoutByIdInput,
	repo repositories.LayoutRepository,
) (*DeleteLayoutByIdOutput, error) {
	l, err := repo.GetLayoutByID(ctx, req.LayoutID)
	if err != nil {
		return nil, err
	}
	err = repo.DeleteLayout(ctx, *l)
	if err != nil {
		return nil, err
	}
	return &DeleteLayoutByIdOutput{Data: *l}, nil
}
