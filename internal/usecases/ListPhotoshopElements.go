package usecases

import (
	"context"
	"errors"

	"algvisual/internal/database"
)

type ListPhotoshopElementsUseCaseRequest struct {
	Limit       int `query:"limit" json:"limit,omitempty"`
	Skip        int `query:"skip"  json:"skip,omitempty"`
	PhotoshopID int `              json:"photoshop_id,omitempty" param:"photoshop_id"`
}

type ListPhotoshopElementsUseCaseResult struct {
	Status string                   `json:"status,omitempty"`
	Data   []database.DesignElement `json:"data,omitempty"`
}

func ListPhotoshopElementsUseCase(
	ctx context.Context,
	req ListPhotoshopElementsUseCaseRequest,
	db *database.Queries,
) (*ListPhotoshopElementsUseCaseResult, error) {
	res, err := db.ListPhotoshopElements(ctx, database.ListPhotoshopElementsParams{
		PhotoshopID: int32(req.PhotoshopID),
		Limit:       int32(req.Limit),
		Offset:      int32(req.Skip),
	})
	if err != nil {
		return nil, errors.Join(err, errors.New("falha ai listar elementos do photoshop"))
	}
	return &ListPhotoshopElementsUseCaseResult{
		Status: "success",
		Data:   res,
	}, nil
}
