package designs

import (
	"context"

	"algvisual/internal/database"
	"algvisual/internal/shared"
)

type ListComponentsByFileIdRequest struct {
	DesignID int32 `param:"photoshop_id"`
}

type ListComponentsByFileIdResult struct {
	Status string                     `json:"status,omitempty"`
	Data   []database.DesignComponent `json:"data,omitempty"`
}

func ListComponentsByFileIdUseCase(
	ctx context.Context,
	req ListComponentsByFileIdRequest,
	queries *database.Queries,
) (*ListComponentsByFileIdResult, error) {
	res, err := queries.GetComponentsByDesignID(ctx, req.DesignID)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao listar componentes por arquivo", "")
		return nil, err
	}
	return &ListComponentsByFileIdResult{
		Status: "success",
		Data:   res,
	}, nil
}
