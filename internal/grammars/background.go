package grammars

import "algvisual/internal/entities"

func PositionBackground(
	world World,
	prancheta entities.Layout,
) (World, entities.Layout) {
	ent := prancheta.Background
	if ent == nil {
		return world, prancheta
	}
	widthRatio := float64(prancheta.Width) / float64(ent.BboxWidth())
	heightRatio := float64(prancheta.Height) / float64(ent.BboxHeigth())
	if widthRatio > heightRatio {
		ent.ScaleTo(widthRatio, widthRatio)
	} else {
		ent.ScaleTo(heightRatio, heightRatio)
	}
	var x, y int32
	if prancheta.Width-ent.Width != 0 {
		x = (prancheta.Width - ent.Width) / 2
	}
	if prancheta.Height-ent.Height != 0 {
		y = (prancheta.Height - ent.Height) / 2
	}
	ent.SetPosition(x, y)
	return world, prancheta
}
