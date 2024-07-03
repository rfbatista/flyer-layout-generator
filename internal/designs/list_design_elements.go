package designs

import (
	"algvisual/database"
	"context"
	"errors"
)

type ListDesignElementsUseCaseRequest struct {
	Limit       int `query:"limit" json:"limit,omitempty"`
	Skip        int `query:"skip"  json:"skip,omitempty"`
	PhotoshopID int `              json:"photoshop_id,omitempty" param:"photoshop_id"`
}

type ListDesignElementsUseCaseResult struct {
	Status string                   `json:"status,omitempty"`
	Data   []database.LayoutElement `json:"data,omitempty"`
}

func ListDesignElementsUseCase(
	ctx context.Context,
	req ListDesignElementsUseCaseRequest,
	db *database.Queries,
) (*ListDesignElementsUseCaseResult, error) {
	res, err := db.ListdesignElements(ctx, database.ListdesignElementsParams{
		DesignID: int32(req.PhotoshopID),
		Limit:    int32(req.Limit),
		Offset:   int32(req.Skip),
	})
	if err != nil {
		return nil, errors.Join(err, errors.New("falha ai listar elementos do photoshop"))
	}
	return &ListDesignElementsUseCaseResult{
		Status: "success",
		Data:   res,
	}, nil
}
