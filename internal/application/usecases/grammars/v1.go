package grammars

import (
	"algvisual/internal/domain/entities"
	"fmt"
	"sort"

	"go.uber.org/zap"
)

func RunV1(
	original entities.Layout,
	template entities.Template,
	gridX, gridY int32,
	log *zap.Logger,
) (*entities.Layout, error) {
	var out entities.Layout
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
	log.Debug("starting stage 1")
	var stage1components []entities.LayoutComponent
	for _, c := range original.Components {
		cell := grid.WhereIsPoint(c.Center())
		if cell == nil {
			continue
		}
		c.Pivot = cell.Position()
		cell.Ocupy(c.ID)
		stage1components = append(stage1components, c)
	}
	sort.Slice(stage1components, func(i, j int) bool {
		return stage1components[i].OrderPriority() < stage1components[j].OrderPriority()
	})

	log.Debug("starting stage 2")
	var stage2components []entities.LayoutComponent
	if err != nil {
		return nil, err
	}
	for _, c := range stage1components {
		cell := grid.WhereIsId(c.ID)
		if cell == nil {
			continue
		}
		c.MoveTo(cell.UpLeft())
		positions, err := grid.FindFreePositionsToFitBasedOnPivot(cell.Position(), c.InnerContainer)
		if err != nil {
			continue
		}
		grid.OcupyByPositionList(positions, c.ID)
		c.Positions = positions
		cont := grid.PositionsToContainer(positions)
		c.ScaleToFitInSize(cont.Width(), cont.Height())
		c.MoveTo(cont.UpperLeft)
		stage2components = append(stage2components, c)
	}

	log.Debug("starting stage 3")
	var stage3components []entities.LayoutComponent
	stage3grid, _ := entities.NewGrid(
		entities.WithDefault(template.Width, template.Height),
		entities.WithCells(gridX, gridY),
	)
	for _, c := range stage2components {
		if len(c.Positions) == 0 {
			continue
		}
		cont := stage3grid.PositionsToContainer(c.Positions)
		stage3grid.OcupyByPositionList(c.Positions, c.ID)
		c.ScaleToFitInSize(cont.Width(), cont.Height())
		c.MoveTo(cont.UpperLeft)
		c.CenterInContainer(cont)
		stage3components = append(stage3components, c)
	}

	log.Debug("starting stage 4")
	var stage4components []entities.LayoutComponent
	for _, c := range stage3components {
		if len(c.Positions) == 0 {
			continue
		}
		if c.Type == "oferta" {
			fmt.Println("here")
		}
		cont, err := stage3grid.FindSpaceToGrow(c.Pivot, c.InnerContainer, c.ID)
		if err != nil || cont == nil {
			stage4components = append(stage4components, c)
			continue
		}
		c.MoveTo(cont.UpperLeft)
		c.ScaleToFitInSize(cont.Width(), cont.Height())
		c.CenterInContainer(*cont)
		stage4components = append(stage4components, c)
	}
	log.Debug("stages finished")
	out.Components = stage4components
	out.Template = template
	out.DesignID = original.DesignID
	out.Width = template.Width
	out.Height = template.Height
	out.Grid = *stage3grid
	return &out, nil
}
