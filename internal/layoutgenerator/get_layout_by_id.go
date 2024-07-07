package layoutgenerator

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type GetLayoutByIDRequest struct {
	LayoutID int32 `param:"layout_id"`
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
	comps, err := GetComponentsByLayoutIDUseCase(
		GetComponentsByLayoutIDInput{LayoutID: req.LayoutID},
		db,
		ctx,
	)
	if err != nil {
		return out, err
	}
	elements, err := db.GetLayoutElementsByLayoutID(ctx, req.LayoutID)
	if err != nil {
		return out, err
	}
	for _, e := range elements {
		delement := mapper.ToDesignElementEntitie(e)
		lay.Elements = append(lay.Elements, delement)
	}
	var components []entities.LayoutComponent
	for _, comp := range comps.Components {
		els, err := db.GetDesignElementsByComponentID(ctx, pgtype.Int4{Int32: comp.ID, Valid: true})
		if err != nil {
			return out, err
		}
		var elements []entities.LayoutElement
		for _, el := range els {
			elements = append(elements, mapper.ToDesignElementEntitie(el))
		}
		comp.Elements = elements
		components = append(components, comp)
	}
	lay.Components = components
	grid := entities.Grid{}
	lay.Grid = grid
	out.Layout = lay
	return out, nil
}
