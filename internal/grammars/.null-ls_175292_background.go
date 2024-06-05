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
	if prancheta.Width-ent.FWidth != 0 {
		x = (prancheta.Width - ent.FWidth) / 2
	}
	if prancheta.Height-ent.FHeight != 0 {
		y = (prancheta.Height - ent.FHeight) / 2
	}
	ent.SetPosition(x, y)
	return world, prancheta
}
