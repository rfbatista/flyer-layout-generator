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
		return original.Components[i].OrderPriority() < original.Components[j].OrderPriority()
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
	// layout3, _, err := Stage3(original, layout2, template, stage2Grid)
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
		positions, err := grid.FindPositionsToFitBasedOnPivot(cell.Position(), c.InnerContainer)
		if err != nil || len(positions) == 0 {
			continue
		}
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
	stageGrid, _ := entities.NewGrid(
		entities.WithDefault(template.Width, template.Height),
		entities.WithCells(prevGrid.SlotsX, prevGrid.SlotsY),
	)
	for _, c := range prevLayout.Components {
		cell := prevGrid.WhereIsId(c.ID)
		if cell == nil {
			continue
		}
		if cell.IsOnlyOcupiedBy(c.ID) {
			stageComponents = append(stageComponents, c)
			stageGrid.OcupyByPositionList(c.Positions, c.ID)
			continue
		}
		positions, err := prevGrid.FindFreePositionsToFitBasedOnPivot(cell.Position(), c.InnerContainer)
		if err != nil {
			continue
		}
		cont := stageGrid.PositionsToContainer(positions)
		stageGrid.OcupyByPositionList(c.Positions, c.ID)
		c.ScaleToFitInSize(cont.Width(), cont.Height())
		c.MoveTo(cont.UpperLeft)
		c.CenterInContainer(cont)
		stageGrid.OcupyByPositionList(positions, c.ID)
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
