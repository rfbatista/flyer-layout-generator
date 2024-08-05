package clients

import (
	"context"
)

func NewClientService(repo ClientRepository) ClientService {
	return ClientService{repo: repo}
}

type ClientService struct {
	repo ClientRepository
}

func (c ClientService) CreateClient(
	ctx context.Context,
	req CreateClientInput,
) (*CreateClientOutput, error) {
	return CreateClientUseCase(
		ctx, req, c.repo,
	)
}
