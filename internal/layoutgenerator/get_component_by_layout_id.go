package layoutgenerator

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type GetComponentsByLayoutIDInput struct {
	LayoutID int32
}

type GetComponentsByLayoutIDOutput struct {
	Components []entities.LayoutComponent
}

func GetComponentsByLayoutIDUseCase(
	req GetComponentsByLayoutIDInput,
	db *database.Queries,
	ctx context.Context,
) (GetComponentsByLayoutIDOutput, error) {
	var out GetComponentsByLayoutIDOutput
	comps, err := db.GetLayoutComponentsByLayoutID(ctx, req.LayoutID)
	if err != nil {
		return out, err
	}
	var compentities []entities.LayoutComponent
	for _, c := range comps {
		comp := mapper.LayoutComponentToDomain(c)
		dbElements, err := db.GetDesignElementsByComponentID(ctx, pgtype.Int4{Int32: c.ID, Valid: true})
		if err != nil {
			return out, err
		}
		var elements []entities.LayoutElement
		for _, el := range dbElements {
			elements = append(elements, mapper.ToDesignElementEntitie(el))
		}
		comp.Elements = elements
		compentities = append(compentities, comp)
	}
	out.Components = compentities
	return out, nil
}
