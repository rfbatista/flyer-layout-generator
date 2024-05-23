package grammars

import "algvisual/internal/entities"

func CenterComponent(
	world World,
	prancheta entities.Prancheta,
	id int32,
) (World, entities.Prancheta) {
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
			xcenter := int32(c.Width/2) + c.Xi
			ycenter := int32(c.Height/2) + c.Yi
			xedge := xcenter - int32(ent.Width/2)
			yedge := ycenter - int32(ent.Height/2)
			if xedge < 0 && c.Type != "modelo" {
				xedge = 5
			}
			if yedge < 0 && c.Type != "modelo" {
				yedge = 5
			}
			if xedge+ent.Width > prancheta.Width && c.Type != "modelo" {
				xedge = xedge - (xedge + ent.Width - prancheta.Width) - 5
			}
			if yedge+ent.Height > prancheta.Height && c.Type != "modelo" {
				yedge = yedge - ((yedge + ent.Height) - prancheta.Height) - 5
			}
			ent.SetPosition(xedge, yedge)
		}
	}
	for idx := range prancheta.Components {
		if prancheta.Components[idx].ID == id {
			prancheta.Components[idx] = *ent
		}
	}
	return world, prancheta
}
