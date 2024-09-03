package repositories

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func NewTemplateRepository(db *database.Queries) (*TemplateRepository, error) {
	return &TemplateRepository{db: db}, nil
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

func (t TemplateRepository) GetByID(
	ctx context.Context,
	id int32,
) (*entities.Template, error) {
	raw, err := t.db.GetTemplateByID(ctx, id)
	if err != nil {
		return nil, err
	}
	temp := mapper.TemplateToDomain(raw)
	return &temp, nil
}
