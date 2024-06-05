package grammars

import (
	"algvisual/internal/entities"
	"sort"
)

func RunV1(
	original entities.Layout,
	template entities.Template,
	gridX, gridY int32,
) (entities.Layout, error) {
	var out entities.Layout
	grid, err := entities.NewGrid(
		entities.WithDefault(original.Width, original.Height),
		entities.WithCells(gridX, gridY),
	)
	if err != nil {
		return out, err
	}
	sort.Slice(original.Components, func(i, j int) bool {
		return original.Components[i].OrderPriority() < original.Components[j].OrderPriority()
	})
	var stage1components []entities.DesignComponent
	for _, c := range original.Components {
		cell := grid.WhereIsPoint(c.Center())
		if cell == nil {
			continue
		}
		cell.Ocupy(c.ID)
		stage1components = append(stage1components, c)
	}
	sort.Slice(stage1components, func(i, j int) bool {
		return stage1components[i].OrderPriority() < stage1components[j].OrderPriority()
	})
	var stage2components []entities.DesignComponent
	// stage2grid, err := entities.NewGrid(
	// 	entities.WithDefault(original.Width, original.Height),
	// 	entities.WithCells(gridX, gridY),
	// )
	if err != nil {
		return out, err
	}
	for _, c := range stage1components {
		cell := grid.WhereIsId(c.ID)
		if cell == nil {
			continue
		}
		positions, err := grid.FindPositionsToFitBasedOnPivot(cell.Position(), c.InnerContainer)
		if err != nil {
			continue
		}
		cont := grid.PositionsToContainer(positions)
		c.MoveTo(cont.UpperLeft)
		c.ScaleToFitInSize(cont.Width(), cont.Height())
		stage2components = append(stage2components, c)
	}
	out.Components = stage2components
	out.Template = template
	out.DesignID = original.DesignID
	out.Width = original.Width
	out.Height = original.Height
	out.Grid = *grid
	return out, nil
}
