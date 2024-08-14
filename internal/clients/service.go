package clients

import (
	"algvisual/database"
	"algvisual/internal/clients/repository"
	"algvisual/internal/clients/usecase"

	"github.com/labstack/echo/v4"
)

func NewClientService(repo repository.ClientRepository, db *database.Queries) ClientService {
	return ClientService{repo: repo, db: db}
}

type ClientService struct {
	db   *database.Queries
	repo repository.ClientRepository
}

func (c ClientService) CreateClient(
	ctx echo.Context,
	req usecase.CreateClientInput,
) (*usecase.CreateClientOutput, error) {
	return usecase.CreateClientUseCase(
		ctx, req, c.repo,
	)
}

func (c ClientService) ListClients(
	ctx echo.Context,
	req usecase.ListClientsInput,
) (*usecase.ListClientsOutput, error) {
	return usecase.ListClientsUseCase(
		ctx, req, c.db,
	)
}
