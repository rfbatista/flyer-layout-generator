package designs

import (
	"context"

	"algvisual/internal/database"
	"algvisual/internal/shared"
)

type RemoveComponentUseCaseRequest struct {
	DesignID int32   `param:"design_id" json:"photoshop_id,omitempty"`
	Elements []int32 `                  json:"elements,omitempty"     form:"elements" body:"elements"`
}

type RemoveComponentUseCaseResult struct {
	Status string                   `json:"status,omitempty"`
	Data   []database.DesignElement `json:"data,omitempty"`
}

func RemoveComponentUseCase(
	ctx context.Context,
	queries *database.Queries,
	req RemoveComponentUseCaseRequest,
) (*RemoveComponentUseCaseResult, error) {
	elUpdated, err := queries.RemoveComponentFromElements(
		ctx,
		database.RemoveComponentFromElementsParams{
			DesignID: req.DesignID,
			Ids:      req.Elements,
		},
	)
	if err != nil {
		return nil, shared.WrapWithAppError(err, "Falha ao atualizar elemento do photoshop", "")
	}
	err = queries.ClearEmptyComponents(ctx)
	if err != nil {
		return nil, shared.WrapWithAppError(err, "Falha ao limpar componentes vazios", "")
	}
	return &RemoveComponentUseCaseResult{
		Status: "success",
		Data:   elUpdated,
	}, nil
}
