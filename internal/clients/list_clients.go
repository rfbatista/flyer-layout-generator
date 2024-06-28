package clients

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"
)

type ListClientsInput struct {
	Page  int32 `query:"page"  json:"page,omitempty"`
	Limit int32 `query:"limit" json:"skip,omitempty"`
}

type ListClientsOutput struct {
	Clients []entities.Client `json:"clients,omitempty"`
}

func ListClientsUseCase(
	ctx context.Context,
	req ListClientsInput,
	db *database.Queries,
) (*ListClientsOutput, error) {
	cls, err := db.ListClients(ctx, database.ListClientsParams{
		Limit:  req.Limit,
		Offset: req.Page,
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
