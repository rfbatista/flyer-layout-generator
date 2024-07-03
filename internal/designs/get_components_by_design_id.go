package designs

import (
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"

	"algvisual/database"
)

type GetComponentsByDesignIdRequest struct {
	ID int32 `json:"id,omitempty"`
}

type GetComponentsByDesignIdResult struct {
	Components []entities.LayoutComponent `json:"components,omitempty"`
}

func GetComponentsByDesignIdUseCase(
	ctx context.Context,
	req GetComponentsByDesignIdRequest,
	db *database.Queries,
) (*GetComponentsByDesignIdResult, error) {
	comps, err := db.GetComponentsByDesignID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	var components []entities.LayoutComponent
	for _, c := range comps {
		comp := mapper.TodesignComponentEntitie(c)
		el, err := db.GetDesignElementsByComponentID(ctx, pgtype.Int4{Int32: c.ID, Valid: true})
		if err != nil {
			return nil, err
		}
		comp.Elements = mapper.ToDesignElementEntitieList(el)
		components = append(components, comp)
	}
	return &GetComponentsByDesignIdResult{Components: components}, nil
}
