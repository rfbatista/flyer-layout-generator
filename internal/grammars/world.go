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
	Confi          entities.LayoutRequestConfig
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
			//return world, prancheta, nil
			continue
		}
		layout.Components = append(layout.Components, *c)
		_, layout = PositionComponent(world, layout, c.ID)
		_, layout = ScaleComponent(world, layout, c.ID)
		_, layout = SnapComponent(world, layout, c.ID)
		_, layout = CenterComponent(world, layout, c.ID)
	}
	// s := rand.NewSource(time.Now().Unix())
	// r := rand.New(s) // initialize local pseudorandom generator
	// candidates := filter(layout.Components, func(component entities.DesignComponent) bool {
	// 	if component.Width > layout.Width/2 {
	// 		return false
	// 	}
	// 	if component.Height > layout.Height {
	// 		return false
	// 	}
	// 	return true
	// })
	// grid, _ := entities.NewGrid(entities.WithDefault(layout.Width, layout.Height),
	// 	entities.WithPivot(candidates[r.Intn(len(candidates))].Width, candidates[r.Intn(len(candidates))].Height),
	// )
	// ngrid, _ := CloneGrid(grid)
	// layout.Grid = *ngrid
	// var components []entities.DesignComponent
	// for idx := range layout.Components {
	// 	if len(grid.Regions) == 0 {
	// 		continue
	// 	}
	// 	_, layout = CalculateGap(world, layout, layout.Components[idx].ID)
	// 	_, layout = SnapGridComponent(world, layout, layout.Components[idx].ID, grid)
	// 	ent := layout.Components[idx]
	// 	if ent.Type != "modelo" {
	// 		grid.RemoveAllRegionsInThisPosition(ent.Xi, ent.Yi, ent.Xii, ent.Yii)
	// 	}
	// 	components = append(components, layout.Components[idx])
	// }
	// layout.Components = components
	origJSON, _ := json.Marshal(layout)
	fmt.Println(string(origJSON))
	return world, layout, nil
}

func filter(ss []entities.DesignComponent, test func(component entities.DesignComponent) bool) (ret []entities.DesignComponent) {
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
