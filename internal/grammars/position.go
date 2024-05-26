package grammars

import (
	"algvisual/internal/entities"
)

func PositionComponent(
	world World,
	prancheta entities.Layout,
	id int32,
) (World, entities.Layout) {
	var ent *entities.DesignComponent
	for _, c := range world.Components {
		if c.ID == id {
			ent = &c
		}
	}
	if ent == nil {
		return world, prancheta
	}
	for _, c := range world.TwistedDesign.Components {
		if c.ID == id {
			ent.SetPosition(c.Xi, c.Yi)
		}
	}
	for idx := range prancheta.Components {
		if prancheta.Components[idx].ID == id {
			prancheta.Components[idx] = *ent
		}
	}
	return world, prancheta
}
