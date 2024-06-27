package grammars

import "algvisual/internal/entities"

func StageRemoveColisions(
	original entities.Layout,
	prevLayout entities.Layout,
	template entities.Template,
	grid entities.Grid,
) (*entities.Layout, *entities.Grid, error) {
	var out entities.Layout
	var stagecomponents []entities.LayoutComponent
	for _, c := range prevLayout.Components {
		if len(c.Positions) == 0 {
			continue
		}
		if grid.IsPositionListOcupiedByOtherThanThisId(c.Positions, c.ID) {
			grid.RemoveFromAllCells(c.ID)
			continue
		}
		stagecomponents = append(stagecomponents, c)
	}
	out.Components = stagecomponents
	out.Template = template
	out.DesignID = original.DesignID
	out.Width = template.Width
	out.Height = template.Height
	out.Grid = grid
	return &out, &grid, nil
}
