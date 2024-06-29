package templates

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
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
	Data  []entities.Template `json:"data,omitempty"`
}

func ListTemplatesByProjectIdUseCase(
	ctx context.Context,
	req ListTemplatesByProjectIdInput,
	db *database.Queries,
) (*ListTemplatesByProjectIdOutput, error) {
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
