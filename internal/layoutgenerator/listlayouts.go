package layoutgenerator

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"
)

func ListLayout(
	ctx context.Context,
	db *database.Queries,
	limit, skip int32,
) ([]entities.Layout, error) {
	var e []entities.Layout
	if limit == 0 || limit > 100 {
		limit = 10
	}
	layouts, err := db.ListLayouts(ctx, database.ListLayoutsParams{Limit: limit, Offset: skip})
	if err != nil {
		return e, err
	}
	for _, l := range layouts {
		lay := mapper.LayoutToDomain(l)
		comps, err := db.GetLayoutComponentsByLayoutID(ctx, int32(l.ID))
		if err != nil {
			return e, err
		}
		for _, c := range comps {
			comp := mapper.LayoutComponentToDomain(c)
			lay.Components = append(lay.Components, comp)
		}
		regions, err := db.GetLayoutRegionByLayoutID(ctx, int32(l.ID))
		if err != nil {
			return e, err
		}
		grid := entities.Grid{}
		for _, t := range regions {
			region := mapper.LayoutRegionToDomain(t)
			grid.Regions = append(grid.Regions, region)
		}
		lay.Grid = grid
		e = append(e, lay)
	}

	return e, nil
}
