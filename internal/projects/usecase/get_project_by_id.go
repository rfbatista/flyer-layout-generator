package usecase

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"

	"github.com/labstack/echo/v4"
)

type GetProjectByIdInput struct {
	ProjectID int32 `param:"project_id" json:"project_id,omitempty"`
}

type GetProjectByIdOutput struct {
	Data entities.Project `json:"data,omitempty"`
}

func GetProjectByIdUseCase(
	c echo.Context,
	req GetProjectByIdInput,
	db *database.Queries,
) (*GetProjectByIdOutput, error) {
	ctx := c.Request().Context()
	pr, err := db.GetProjectByID(ctx, int64(req.ProjectID))
	if err != nil {
		return nil, err
	}
	project := mapper.ProjectToDomain(pr)
	cl, err := db.GetClientByID(ctx, int64(pr.ClientID.Int32))
	if err != nil {
		return nil, err
	}
	client := mapper.ClientToDomain(cl)
	project.Client = client
	ad, err := db.GetAdvertiserByID(ctx, int64(pr.AdvertiserID.Int32))
	if err != nil {
		return nil, err
	}
	advertiser := mapper.AdvertiserToDomain(ad)
	project.Advertiser = advertiser
	return &GetProjectByIdOutput{
		Data: project,
	}, nil
}
