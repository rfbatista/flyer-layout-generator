package clients

import (
	"algvisual/internal/clients/repository"
	"algvisual/internal/ports"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		ports.AsController(NewClientsController),
		NewClientService,
		repository.NewClientRepository,
	),
)
