package usecase

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/infra/middlewares"
	"algvisual/internal/mapper"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type ListClientsInput struct {
	Page  int32 `query:"page"  json:"page,omitempty"`
	Limit int32 `query:"limit" json:"skip,omitempty"`
}

type ListClientsOutput struct {
	Clients []entities.Client `json:"clients,omitempty"`
}

func ListClientsUseCase(
	ctx echo.Context,
	req ListClientsInput,
	db *database.Queries,
) (*ListClientsOutput, error) {
	cc := ctx.(*middlewares.ApplicationContext)
	cls, err := db.ListClients(ctx.Request().Context(), database.ListClientsParams{
		Limit:     req.Limit,
		Offset:    req.Page,
		CompanyID: pgtype.Int4{Int32: int32(cc.UserSession().CompanyID), Valid: true},
	})
	if err != nil {
		return nil, err
	}
	var clients []entities.Client
	for _, c := range cls {
		clients = append(clients, mapper.ClientToDomain(c))
	}
	return &ListClientsOutput{
		Clients: clients,
	}, nil
}
