package repository

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func NewTemplateRepository() (*TemplateRepository, error) {
	return &TemplateRepository{}, nil
}

type TemplateRepository struct {
	db *database.Queries
}

type ListTemplatesParams struct {
	Limit           int32 `json:"limit"`
	Offset          int32 `json:"offset"`
	CompanyID       int32
	Type            entities.TemplateType
	FilterByCompany bool
	FilterByType    bool
	FilterByProject bool
}

func (t TemplateRepository) List(
	ctx context.Context,
	p ListTemplatesParams,
) ([]entities.Template, error) {
	var list []entities.Template
	raw, err := t.db.ListTemplates(ctx, database.ListTemplatesParams{
		Limit:     p.Limit,
		Offset:    p.Offset,
		CompanyID: pgtype.Int4{Int32: p.CompanyID, Valid: p.CompanyID != 0},
		Type: database.NullTemplateType{
			TemplateType: mapper.TemplateTypeToDatabase(p.Type),
			Valid:        p.Type != "",
		},
		FilterByCompany: p.FilterByCompany,
		FilterByType:    p.FilterByType,
		FilterByProject: p.FilterByProject,
	})
	if err != nil {
		return list, err
	}
	for _, r := range raw {
		list = append(list, mapper.TemplateToDomain(r))
	}
	return list, nil
}
