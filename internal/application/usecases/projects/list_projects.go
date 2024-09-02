package projects

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories/mapper"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type ListProjectsUseCase struct {
	db *database.Queries
}

func NewListProjectsUseCase(
	db *database.Queries,
) ListProjectsUseCase {
	return ListProjectsUseCase{
		db: db,
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
	c echo.Context,
	req ListProjectsInput,
) (*ListProjectsOutput, error) {
	session := req.UserSession
	ctx := c.Request().Context()
	pr, err := l.db.ListProjects(ctx, database.ListProjectsParams{
		Limit:     req.Limit,
		Offset:    req.Page,
		CompanyID: pgtype.Int4{Int32: int32(session.CompanyID), Valid: true},
	})
	if err != nil {
		return nil, err
	}
	var projects []entities.Project
	for _, p := range pr {
		ad, err := l.db.GetAdvertiserByID(ctx, int64(p.AdvertiserID.Int32))
		if err != nil {
			return nil, err
		}
		cl, err := l.db.GetClientByID(ctx, int64(p.ClientID.Int32))
		if err != nil {
			return nil, err
		}
		project := mapper.ProjectToDomain(p)
		project.Client = mapper.ClientToDomain(cl)
		project.Advertiser = mapper.AdvertiserToDomain(ad)
		projects = append(projects, project)
	}
	return &ListProjectsOutput{
		Page:  req.Page,
		Limit: req.Limit,
		Data:  projects,
	}, nil
}
