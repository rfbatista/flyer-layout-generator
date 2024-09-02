package advertisers

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/middlewares"
	"algvisual/internal/infrastructure/repositories"

	"github.com/labstack/echo/v4"
)

type CreateAdvertiserInput struct {
	Name string `json:"name,omitempty"`
}

type CreateAdvertiserOutput struct {
	Data entities.Advertiser `json:"data,omitempty"`
}

func CreateAdvertiserUseCase(
	ctx echo.Context,
	req CreateAdvertiserInput,
	repo repositories.AdvertiserRepository,
) (*CreateAdvertiserOutput, error) {
	cc := ctx.(*middlewares.ApplicationContext)
	e := entities.Advertiser{
		Name:      req.Name,
		CompanyID: int32(cc.UserSession().CompanyID),
	}
	eCreated, err := repo.Create(ctx, e)
	if err != nil {
		return nil, err
	}
	return &CreateAdvertiserOutput{
		Data: eCreated,
	}, nil
}
