package projects

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories"
	"context"
)

type ListProjectsUseCase struct {
	db   *database.Queries
	repo *repositories.ProjectRepository
}

func NewListProjectsUseCase(
	db *database.Queries,
	repo *repositories.ProjectRepository,
) ListProjectsUseCase {
	return ListProjectsUseCase{
		db:   db,
		repo: repo,
	}
}

type ListProjectsInput struct {
	Page        int32 `query:"page"  json:"page,omitempty"`
	Limit       int32 `query:"limit" json:"limit,omitempty"`
	Order       int32 `query:"order"`
	UserSession entities.UserSession
}

type ListProjectsOutput struct {
	Page  int32              `query:"page"  json:"page"`
	Limit int32              `query:"limit" json:"limit"`
	Data  []entities.Project `              json:"data"`
}

func (l ListProjectsUseCase) Execute(
	ctx context.Context,
	req ListProjectsInput,
) (*ListProjectsOutput, error) {
	pr, err := l.repo.ListProjects(ctx, int32(req.UserSession.CompanyID))
	if err != nil {
		return nil, err
	}
	return &ListProjectsOutput{
		Page:  req.Page,
		Limit: req.Limit,
		Data:  pr,
	}, nil
}
