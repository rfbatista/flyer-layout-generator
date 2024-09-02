package templates

import (
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories"
	"context"
)

func NewTemplateService(
	db *database.Queries,
	repo *repositories.TemplateRepository,
) TemplatesService {
	return TemplatesService{db: db, repo: repo}
}

type TemplatesService struct {
	db   *database.Queries
	repo *repositories.TemplateRepository
}

func (t TemplatesService) GetTemplateByID(
	ctx context.Context,
	in GetTemplateByIdInput,
) (*GetTemplateByIdOutput, error) {
	return GetTemplateByIdUseCase(ctx, in, t.db)
}

func (t TemplatesService) ListAdaptationTemplates(
	ctx context.Context,
	in ListAdaptationTemplatesInput,
) (*ListAdaptationTemplatesOutput, error) {
	return ListAdaptationTemplatesUseCase(ctx, in, t.repo)
}
