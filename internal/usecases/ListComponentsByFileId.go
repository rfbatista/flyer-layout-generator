package usecases

import (
	"context"

	"algvisual/internal/database"
	"algvisual/internal/shared"
)

type ListComponentsByFileIdRequest struct {
	PhotoshopID int32 `param:"photoshop_id"`
}

type ListComponentsByFileIdResult struct {
	Status string                        `json:"status,omitempty"`
	Data   []database.PhotoshopComponent `json:"data,omitempty"`
}

func ListComponentsByFileIdUseCase(
	ctx context.Context,
	req ListComponentsByFileIdRequest,
	queries *database.Queries,
) (*ListComponentsByFileIdResult, error) {
	res, err := queries.ListComponentByFileId(ctx, req.PhotoshopID)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao listar componentes por arquivo", "")
		return nil, err
	}
	return &ListComponentsByFileIdResult{
		Status: "success",
		Data:   res,
	}, nil
}
