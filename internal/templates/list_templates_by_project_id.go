package templates

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type ListTemplatesByProjectIdInput struct {
	ProjectID int32 `form:"project_id" json:"project_id,omitempty"`
}

type ListTemplatesByProjectIdOutput struct {
	Data []entities.Template
}

func ListTemplatesByProjectIdUseCase(
	ctx context.Context,
	req ListTemplatesByProjectIdInput,
	db *database.Queries,
) (*ListTemplatesByProjectIdOutput, error) {
	ts, err := db.ListTemplatesByProjectID(ctx, database.ListTemplatesByProjectIDParams{
		ProjectID: pgtype.Int4{Int32: req.ProjectID, Valid: true},
		Limit:     100,
		Offset:    0,
	})
	if err != nil {
		return nil, err
	}
	var templates []entities.Template
	for _, t := range ts {
		templates = append(templates, mapper.TemplateToDomain(t))
	}
	return &ListTemplatesByProjectIdOutput{
		Data: templates,
	}, nil
}
