package editor

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"context"

	"go.uber.org/zap"
)

type request struct{}

func Props(
	ctx context.Context,
	db *database.Queries,
	log *zap.Logger,
	req request,
) (PageProps, error) {
	var props PageProps
	props.types = []string{
		entities.ComponentTypeProduto.ToString(),
		entities.ComponentTypeCallToAction.ToString(),
		entities.ComponentTypeMarca.ToString(),
		entities.ComponentTypeCelebridade.ToString(),
		entities.ComponentTypeGrafismo.ToString(),
		entities.ComponentTypeOferta.ToString(),
		entities.ComponentTypePackshot.ToString(),
		entities.ComponentTypeModelo.ToString(),
		entities.ComponentTypePlanoDeFundo.ToString(),
	}
	return props, nil
}
