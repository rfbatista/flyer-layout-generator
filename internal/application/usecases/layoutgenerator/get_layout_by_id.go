package layoutgenerator

import (
	"algvisual/internal/application/usecases/designassets"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories/mapper"
	"algvisual/internal/shared"
	"context"
)

type GetLayoutByIDRequest struct {
	LayoutID int32 `param:"layout_id"`
}

type GetLayoutByIDOutput struct {
	Layout entities.Layout
}

func GetLayoutByIDUseCase(
	ctx context.Context,
	db *database.Queries,
	req GetLayoutByIDRequest,
	das *designassets.DesignAssetService,
) (GetLayoutByIDOutput, error) {
	var out GetLayoutByIDOutput
	l, err := db.GetLayoutByID(ctx, int64(req.LayoutID))
	if err != nil {
		return out, err
	}
	lay := mapper.LayoutToDomain(l)
	elements, err := db.GetLayoutElementsByLayoutID(ctx, req.LayoutID)
	if err != nil {
		return out, err
	}
	for _, e := range elements {
		assets, err := das.GetDesignAssetByID(
			ctx,
			designassets.GetDesignAssetByIdInput{ID: e.AssetID},
		)
		if err != nil {
			return out, err
		}
		delement := mapper.ToDesignElementEntitie(e)
		delement.Properties = append(delement.Properties, assets.Data.Properties...)
		lay.Elements = append(lay.Elements, delement)
	}
	compHash := make(map[int32][]entities.LayoutElement)
	for _, c := range lay.Elements {
		if c.ComponentID != 0 {
			compHash[c.ComponentID] = append(compHash[c.ComponentID], c)
		}
	}
	var components []entities.LayoutComponent
	var bg *entities.LayoutComponent
	for k := range compHash {
		data, compErr := db.GetComponentByID(ctx, k)
		if compErr != nil {
			compErr = shared.WrapWithAppError(
				compErr,
				"Não foi possivel encontrar os componentes do arquivo Photoshop",
				"",
			)
			return out, compErr
		}
		comp := mapper.TodesignComponentEntitie(data)
		comp.Elements = compHash[k]
		if comp.IsBackground() {
			bg = &comp
		} else {
			components = append(components, comp)
		}
	}
	for idx := range components {
		for elidx := range components[idx].Elements {
			components[idx].Elements[elidx].Type = components[idx].Type
		}
	}
	if bg != nil {
		for elidx := range bg.Elements {
			bg.Elements[elidx].Type = bg.Type
		}
	}
	lay.Components = components
	lay.Background = bg
	grid := entities.Grid{}
	lay.Grid = grid
	out.Layout = lay
	return out, nil
}
