package usecases

import (
	"context"

	"algvisual/internal/database"
	"algvisual/internal/shared"
)

type RemoveComponentUseCaseRequest struct {
	PhotoshopID int   `params:"PhotoshopID" json:"photoshop_id,omitempty"`
	Elements    []int `                     json:"elements,omitempty"     body:"elements"`
}

type RemoveComponentUseCaseResult struct {
	Data []database.PhotoshopElement
}

func RemoveComponentUseCase(
	ctx context.Context,
	queries *database.Queries,
	req RemoveComponentUseCaseRequest,
) (*RemoveComponentUseCaseResult, error) {
	elUpdated, err := queries.UpdateManyPhotoshopElement(
		ctx,
		database.UpdateManyPhotoshopElementParams{
			PhotoshopID:         int32(req.PhotoshopID),
			ComponentIDDoUpdate: true,
			ComponentID:         0,
		},
	)
	if err != nil {
		return nil, shared.WrapWithAppError(err, "Falha ao atualizar elemento do photoshop", "")
	}
	return &RemoveComponentUseCaseResult{
		Data: elUpdated,
	}, nil
}
