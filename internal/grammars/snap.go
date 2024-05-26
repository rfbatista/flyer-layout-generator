package grammars

import (
	"algvisual/internal/entities"
)

func SnapComponent(
	world World,
	prancheta entities.Layout,
	id int32,
) (World, entities.Layout) {
	var ent *entities.DesignComponent
	for _, c := range prancheta.Components {
		if c.ID == id {
			ent = &c
		}
	}
	if ent == nil {
		return world, prancheta
	}
	for _, c := range world.TwistedDesign.Components {
		if c.ID == id {
			if c.Xi < 10 {
				ent.SetPosition(c.Xi, ent.Yi)
				ent.Xsnaped = true
			}
			if c.Xii > world.TwistedDesign.Width-10 {
				ent.SetPosition(c.Xii-ent.Width, ent.Yi)
				ent.Xsnaped = true
			}
			if c.Yi < 10 {
				ent.SetPosition(ent.Xi, c.Yi)
				ent.Ysnaped = true
			}
			if c.Yii > world.TwistedDesign.Height-10 {
				ent.SetPosition(ent.Xi, c.Yii-ent.Height)
				ent.Ysnaped = true
			}
		}
	}
	for idx := range prancheta.Components {
		if prancheta.Components[idx].ID == id {
			prancheta.Components[idx] = *ent
		}
	}
	return world, prancheta
}
