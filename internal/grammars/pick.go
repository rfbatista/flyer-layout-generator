package grammars

import (
	"algvisual/internal/entities"
	"go.uber.org/zap"
)

func Pick(world World, prancheta entities.Layout, log *zap.Logger) *entities.DesignComponent {
	for x := 0; x <= int(prancheta.Width); x++ {
		for y := 0; y <= int(prancheta.Height); y++ {
			c := whoIsIn(world, prancheta, world.TwistedDesign.Components, x, y)
			if c != nil {
				choosenComp := *c
				return &choosenComp
			}
		}
	}
	log.Info("no design was picked")
	return nil
}

func whoIsIn(world World, prancheta entities.Layout, comps []entities.DesignComponent, x, y int) *entities.DesignComponent {
	for c := range comps {
		if comps[c].ID == 21 {
		}
		if x >= int(comps[c].Xi) && x <= int(comps[c].Xii) && y >= int(comps[c].Yi) && y <= int(comps[c].Yii) {
			if !isInComp(prancheta.Components, comps[c]) {
				return findCompInWorld(world.Components, comps[c])
			}
		}
	}
	return nil
}

func isInComp(in []entities.DesignComponent, toCheck entities.DesignComponent) bool {
	if toCheck.ID == 21 {
	}
	for _, c := range in {
		if c.ID == toCheck.ID {
			return true
		}
	}
	return false
}

func findCompInWorld(
	d []entities.DesignComponent,
	c entities.DesignComponent,
) *entities.DesignComponent {
	for _, co := range d {
		if co.ID == c.ID {
			return &co
		}
	}
	return nil
}
