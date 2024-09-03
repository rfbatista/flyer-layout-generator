package projects

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/shared"
	"context"

	"go.uber.org/multierr"
)

type SaveProjectLayoutUseCase struct {
	repo repositories.LayoutRepository
}

func NewSaveProjectLayoutUseCase(
	repo repositories.LayoutRepository,
) (*SaveProjectLayoutUseCase, error) {
	return &SaveProjectLayoutUseCase{repo: repo}, nil
}

type SaveProjectLayoutInput struct {
	ProjectID int32 `json:"project_id,omitempty" param:"project_id"`
	LayoutID  int32 `json:"layout_id,omitempty"  param:"layout_id"`
}

type SaveProjectLayoutOutput struct {
	Data *entities.ProjectLayout `json:"data,omitempty"`
}

func (u SaveProjectLayoutUseCase) Execute(
	ctx context.Context,
	req SaveProjectLayoutInput,
) (*SaveProjectLayoutOutput, error) {
	ent := entities.ProjectLayout{
		ProjectID: req.ProjectID,
		LayoutID:  req.LayoutID,
	}
	layouts, err := u.repo.ListLayoutsByProjectID(ctx, req.ProjectID)
	if err != nil {
		return nil, err
	}
	for _, l := range layouts {
		if l.ID == req.LayoutID {
			return &SaveProjectLayoutOutput{Data: &ent}, nil
		}
	}
	entCreated, err := u.repo.SaveProjectLayout(ctx, ent)
	if err != nil {
		return nil, multierr.Append(err, shared.NewInternalError("failed to save project layout"))
	}
	return &SaveProjectLayoutOutput{
		Data: entCreated,
	}, nil
}
