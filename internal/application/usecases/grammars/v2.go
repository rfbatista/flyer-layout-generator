package grammars

import (
	"algvisual/internal/domain/entities"
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
	log.Debug("summary of stage 0",
		zap.Int("total of components", len(original.Components)),
	)

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
	log.Debug("summary of stage 1",
		zap.Int("total of components", len(layout1.Components)),
	)

	// *************************************************
	// STAGE 2
	// Position elements in target template grid
	// *************************************************
	inlay1, _ := Clone(layout1)
	layout2, stage2Grid, err := PositionElementsInGrid(original, *inlay1, template, *stage1Grid)
	if err != nil {
		return nil, err
	}
	layout2.Grid = *stage2Grid
	jstage2, err := json.Marshal(layout2)
	if err != nil {
		return nil, err
	}
	stages = append(stages, string(jstage2))

	log.Debug("summary of stage 2",
		zap.Int("total of components", len(layout2.Components)),
	)
	// *************************************************
	// STAGE 3
	// Move elements that have colision
	// *************************************************
	sort.Slice(layout2.Components, func(i, j int) bool {
		it := original.Components[i].Type
		jt := original.Components[j].Type
		return original.Config.Priorities[it] < original.Config.Priorities[jt]
	})
	inlay2, _ := Clone(layout2)
	layout3, stage3Grid, err := StageFindColision(original, *inlay2, template, *stage2Grid)
	if err != nil {
		return nil, err
	}
	layout2.Grid = *stage3Grid
	jstage3, err := json.Marshal(layout3)
	if err != nil {
		return nil, err
	}
	stages = append(stages, string(jstage3))

	log.Debug("summary of stage 3",
		zap.Int("total of components", len(layout3.Components)),
	)

	// *************************************************
	// STAGE 4
	// Expand elements
	// *************************************************
	inlay3, _ := Clone(layout3)
	layout4, stage4Grid, err := ExpandElements(original, inlay3, template, stage3Grid)
	if err != nil {
		return nil, err
	}
	layout4.Grid = *stage4Grid
	jstage4, err := json.Marshal(layout4)
	if err != nil {
		return nil, err
	}
	stages = append(stages, string(jstage4))

	log.Debug("summary of stage 4",
		zap.Int("total of components", len(layout4.Components)),
	)

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
	inlay4, _ := Clone(layout4)
	layout5, stage5Grid, err := StageRemoveColisions(original, *inlay4, template, *stage4Grid)
	if err != nil {
		return nil, err
	}
	layout5.Grid = *stage5Grid
	jstage5, err := json.Marshal(layout5)
	if err != nil {
		return nil, err
	}
	stages = append(stages, string(jstage5))

	log.Debug("summary of stage 5",
		zap.Int("total of components", len(layout5.Components)),
	)

	// *************************************************
	// log.Debug("((((((((((((((((()))))))))))))))))")
	// log.Debug(fmt.Sprintf("grid x: %d, y %d", gridX, gridY))
	// log.Debug("((((((((((((((((()))))))))))))))))")

	if original.Background != nil {
		original.Background.ScaleToFillInSize(template.Width, template.Height)
		original.Background.MoveTo(entities.NewPoint(0, 0))
	}

	layout5.Background = original.Background
	layout5.Stages = stages
	return layout5, nil
}

func Clone(orig *entities.Layout) (*entities.Layout, error) {
	origJSON, err := json.Marshal(orig)
	if err != nil {
		return nil, err
	}

	clone := entities.Layout{}
	if err = json.Unmarshal(origJSON, &clone); err != nil {
		return nil, err
	}

	return &clone, nil
}
