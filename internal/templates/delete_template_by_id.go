package templates

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"
)

type DeleteTemplateByIdInput struct {
	TemplateID int32 `json:"template_id,omitempty" form:"template_id" param:"template_id"`
}

type DeleteTemplateByIdOutput struct {
	Data entities.Template
}

func DeleteTemplateByIdUseCase(
	ctx context.Context,
	req DeleteTemplateByIdInput,
	db *database.Queries,
) (*DeleteTemplateByIdOutput, error) {
	t, err := db.GetTemplateByID(ctx, req.TemplateID)
	if err != nil {
		return nil, err
	}
	err = db.DeleteTemplateByID(ctx, req.TemplateID)
	if err != nil {
		return nil, err
	}
	return &DeleteTemplateByIdOutput{
		Data: mapper.TemplateToDomain(t),
	}, nil
}
