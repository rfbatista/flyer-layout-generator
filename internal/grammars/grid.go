package grammars

import (
	"algvisual/internal/entities"
)

func ResizeToFitInRegion(
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
	if len(world.Config.Grid.Regions) == 0 {
		prancheta.Components = filter(
			prancheta.Components,
			func(component entities.DesignComponent) bool {
				return component.ID == id
			},
		)
		return world, prancheta
	}
	// regions, _ := world.Config.Grid.WhereToSnap(*ent)
	regions := world.Config.Grid.FindOverlappingRegions(*ent)
	if len(regions) == 0 {
		return world, prancheta
	}
	var region entities.Region
	var size int32
	for _, r := range regions {
		if r.Overlap > size {
			region = r.Region
			size = r.Overlap
		}
	}
	ent.ScaleToFitInSize(
		region.Width()-world.Config.Padding*2,
		region.Height()-world.Config.Padding*2,
	)
	ent.CenterInRegion(region)
	for idx := range prancheta.Components {
		if prancheta.Components[idx].ID == id {
			prancheta.Components[idx] = *ent
		}
	}
	return world, prancheta
}

func RepositonButKeepProportions(
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
	if len(world.Config.Grid.Regions) == 0 {
		prancheta.Components = filter(
			prancheta.Components,
			func(component entities.DesignComponent) bool {
				return ent.ID != component.ID
			},
		)
		return world, prancheta
		return world, prancheta
	}
	if ent.Type == "modelo" {
		return world, prancheta
	}
	// region, _ := g.WhereToSnap(*ent)
	var regions []entities.Region
	for _, region := range world.Config.Grid.Regions {
		if region.Xi > ent.Xii || region.Xii < ent.Xi {
			continue
		}
		if region.Yi > ent.Yii || region.Yii < ent.Yi {
			continue
		}
		regions = append(regions, region)
	}
	if len(regions) == 0 {
		prancheta.Components = filter(
			prancheta.Components,
			func(component entities.DesignComponent) bool {
				return ent.ID != component.ID
			},
		)
		return world, prancheta
	}
	xi := regions[0].Xi
	yi := regions[0].Yi
	xii := regions[0].Xii
	yii := regions[0].Yii
	for _, e := range regions {
		if xi > e.Xi {
			xi = e.Xi
		}
		if yi > e.Yi {
			yi = e.Yi
		}
		if xii < e.Xii {
			xii = e.Xi
		}
		if yii > e.Yii {
			yii = e.Yii
		}
	}
	ent.CenterInRegion(*entities.NewRegion(xi, yi, xii, yii))
	for idx := range prancheta.Components {
		if prancheta.Components[idx].ID == id {
			prancheta.Components[idx] = *ent
		}
	}
	return world, prancheta
}
