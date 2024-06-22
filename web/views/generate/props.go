package generate

import (
	"algvisual/internal/database"
	"algvisual/internal/designprocessor"
	"algvisual/internal/entities"
	"algvisual/internal/templates"
	"context"

	"go.uber.org/zap"
)

func Props(ctx context.Context, db *database.Queries, log *zap.Logger) (PageProps, error) {
	var props PageProps
	out, err := designprocessor.ListDesignFiles(ctx, designprocessor.ListDesignFilesRequest{}, db, log)
	if err != nil {
		return props, err
	}
	templateOut, err := templates.ListTemplatesUseCase(ctx, templates.ListTemplatesUseCaseRequest{}, db, log)
	if err != nil {
		return props, err
	}
	props.files = out.Data
	props.template = templateOut.Data
	props.types = []string{
		entities.ComponentTypeGrafismo.ToString(),
		entities.ComponentTypeMarca.ToString(),
		entities.ComponentTypeModelo.ToString(),
		entities.ComponentTypeOferta.ToString(),
		entities.ComponentTypeCelebridade.ToString(),
	}
	return props, nil
}
