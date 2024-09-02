package templates

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories/mapper"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type ListTemplatesByProjectIdInput struct {
	ProjectID int32 `param:"project_id" json:"project_id,omitempty"`
	Page      int32 `                   json:"page,omitempty"       query:"page"`
	Limit     int32 `                   json:"limit,omitempty"      query:"limit"`
}

type ListTemplatesByProjectIdOutput struct {
	Page  int32               `json:"page"`
	Limit int32               `json:"limit,omitempty"`
	Total int32               `json:"total,omitempty"`
	Data  []entities.Template `json:"data"`
}

func ListTemplatesByProjectIdUseCase(
	c echo.Context,
	req ListTemplatesByProjectIdInput,
	db *database.Queries,
) (*ListTemplatesByProjectIdOutput, error) {
	ctx := c.Request().Context()
	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.Limit > 50 {
		req.Limit = 50
	}
	ts, err := db.ListTemplatesByProjectID(ctx, database.ListTemplatesByProjectIDParams{
		ProjectID: pgtype.Int4{Int32: req.ProjectID, Valid: true},
		Limit:     req.Limit,
		Offset:    req.Page,
	})
	if err != nil {
		return nil, err
	}
	var templates []entities.Template
	for _, t := range ts {
		templates = append(templates, mapper.TemplateToDomain(t))
	}
	total, err := db.TotalTemplatesByProjectID(ctx, pgtype.Int4{Int32: req.ProjectID, Valid: true})
	if err != nil {
		return nil, err
	}
	return &ListTemplatesByProjectIdOutput{
		Page:  req.Page,
		Limit: req.Limit,
		Total: int32(total),
		Data:  templates,
	}, nil
}
