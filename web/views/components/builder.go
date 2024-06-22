package components

import (
	"algvisual/internal/database"
	"algvisual/internal/designs"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"
	"fmt"
)

type pageRequest struct {
	DesignID int32 `param:"design_id" json:"id,omitempty"`
}

func Props(
	ctx context.Context,
	db *database.Queries,
	req pageRequest,
) (PageProps, error) {
	var props PageProps
	el, err := db.GetElements(ctx, req.DesignID)
	if err != nil {
		return props, err
	}
	fmt.Println(req.DesignID)
	comps, err := designs.GetComponentsByDesignIdUseCase(
		ctx,
		designs.GetComponentsByDesignIdRequest{ID: req.DesignID},
		db,
	)
	if err != nil {
		return props, err
	}
	for _, c := range comps.Components {
		for _, ce := range c.Elements {
			for idx, e := range el {
				if e.ID == ce.ID {
					el = RemoveIndex(el, idx)
				}
			}
		}
	}
	var background entities.LayoutComponent
	n := 0
	for _, c := range comps.Components {
		if c.Type == string(database.ComponentTypeBackground) {
			background = c
		} else {
			comps.Components[n] = c
			n++
		}
	}
	comps.Components = comps.Components[:n]
	props.Components = comps.Components
	props.Elements = mapper.ToDesignElementEntitieList(el)
	props.Background = background
	return props, nil
}

func RemoveDesignComponentIndex(
	s []entities.LayoutComponent,
	index int,
) []entities.LayoutComponent {
	return append(s[:index], s[index+1:]...)
}

func RemoveIndex(s []database.LayoutElement, index int) []database.LayoutElement {
	return append(s[:index], s[index+1:]...)
}
