package grammars

import "algvisual/internal/domain/entities"

func StageFindCells(
	original entities.Layout,
	template entities.Template,
	grid entities.Grid,
) (*entities.Layout, *entities.Grid, error) {
	var out entities.Layout
	var stage1components []entities.LayoutComponent
	for _, c := range original.Components {
		cell := grid.WhereIsPoint(c.Center())
		if cell == nil {
			continue
		}
		if c.IsBackground() {
			continue
		}
		c.Pivot = cell.Position()
		grid.OcupyByPosition(cell.Position(), c.ID)
		c.MoveTo(cell.UpLeft())
		gridcont, found, err := grid.FindPositionToFitGridContainerDontCheckColision(
			cell.Position(),
			grid.ContainerToGridContainer(c.InnerContainer),
			c.ID,
		)
		if err != nil || !found {
			continue
		}
		positions := grid.ContainerToPositions(
			gridcont.ToContainer(grid.CellWidth(), grid.CellHeight()),
		)
		grid.OcupyByPositionList(positions, c.ID)
		c.Positions = positions
		cont := grid.PositionsToContainer(positions)
		c.ScaleToFitInSize(cont.Width(), cont.Height())
		c.MoveTo(cont.UpperLeft)
		stage1components = append(stage1components, c)
	}
	out.Components = stage1components
	out.Template = template
	out.DesignID = original.DesignID
	out.Width = original.Width
	out.Height = original.Height
	out.Grid = grid
	return &out, &grid, nil
}
