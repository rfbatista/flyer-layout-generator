package grammars

import (
	"algvisual/internal/entities"
)

func ScaleComponent(
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
	wprorp, hprop := getOriginalProportion(world, id)
	if wprorp > hprop && ent.FWidth < ent.FHeight {
		// calcula a escala com base na proporção do elemento no design original
		nwidth := float64(prancheta.Width) * wprorp
		scaleTo := (nwidth / float64(ent.FWidth))
		ent.ScaleWithoutMoving(scaleTo, scaleTo)
	} else {
		// calcula a escala com base na proporção do elemento no design original
		nheight := float64(prancheta.Height) * hprop
		scaleTo := (nheight / float64(ent.FHeight))
		ent.ScaleWithoutMoving(scaleTo, scaleTo)
	}
	if doesItFit(&prancheta, ent) {
		for idx, c := range prancheta.Components {
			if c.ID == id {
				prancheta.Components[idx] = *ent
			}
		}
		return world, prancheta
	}
	return world, prancheta
}

func getOriginalProportion(world World, id int32) (float64, float64) {
	var ent *entities.DesignComponent
	for _, c := range world.Components {
		if c.ID == id {
			ent = &c
		}
	}
	if ent == nil {
		return 1, 1
	}
	return float64(
			ent.FWidth) / float64(world.OriginalDesign.Width), float64(
			ent.FHeight) / float64(world.OriginalDesign.Height)
}

func doesItFit(pr *entities.Layout, ent *entities.DesignComponent) bool {
	return true
}
