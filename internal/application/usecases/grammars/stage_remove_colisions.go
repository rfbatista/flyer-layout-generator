package grammars

import (
	"algvisual/internal/domain/entities"
	"fmt"
)

func StageRemoveColisions(
	original entities.Layout,
	prevLayout entities.Layout,
	template entities.Template,
	grid entities.Grid,
) (*entities.Layout, *entities.Grid, error) {
	var out entities.Layout
	var stagecomponents []entities.LayoutComponent
	stagegrid, _ := entities.NewGrid(
		entities.WithDefault(template.Width, template.Height),
		entities.WithCells(grid.SlotsX, grid.SlotsY),
	)
	for _, c := range prevLayout.Components {
		if len(c.Positions) == 0 {
			continue
		}
		ctype := c.Type
		if ctype == "o" {
			fmt.Println(ctype)
		}
		positions := stagegrid.ContainerToPositions(c.InnerContainer)
		if stagegrid.IsPositionListOcupiedByOtherThanThisId(positions, c.ID) {
			stagegrid.RemoveFromAllCells(c.ID)
			continue
		}
		stagegrid.OcupyByPositionList(positions, c.ID)
		stagecomponents = append(stagecomponents, c)
	}
	out.Components = stagecomponents
	out.Template = template
	out.DesignID = original.DesignID
	out.Width = template.Width
	out.Height = template.Height
	out.Grid = *stagegrid
	return &out, stagegrid, nil
}
