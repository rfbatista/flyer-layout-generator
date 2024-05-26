package grammars

import (
	"algvisual/internal/entities"
)

func CalculateGap(
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
	rx, ry := rightGap(prancheta, ent.Xii, ent.Yi, prancheta.Width, ent.Yii, id)
	lx, ly := leftGap(prancheta, 0, ent.Yi, ent.Xi, ent.Yii, id)
	ux, uy := upGap(prancheta, ent.Xi, 0, ent.Xii, ent.Yi, id)
	dx, dy := downGap(prancheta, ent.Xi, ent.Yii, ent.Xii, prancheta.Height, id)

	ent.UpGap = entities.NewPosition(ux, uy)
	ent.RightGap = entities.NewPosition(rx, ry)
	ent.LeftGap = entities.NewPosition(lx, ly)
	ent.DownGap = entities.NewPosition(dx, dy)

	for _, c := range world.TwistedDesign.Components {
		if c.ID == id {
		}
	}
	for idx := range prancheta.Components {
		if prancheta.Components[idx].ID == id {
			prancheta.Components[idx] = *ent
		}
	}
	return world, prancheta
}

func rightGap(p entities.Layout, fromx, fromy, tox, toy, fromID int32) (int32, int32) {
	collisionX := tox
	collisionY := toy
	for i := fromx; i <= tox; i++ {
		for j := fromy; j <= toy; j++ {
			var comp *entities.DesignComponent
			for c := range p.Components {
				if p.Components[c].ID == fromID {
					continue
				}
				if p.Components[c].Type == "modelo" {
					continue
				}
				if i >= p.Components[c].Xi && i <= p.Components[c].Xii && j >= p.Components[c].Yi && j <= p.Components[c].Yii {
					comp = &p.Components[c]
				}
			}
			if comp != nil {
				if i < collisionX {
					collisionX = i
				}
				if j < collisionY {
					collisionY = j
				}
			}
		}
	}
	return collisionX, collisionY
}

func leftGap(p entities.Layout, fromx, fromy, tox, toy, fromID int32) (int32, int32) {
	collisionX := fromx
	collisionY := fromy
	for i := fromx; i <= tox; i++ {
		for j := fromy; j <= toy; j++ {
			var comp *entities.DesignComponent
			for c := range p.Components {
				if p.Components[c].ID == fromID {
					continue
				}
				if p.Components[c].Type == "modelo" {
					continue
				}
				if i >= p.Components[c].Xi && i <= p.Components[c].Xii && j >= p.Components[c].Yi && j <= p.Components[c].Yii {
					comp = &p.Components[c]
				}
			}
			if comp != nil {
				if i > collisionX {
					collisionX = i
				}
				if j > collisionY {
					collisionY = j
				}
			}
		}
	}
	return collisionX, collisionY
}

func upGap(p entities.Layout, fromx, fromy, tox, toy, fromID int32) (int32, int32) {
	collisionX := fromx
	collisionY := fromy
	for i := fromx; i <= tox; i++ {
		for j := fromy; j <= toy; j++ {
			var comp *entities.DesignComponent
			for c := range p.Components {
				if p.Components[c].ID == fromID {
					continue
				}
				if p.Components[c].Type == "modelo" {
					continue
				}
				if i >= p.Components[c].Xi && i <= p.Components[c].Xii && j >= p.Components[c].Yi && j <= p.Components[c].Yii {
					comp = &p.Components[c]
				}
			}
			if comp != nil {
				if i > collisionX {
					collisionX = i
				}
				if j > collisionY {
					collisionY = j
				}
			}
		}
	}
	return collisionX, collisionY
}

func downGap(p entities.Layout, fromx, fromy, tox, toy, fromID int32) (int32, int32) {
	collisionX := tox
	collisionY := toy
	for i := fromx; i <= tox; i++ {
		for j := fromy; j <= toy; j++ {
			var comp *entities.DesignComponent
			for c := range p.Components {
				if p.Components[c].ID == fromID {
					continue
				}
				if p.Components[c].Type == "modelo" {
					continue
				}
				if i >= p.Components[c].Xi && i <= p.Components[c].Xii && j >= p.Components[c].Yi && j <= p.Components[c].Yii {
					comp = &p.Components[c]
				}
			}
			if comp != nil {
				if i < collisionX {
					collisionX = i
				}
				if j < collisionY {
					collisionY = j
				}
			}
		}
	}
	return collisionX, collisionY
}
