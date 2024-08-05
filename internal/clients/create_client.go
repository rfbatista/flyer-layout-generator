package clients

import (
	"algvisual/internal/entities"
	"algvisual/internal/repositories"
	"context"
)

type CreateClientInput struct {
	Name      string `json:"name,omitempty"`
	CompanyID int32
}

type CreateClientOutput struct {
	Data entities.Client `json:"data,omitempty"`
}

func CreateClientUseCase(
	ctx context.Context,
	req CreateClientInput,
	repo repositories.ClientRepository,
) (*CreateClientOutput, error) {
	e := entities.Client{
		Name:      req.Name,
		CompanyID: req.CompanyID,
	}
	clientCreated, err := repo.Create(ctx, e)
	if err != nil {
		return nil, err
	}
	return &CreateClientOutput{
		Data: clientCreated,
	}, nil
}
