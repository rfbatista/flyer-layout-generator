package grammars

import (
	"algvisual/internal/entities"
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

type World struct {
	OriginalDesign entities.DesignFile
	Components     []entities.DesignComponent
	Elements       []entities.DesignElement
	PivotWidth     int32
	PivotHeight    int32
	TwistedDesign  entities.Layout
	Config         entities.LayoutRequestConfig
}

type Grammar func(world World, prancheta entities.Layout) (*World, *entities.Layout, error)

func Run(
	world World,
	layout entities.Layout,
	log *zap.Logger,
) (World, entities.Layout, error) {
	world.TwistedDesign = TwistDesign(world, layout, log)
	for i := 0; i <= len(world.Components); i++ {
		c := Pick(world, layout, log)
		if c == nil {
			// return world, prancheta, nil
			continue
		}
		layout.Components = append(layout.Components, *c)
		_, layout = PositionComponent(world, layout, c.ID)
		_, layout = ScaleComponent(world, layout, c.ID)
		_, layout = SnapComponent(world, layout, c.ID)
		_, layout = CenterComponent(world, layout, c.ID)
	}
	if layout.Background != nil {
		_, layout = PositionBackground(world, layout)
	}

	if true {
		layout.Grid = world.Config.Grid
		for idx := range layout.Components {
			ent := &layout.Components[idx]
			if len(world.Config.Grid.Cells) == 0 {
				continue
			}
			if world.Config.KeepProportions {
				_, layout = RepositonButKeepProportions(world, layout, ent.ID)
			} else {
				_, layout = ResizeToFitInRegion(world, layout, ent.ID)
			}
			if ent.Type != "modelo" {
				world.Config.Grid.RemoveAllRegionsInThisPosition(ent.Xi, ent.Yi, ent.Xii, ent.Yii)
			}
		}
	}
	origJSON, _ := json.Marshal(layout)
	fmt.Println(string(origJSON))
	log.Debug(fmt.Sprintf("regions %d components in %d", len(layout.Grid.Cells), len(layout.Components)))
	return world, layout, nil
}

func filter(
	ss []entities.DesignComponent,
	test func(component entities.DesignComponent) bool,
) (ret []entities.DesignComponent) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func CloneGrid(orig *entities.Grid) (*entities.Grid, error) {
	origJSON, err := json.Marshal(orig)
	if err != nil {
		return nil, err
	}

	clone := entities.Grid{}
	if err = json.Unmarshal(origJSON, &clone); err != nil {
		return nil, err
	}

	return &clone, nil
}
