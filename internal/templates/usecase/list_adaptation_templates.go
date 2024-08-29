package usecase

import (
	"algvisual/internal/entities"
	"algvisual/internal/templates/repository"
	"context"
)

type ListAdaptationTemplatesInput struct{}

type ListAdaptationTemplatesOutput struct {
	Data []entities.Template
}

func ListAdaptationTemplatesUseCase(
	ctx context.Context,
	req ListAdaptationTemplatesInput,
	repo *repository.TemplateRepository,
) (*ListAdaptationTemplatesOutput, error) {
	templates, err := repo.List(ctx, repository.ListTemplatesParams{
		Limit:        20,
		Offset:       0,
		FilterByType: true,
		Type:         entities.TemplateTypeAdaptation,
	})
	if err != nil {
		return nil, err
	}
	return &ListAdaptationTemplatesOutput{
		Data: templates,
	}, nil
}
