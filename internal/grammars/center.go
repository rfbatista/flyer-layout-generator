package grammars

import (
	"algvisual/internal/entities"
)

func CenterComponent(
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
			xcenter := int32(c.FWidth/2) + c.Xi
			ycenter := int32(c.FHeight/2) + c.Yi
			xedge := xcenter - int32(ent.FWidth/2)
			yedge := ycenter - int32(ent.FHeight/2)
			if xedge < 0 && c.Type != "modelo" && ent.Xsnaped == false {
				xedge = 5
			}
			if yedge < 0 && c.Type != "modelo" && ent.Ysnaped == false {
				yedge = 5
			}
			if xedge+ent.FWidth > prancheta.Width && c.Type != "modelo" && ent.Xsnaped == false {
				xedge = xedge - (xedge + ent.FWidth - prancheta.Width) - 5
			}
			if yedge+ent.FHeight > prancheta.Height && c.Type != "modelo" && ent.Ysnaped == false {
				yedge = yedge - ((yedge + ent.FHeight) - prancheta.Height) - 5
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
