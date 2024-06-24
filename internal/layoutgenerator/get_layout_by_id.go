package layoutgenerator

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"
)

type GetLayoutByIDRequest struct {
	LayoutID int32 `param:"request_id"`
}

type GetLayoytByIDOutput struct {
	Layout entities.Layout
}

func GetLayoutByIDUseCase(
	ctx context.Context,
	db *database.Queries,
	req GetLayoutByIDRequest,
) (GetLayoytByIDOutput, error) {
	var out GetLayoytByIDOutput
	l, err := db.GetLayoutByID(ctx, int64(req.LayoutID))
	if err != nil {
		return out, err
	}
	lay := mapper.LayoutToDomain(l)
	comps, err := db.GetLayoutComponentsByLayoutID(ctx, int32(l.ID))
	if err != nil {
		return out, err
	}
	for _, c := range comps {
		comp := mapper.LayoutComponentToDomain(c)
		lay.Components = append(lay.Components, comp)
	}
	elements, err := db.GetLayoutElementsByLayoutID(ctx, req.LayoutID)
	if err != nil {
		return out, err
	}
	for _, e := range elements {
		delement := mapper.ToDesignElementEntitie(e)
		lay.Elements = append(lay.Elements, delement)
	}
	grid := entities.Grid{}
	lay.Grid = grid
	out.Layout = lay
	return out, nil
}
