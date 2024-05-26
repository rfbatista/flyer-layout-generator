package grammars

import (
	"algvisual/internal/entities"
	"encoding/json"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

type World struct {
	OriginalDesign entities.DesignFile
	Components     []entities.DesignComponent
	Elements       []entities.DesignElement
	PivotWidth     int32
	PivotHeight    int32
	TwistedDesign  entities.Layout
}

type Grammar func(world World, prancheta entities.Layout) (*World, *entities.Layout, error)

func Run(
	world World,
	prancheta entities.Layout,
	log *zap.Logger,
) (World, entities.Layout, error) {
	world.TwistedDesign = TwistDesign(world, prancheta, log)
	for i := 0; i <= len(world.Components); i++ {
		c := Pick(world, prancheta, log)
		if c == nil {
			//return world, prancheta, nil
			continue
		}
		prancheta.Components = append(prancheta.Components, *c)
		_, prancheta = PositionComponent(world, prancheta, c.ID)
		_, prancheta = ScaleComponent(world, prancheta, c.ID)
		_, prancheta = SnapComponent(world, prancheta, c.ID)
		_, prancheta = CenterComponent(world, prancheta, c.ID)
	}
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	candidates := filter(prancheta.Components, func(component entities.DesignComponent) bool {
		if component.Width > prancheta.Width/2 {
			return false
		}
		if component.Height > prancheta.Height {
			return false
		}
		return true
	})
	grid, _ := entities.NewGrid(entities.WithDefault(prancheta.Width, prancheta.Height),
		entities.WithPivot(candidates[r.Intn(len(candidates))].Width, candidates[r.Intn(len(candidates))].Height),
	)
	ngrid, _ := CloneGrid(grid)
	prancheta.Grid = *ngrid
	var components []entities.DesignComponent
	for idx := range prancheta.Components {
		if len(grid.Regions) == 0 {
			continue
		}
		_, prancheta = CalculateGap(world, prancheta, prancheta.Components[idx].ID)
		_, prancheta = SnapGridComponent(world, prancheta, prancheta.Components[idx].ID, grid)
		ent := prancheta.Components[idx]
		if ent.Type != "modelo" {
			grid.RemoveAllRegionsInThisPosition(ent.Xi, ent.Yi, ent.Xii, ent.Yii)
		}
		components = append(components, prancheta.Components[idx])
	}
	prancheta.Components = components
	return world, prancheta, nil
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
