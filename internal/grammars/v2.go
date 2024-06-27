package grammars

import (
	"algvisual/internal/entities"
	"encoding/json"
	"sort"

	"go.uber.org/zap"
)

func RunV2(
	original entities.Layout,
	template entities.Template,
	gridX, gridY int32,
	log *zap.Logger,
) (*entities.Layout, error) {
	var stages []string
	grid, err := entities.NewGrid(
		entities.WithDefault(original.Width, original.Height),
		entities.WithCells(gridX, gridY),
	)
	if err != nil {
		return nil, err
	}

	// *************************************************
	// STAGE 1
	// Find cells for each component in original design
	// *************************************************
	sort.Slice(original.Components, func(i, j int) bool {
		it := original.Components[i].Type
		jt := original.Components[j].Type
		return original.Config.Priorities[it] < original.Config.Priorities[jt]
	})
	layout1, stage1Grid, err := StageFindCells(original, template, *grid)
	if err != nil {
		return nil, err
	}
	layout1.Grid = *stage1Grid
	jstage1, err := json.Marshal(layout1)
	if err != nil {
		return nil, err
	}
	stages = append(stages, string(jstage1))

	// *************************************************
	// STAGE 2
	// Position elements in target template grid
	// *************************************************
	layout2, stage2Grid, err := PositionElementsInGrid(original, *layout1, template, *stage1Grid)
	if err != nil {
		return nil, err
	}
	layout2.Grid = *stage2Grid
	jstage2, err := json.Marshal(layout2)
	if err != nil {
		return nil, err
	}
	stages = append(stages, string(jstage2))

	// *************************************************
	// STAGE 3
	// Move elements that have colision
	// *************************************************
	sort.Slice(layout2.Components, func(i, j int) bool {
		it := original.Components[i].Type
		jt := original.Components[j].Type
		return original.Config.Priorities[it] > original.Config.Priorities[jt]
	})
	layout3, stage3Grid, err := StageFindColision(original, *layout2, template, *stage2Grid)
	if err != nil {
		return nil, err
	}
	layout2.Grid = *stage3Grid
	jstage3, err := json.Marshal(layout3)
	if err != nil {
		return nil, err
	}
	stages = append(stages, string(jstage3))

	// *************************************************
	// STAGE 4
	// Expand elements
	// *************************************************
	layout4, stage4Grid, err := Stage4(original, layout3, template, stage3Grid)
	if err != nil {
		return nil, err
	}
	layout4.Grid = *stage4Grid
	jstage4, err := json.Marshal(layout4)
	if err != nil {
		return nil, err
	}
	stages = append(stages, string(jstage4))

	// *************************************************
	// STAGE 5
	// Remove colission
	// *************************************************
	sort.Slice(layout4.Components, func(i, j int) bool {
		it := original.Components[i].Type
		jt := original.Components[j].Type
		// Put the one with less priorities to first
		return original.Config.Priorities[it] < original.Config.Priorities[jt]
	})
	layout5, stage5Grid, err := StageRemoveColisions(original, *layout4, template, *stage4Grid)
	if err != nil {
		return nil, err
	}
	layout5.Grid = *stage5Grid
	jstage5, err := json.Marshal(layout5)
	if err != nil {
		return nil, err
	}
	stages = append(stages, string(jstage5))
	// *************************************************

	if original.Background != nil {
		original.Background.ScaleToFillInSize(template.Width, template.Height)
		original.Background.MoveTo(entities.NewPoint(0, 0))
	}

	layout5.Background = original.Background
	layout5.Stages = stages
	return layout5, nil
}
