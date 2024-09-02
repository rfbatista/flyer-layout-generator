package clients

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/middlewares"
	"algvisual/internal/infrastructure/repositories"

	"github.com/labstack/echo/v4"
)

type CreateClientInput struct {
	Name      string `json:"name,omitempty"`
	CompanyID int32
}

type CreateClientOutput struct {
	Data entities.Client `json:"data,omitempty"`
}

func CreateClientUseCase(
	ctx echo.Context,
	req CreateClientInput,
	repo repositories.ClientRepository,
) (*CreateClientOutput, error) {
	cc := ctx.(*middlewares.ApplicationContext)
	e := entities.Client{
		Name:      req.Name,
		CompanyID: int32(cc.UserSession().CompanyID),
	}
	clientCreated, err := repo.Create(ctx, e)
	if err != nil {
		return nil, err
	}
	return &CreateClientOutput{
		Data: clientCreated,
	}, nil
}
