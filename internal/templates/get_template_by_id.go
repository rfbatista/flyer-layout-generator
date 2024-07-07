package templates

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"
)

type GetTemplateByIdInput struct {
	TemplateID int32 `json:"template_id,omitempty"`
}

type GetTemplateByIdOutput struct {
	Data entities.Template
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
