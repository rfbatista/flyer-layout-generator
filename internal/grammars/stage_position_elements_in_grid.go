package grammars

import "algvisual/internal/entities"

func PositionElementsInGrid(
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
		cont := stagegrid.PositionsToContainer(c.Positions)
		stagegrid.OcupyByPositionList(c.Positions, c.ID)
		c.ScaleToFitInSize(cont.Width(), cont.Height())
		c.MoveTo(cont.UpperLeft)
		c.CenterInContainer(cont)
		c.GridContainer = cont
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
