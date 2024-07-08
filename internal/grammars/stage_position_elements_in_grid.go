package grammars

import "algvisual/internal/entities"

func PositionElementsInGrid(
	original entities.Layout,
	prevLayout entities.Layout,
	template entities.Template,
	prevGrid entities.Grid,
) (*entities.Layout, *entities.Grid, error) {
	var out entities.Layout
	var stagecomponents []entities.LayoutComponent
	stagegrid, _ := entities.NewGrid(
		entities.WithDefault(template.Width, template.Height),
		entities.WithCells(prevGrid.SlotsX, prevGrid.SlotsY),
	)
	for _, c := range prevLayout.Components {
		if len(c.Positions) == 0 {
			continue
		}
		cont := stagegrid.PositionsToContainer(c.Positions)
		c.ScaleToFitInSize(cont.Width(), cont.Height())
		gridcont, found, err := stagegrid.FindPositionToFitGridContainerDontCheckColision(
			c.Pivot,
			stagegrid.ContainerToGridContainer(c.InnerContainer),
			c.ID,
		)
		if err != nil || !found {
			continue
		}
		positions := stagegrid.ContainerToPositions(
			gridcont.ToContainer(stagegrid.CellWidth(), stagegrid.CellHeight()),
		)
		stagegrid.OcupyByPositionList(c.Positions, c.ID)
		c.Positions = positions
		ncont := stagegrid.PositionsToContainer(positions)
		c.MoveTo(ncont.UpperLeft)
		c.CenterInContainer(ncont)
		c.GridContainer = ncont
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
