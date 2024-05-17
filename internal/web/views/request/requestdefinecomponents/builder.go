package requestdefinecomponents

import (
	"context"

	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/usecases/componentusecase"
)

type pageRequest struct {
	DesignID int32 `json:"id,omitempty"`
}

func PagePropsAssembler(
	ctx context.Context,
	db *database.Queries,
	req pageRequest,
) (*PageProps, error) {
	el, err := db.GetElements(ctx, req.DesignID)
	if err != nil {
		return nil, err
	}
	comps, err := componentusecase.GetComponentsByDesignIdUseCase(
		ctx,
		componentusecase.GetComponentsByDesignIdRequest{ID: req.DesignID},
		db,
	)
	if err != nil {
		return nil, err
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
	var background entities.DesignComponent
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
	return &PageProps{
		Components: comps.Components,
		Elements:   database.ToDesignElementEntitieList(el),
		Background: background,
	}, nil
}

func RemoveDesignComponentIndex(
	s []entities.DesignComponent,
	index int,
) []entities.DesignComponent {
	return append(s[:index], s[index+1:]...)
}

func RemoveIndex(s []database.DesignElement, index int) []database.DesignElement {
	return append(s[:index], s[index+1:]...)
}
