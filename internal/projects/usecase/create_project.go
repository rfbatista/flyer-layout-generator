package usecase

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/infra/middlewares"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type CreateProjectInput struct {
	Name         string `form:"name"          json:"name,omitempty"`
	ClientID     int32  `form:"client_id"     json:"client_id,omitempty"`
	AdvertiserID int32  `form:"advertiser_id" json:"advertiser_id,omitempty"`
}

type CreateProjectOutput struct {
	Project entities.Project `json:"project,omitempty"`
}

func CreateProjectUseCase(
	c echo.Context,
	req CreateProjectInput,
	db *database.Queries,
) (*CreateProjectOutput, error) {
	ctx := c.Request().Context()
	session := c.(*middlewares.ApplicationContext)
	pr, _ := db.CreateProject(ctx, database.CreateProjectParams{
		Name:         req.Name,
		ClientID:     pgtype.Int4{Int32: req.ClientID, Valid: true},
		CompanyID:    pgtype.Int4{Int32: int32(session.UserSession().CompanyID), Valid: true},
		AdvertiserID: pgtype.Int4{Int32: req.AdvertiserID, Valid: true},
	})
	project, err := GetProjectByIdUseCase(c, GetProjectByIdInput{ProjectID: int32(pr.ID)}, db)
	if err != nil {
		return nil, err
	}
	return &CreateProjectOutput{
		Project: project.Data,
	}, nil
}
