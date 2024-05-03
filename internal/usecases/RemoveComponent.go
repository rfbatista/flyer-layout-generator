package usecases

import (
	"context"

	"algvisual/internal/database"
	"algvisual/internal/shared"
)

type RemoveComponentUseCaseRequest struct {
	PhotoshopID int32   `param:"photoshop_id" json:"photoshop_id,omitempty"`
	Elements    []int32 `                      json:"elements,omitempty"     body:"elements"`
}

type RemoveComponentUseCaseResult struct {
	Status string                      `json:"status,omitempty"`
	Data   []database.PhotoshopElement `json:"data,omitempty"`
}

func RemoveComponentUseCase(
	ctx context.Context,
	queries *database.Queries,
	req RemoveComponentUseCaseRequest,
) (*RemoveComponentUseCaseResult, error) {
	elUpdated, err := queries.RemoveComponentFromElements(
		ctx,
		database.RemoveComponentFromElementsParams{
			PhotoshopID: req.PhotoshopID,
			Ids:         req.Elements,
		},
	)
	if err != nil {
		return nil, shared.WrapWithAppError(err, "Falha ao atualizar elemento do photoshop", "")
	}
	return &RemoveComponentUseCaseResult{
		Status: "success",
		Data:   elUpdated,
	}, nil
}
