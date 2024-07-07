package templates

import (
	"algvisual/database"
	"context"
)

func NewTemplateService(db *database.Queries) TemplatesService {
	return TemplatesService{db: db}
}

type TemplatesService struct {
	db *database.Queries
}

func (t TemplatesService) GetTemplateByID(
	ctx context.Context,
	in GetTemplateByIdInput,
) (*GetTemplateByIdOutput, error) {
	return GetTemplateByIdUseCase(ctx, in, t.db)
}
