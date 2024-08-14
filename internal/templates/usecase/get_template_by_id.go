package usecase

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"

	"github.com/labstack/echo/v4"
)

type GetTemplateByIdInput struct {
	TemplateID int32 `json:"template_id,omitempty"`
}

type GetTemplateByIdOutput struct {
	Data entities.Template
}

func GetTemplateByIdUseCase(
	c echo.Context,
	req GetTemplateByIdInput,
	queries *database.Queries,
) (*GetTemplateByIdOutput, error) {
	ctx := c.Request().Context()
	raw, err := queries.GetTemplate(ctx, req.TemplateID)
	if err != nil {
		return nil, err
	}
	return &GetTemplateByIdOutput{
		Data: mapper.TemplateToDomain(raw.Template),
	}, nil
}
