package grammars

import (
	"algvisual/internal/entities"
)

func SnapGridComponent(
	world World,
	prancheta entities.Layout,
	id int32,
	grid *entities.Grid,
) (World, entities.Layout) {
	var ent *entities.DesignComponent
	for _, c := range prancheta.Components {
		if c.ID == id {
			ent = &c
		}
	}
	if ent == nil || ent.Type == "modelo" {
		return world, prancheta
	}
	regionToSnap, snapToLeft := grid.WhereToSnap(*ent)
	if snapToLeft == true {
		ent.SetPosition(regionToSnap.Xi, regionToSnap.Yi)
		grid.RemoveRegion(regionToSnap)
	} else {
		ent.SetPosition(regionToSnap.Xii-ent.Width, regionToSnap.Yi)
		grid.RemoveRegion(regionToSnap)
	}
	for idx := range prancheta.Components {
		if prancheta.Components[idx].ID == id {
			prancheta.Components[idx] = *ent
		}
	}
	return world, prancheta
}
