package templates

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories/mapper"
	"context"
)

type GetTemplateByIdInput struct {
	TemplateID int32 `json:"template_id,omitempty" param:"template_id"`
}

type GetTemplateByIdOutput struct {
	Data entities.Template `json:"data,omitempty"`
}

func GetTemplateByIdUseCase(
	ctx context.Context,
	req GetTemplateByIdInput,
	queries *database.Queries,
) (*GetTemplateByIdOutput, error) {
	raw, err := queries.GetTemplate(ctx, req.TemplateID)
	if err != nil {
		return nil, err
	}
	return &GetTemplateByIdOutput{
		Data: mapper.TemplateToDomain(raw.Template),
	}, nil
}
