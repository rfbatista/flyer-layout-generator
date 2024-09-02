package templates

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/middlewares"
	"algvisual/internal/infrastructure/repositories/mapper"
	"algvisual/internal/shared"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ListTemplatesUseCaseRequest struct {
	Limit int `query:"limit" json:"limit,omitempty"`
	Skip  int `query:"skip"  json:"skip,omitempty"`
}

type ListTemplatesUseCaseResult struct {
	Data []entities.Template `json:"data,omitempty"`
}

func ListTemplatesUseCase(
	c echo.Context,
	req ListTemplatesUseCaseRequest,
	queries *database.Queries,
	log *zap.Logger,
) (*ListTemplatesUseCaseResult, error) {
	ctx := c.Request().Context()
	session := c.(*middlewares.ApplicationContext)
	limit := req.Limit
	if limit == 0 {
		limit = 10
	}
	result, err := queries.ListTemplates(ctx, database.ListTemplatesParams{
		Limit:           int32(limit),
		Offset:          int32(req.Skip),
		FilterByCompany: true,
		CompanyID:       pgtype.Int4{Int32: int32(session.UserSession().CompanyID), Valid: true},
	})
	if err != nil {
		err = shared.WrapWithAppError(err, "failed to list templates", "")
		log.Error(err.Error())
		return nil, err
	}
	var templates []entities.Template
	for _, t := range result {
		templates = append(templates, mapper.TemplateToDomain(t))
	}
	return &ListTemplatesUseCaseResult{
		Data: templates,
	}, nil
}
