package templates

import (
	"algvisual/database"
	"algvisual/internal/templates/repository"
	"algvisual/internal/templates/usecase"
	"context"
)

func NewTemplateService(
	db *database.Queries,
	repo *repository.TemplateRepository,
) TemplatesService {
	return TemplatesService{db: db, repo: repo}
}

type TemplatesService struct {
	db   *database.Queries
	repo *repository.TemplateRepository
}

func (t TemplatesService) GetTemplateByID(
	ctx context.Context,
	in usecase.GetTemplateByIdInput,
) (*usecase.GetTemplateByIdOutput, error) {
	return usecase.GetTemplateByIdUseCase(ctx, in, t.db)
}

func (t TemplatesService) ListAdaptationTemplates(
	ctx context.Context,
	in usecase.ListAdaptationTemplatesInput,
) (*usecase.ListAdaptationTemplatesOutput, error) {
	return usecase.ListAdaptationTemplatesUseCase(ctx, in, t.repo)
}
