package generate

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/layoutgenerator"
	"algvisual/internal/templates"
	"context"
	"encoding/json"

	"go.uber.org/zap"
)

type request struct {
	DesignID int32 `param:"design_id"`
	LayoutID int32 `param:"layout_id"`
}

func Props(
	ctx context.Context,
	db *database.Queries,
	log *zap.Logger,
	req request,
) (PageProps, error) {
	var props PageProps
	props.designID = req.DesignID
	layout, err := layoutgenerator.GetLayoutByIDUseCase(
		ctx,
		db,
		layoutgenerator.GetLayoutByIDRequest{
			LayoutID: req.LayoutID,
		},
	)
	if err != nil {
		return props, err
	}
	props.layout = layout.Layout
	out, err := json.Marshal(layout.Layout)
	if err != nil {
		panic(err)
	}
	props.layoutjson = string(out)
	templateOut, err := templates.ListTemplatesUseCase(
		ctx,
		templates.ListTemplatesUseCaseRequest{},
		db,
		log,
	)
	if err != nil {
		return props, err
	}
	props.template = templateOut.Data
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
