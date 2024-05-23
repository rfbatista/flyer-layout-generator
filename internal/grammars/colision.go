package grammars

import "algvisual/internal/entities"

func RemoveColision(
	world World,
	prancheta entities.Prancheta,
	id int32,
) (World, entities.Prancheta) {
	var ent *entities.DesignComponent
	for _, c := range world.Components {
		if c.ID == id {
			ent = &c
		}
	}
	if ent == nil {
		return world, prancheta
	}
	for x := ent.Xi; x <= ent.Xii; x++ {
		for y := ent.Yi; y <= ent.Yii; y++ {
			c := isCompIn(prancheta.Components, int(x), int(y), id)
			if c != nil {
				return world, prancheta
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

func isCompIn(comps []entities.DesignComponent, x, y int, fromID int32) *entities.DesignComponent {
	for c := range comps {
		if comps[c].ID == fromID {
			continue
		}
		if comps[c].Type == "modelo" {
			continue
		}
		if x >= int(comps[c].Xi) && x <= int(comps[c].Xii) && y >= int(comps[c].Yi) && y <= int(comps[c].Yii) {
			return &comps[c]
		}
	}
	return nil
}
