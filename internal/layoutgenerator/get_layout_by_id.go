package layoutgenerator

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"
	"fmt"
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
	fmt.Println("oooooooooooooooooooooooooo")
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
	lay.Components = comps.Components
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
	fmt.Println(fmt.Sprintf("found %d", len(out.Layout.Elements)))
	return out, nil
}
