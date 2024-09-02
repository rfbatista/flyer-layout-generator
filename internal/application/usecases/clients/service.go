package clients

import (
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories"

	"github.com/labstack/echo/v4"
)

func NewClientService(repo repositories.ClientRepository, db *database.Queries) ClientService {
	return ClientService{repo: repo, db: db}
}

type ClientService struct {
	db   *database.Queries
	repo repositories.ClientRepository
}

func (c ClientService) CreateClient(
	ctx echo.Context,
	req CreateClientInput,
) (*CreateClientOutput, error) {
	return CreateClientUseCase(
		ctx, req, c.repo,
	)
}

func (c ClientService) ListClients(
	ctx echo.Context,
	req ListClientsInput,
) (*ListClientsOutput, error) {
	return ListClientsUseCase(
		ctx, req, c.db,
	)
}
