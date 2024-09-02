package adaptations

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/repositories"
	"context"
)

type ListAdaptationTemplatesInput struct{}

type ListAdaptationTemplatesOutput struct {
	Data []entities.Template
}

func ListAdaptationTemplatesUseCase(
	ctx context.Context,
	req ListAdaptationTemplatesInput,
	repo *repositories.TemplateRepository,
) (*ListAdaptationTemplatesOutput, error) {
	templates, err := repo.List(ctx, repositories.ListTemplatesParams{
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
