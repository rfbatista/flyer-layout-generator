package componentusecase

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

	"algvisual/internal/database"
	"algvisual/internal/entities"
)

type GetComponentsByDesignIdRequest struct {
	ID int32 `json:"id,omitempty"`
}

type GetComponentsByDesignIdResult struct {
	Components []entities.DesignComponent `json:"components,omitempty"`
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
	var components []entities.DesignComponent
	for _, c := range comps {
		comp := database.TodesignComponentEntitie(c)
		el, err := db.GetDesignElementsByComponentID(ctx, pgtype.Int4{Int32: c.ID, Valid: true})
		if err != nil {
			return nil, err
		}
		comp.Elements = database.ToDesignElementEntitieList(el)
		components = append(components, comp)
	}
	return &GetComponentsByDesignIdResult{Components: components}, nil
}
