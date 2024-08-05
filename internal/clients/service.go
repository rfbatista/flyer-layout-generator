package clients

import (
	"algvisual/internal/repositories"
	"context"
)

func NewClientService(repo repositories.ClientRepository) ClientService {
	return ClientService{repo: repo}
}

type ClientService struct {
	repo repositories.ClientRepository
}

func (c ClientService) CreateClient(
	ctx context.Context,
	req CreateClientInput,
) (*CreateClientOutput, error) {
	return CreateClientUseCase(
		ctx, req, c.repo,
	)
}
