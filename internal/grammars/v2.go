package grammars

import (
	"algvisual/internal/entities"
	"sort"

	"go.uber.org/zap"
)

func RunV2(
	original entities.Layout,
	template entities.Template,
	gridX, gridY int32,
	log *zap.Logger,
) (*entities.Layout, error) {
	grid, err := entities.NewGrid(
		entities.WithDefault(original.Width, original.Height),
		entities.WithCells(gridX, gridY),
	)
	if err != nil {
		return nil, err
	}
	sort.Slice(original.Components, func(i, j int) bool {
		return original.Components[i].OrderPriority() > original.Components[j].OrderPriority()
	})
	// Find cells for each component in original design
	layout1, _, err := Stage1(original, template, *grid)
	if err != nil {
		return nil, err
	}
	// // Position elements in target template grid
	// layout2, _, err := Stage2(original, *layout1, template, *stage1Grid)
	// if err != nil {
	// 	return nil, err
	// }

	// // Move elements that have colision
	// sort.Slice(layout2.Components, func(i, j int) bool {
	// 	return layout2.Components[i].OrderPriority() > layout2.Components[j].OrderPriority()
	// })
	// layout3, _, err := Stage3(original, layout2, template, stage2Grid)
	// if err != nil {
	// 	return nil, err
	// }

	// // expand elements
	// layout4, _, err := Stage4(original, layout3, template, stage3Grid)
	// if err != nil {
	// 	return nil, err
	// }
	return layout1, nil
}

func Stage1(
	original entities.Layout,
	template entities.Template,
	grid entities.Grid,
) (*entities.Layout, *entities.Grid, error) {
	var out entities.Layout
	var stage1components []entities.DesignComponent
	for _, c := range original.Components {
		cell := grid.WhereIsPoint(c.Center())
		if cell == nil {
			continue
		}
		c.Pivot = cell.Position()
		cell.Ocupy(c.ID)
		c.MoveTo(cell.UpLeft())
		gridcont, found, err := grid.FindPositionToFitGridContainerDontCheckColision(cell.Position(), c.InnerContainer)
		if err != nil || !found {
			continue
		}
		positions := grid.ContainerToPositions(gridcont.ToContainer(grid.CellWidth(), grid.CellHeight()))
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

func Stage2(
	original entities.Layout,
	prevLayout entities.Layout,
	template entities.Template,
	grid entities.Grid,
) (*entities.Layout, *entities.Grid, error) {
	var out entities.Layout
	var stagecomponents []entities.DesignComponent
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

func Stage3(
	original entities.Layout,
	prevLayout *entities.Layout,
	template entities.Template,
	prevGrid *entities.Grid,
) (*entities.Layout, *entities.Grid, error) {
	var out entities.Layout
	var stageComponents []entities.DesignComponent
	for _, c := range prevLayout.Components {
		if !prevGrid.HaveColisionInList(c.Positions, c.ID) {
			stageComponents = append(stageComponents, c)
			continue
		}
		positions, err := prevGrid.FindFreePositionsToFitBasedOnPivot(
			c.Pivot,
			c.InnerContainer,
		)
		if err != nil {
			continue
		}
		cont := prevGrid.PointsToContainer(positions)
		prevGrid.RemoveFromAllCells(c.ID)
		prevGrid.OcupyByPointList(positions, c.ID)
		c.Positions = positions
		c.ScaleToFitInSize(cont.Width(), cont.Height())
		c.MoveTo(cont.UpperLeft)
		c.CenterInContainer(cont)
		stageComponents = append(stageComponents, c)
	}
	out.Components = stageComponents
	out.Template = template
	out.DesignID = original.DesignID
	out.Width = template.Width
	out.Height = template.Height
	out.Grid = *prevGrid
	return &out, prevGrid, nil
}

func Stage4(
	original entities.Layout,
	prevLayout *entities.Layout,
	template entities.Template,
	prevGrid *entities.Grid,
) (*entities.Layout, *entities.Grid, error) {
	var out entities.Layout
	var stageComponents []entities.DesignComponent
	stageGrid, _ := entities.NewGrid(
		entities.WithDefault(template.Width, template.Height),
		entities.WithCells(prevGrid.SlotsX, prevGrid.SlotsY),
	)
	for _, c := range prevLayout.Components {
		if !prevGrid.CantItGrow(c.Pivot, c.InnerContainer, c.ID) {
			stageComponents = append(stageComponents, c)
			continue
		}
		cont, err := prevGrid.FindSpaceToGrow(c.Pivot, c.InnerContainer, c.ID)
		if err != nil || cont == nil {
			continue
		}
		prevGrid.OcupyWithContainer(*cont, c.ID)
		c.MoveTo(cont.UpperLeft)
		c.ScaleToFitInSize(cont.Width(), cont.Height())
		c.CenterInContainer(*cont)
		stageComponents = append(stageComponents, c)
	}
	out.Components = stageComponents
	out.Template = template
	out.DesignID = original.DesignID
	out.Width = template.Width
	out.Height = template.Height
	out.Grid = *stageGrid
	return &out, stageGrid, nil
}
