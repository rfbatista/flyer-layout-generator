package templates

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/shared"
	"context"

	"go.uber.org/zap"
)

type ListTemplatesUseCase struct {
	repo    *repositories.TemplateRepository
	queries *database.Queries
	log     *zap.Logger
}

func NewListTemplatesUseCase(
	repo *repositories.TemplateRepository,
	queries *database.Queries,
	log *zap.Logger,
) (*ListTemplatesUseCase, error) {
	return &ListTemplatesUseCase{
		repo:    repo,
		queries: queries,
		log:     log,
	}, nil
}

type ListTemplatesUseCaseRequest struct {
	Limit     int                  `query:"limit" json:"limit,omitempty" param:"limit"`
	Skip      int                  `query:"skip"  json:"skip,omitempty"  param:"skip"`
	ProjectID int32                `                                     param:"project_id"`
	Session   entities.UserSession `                                     param:"session"`
}

type ListTemplatesUseCaseResult struct {
	Data []entities.Template `json:"data,omitempty"`
}

func (l ListTemplatesUseCase) Execute(
	ctx context.Context,
	req ListTemplatesUseCaseRequest,
) (*ListTemplatesUseCaseResult, error) {
	limit := req.Limit
	if limit == 0 {
		limit = 10
	}
	templates, err := l.repo.List(ctx, repositories.ListTemplatesParams{
		Limit:              int32(limit),
		Offset:             int32(req.Skip),
		FilterByCompany:    true,
		CompanyID:          req.Session.UserID,
		FilterByProject:    req.ProjectID != 0,
		ProjectID:          req.ProjectID,
		AddPublicTemplates: true,
	})
	if err != nil {
		l.log.Error("failed to find templates", zap.Error(err))
		return nil, shared.NewInternalError("failed to list templates")
	}
	return &ListTemplatesUseCaseResult{
		Data: templates,
	}, nil
}
