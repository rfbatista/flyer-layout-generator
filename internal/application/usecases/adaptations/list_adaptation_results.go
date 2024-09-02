package adaptations

import (
	"algvisual/internal/domain/entities"
	"context"
)

type ListAdaptationResultsInput struct {
	AdaptationID int32 `param:"adaptation_id"`
}

type ListAdaptationResultsOutput struct {
	Data []entities.Layout `json:"data"`
}

func ListAdaptationResultsUseCase(
	ctx context.Context,
	req ListAdaptationResultsInput,
) (*ListAdaptationResultsOutput, error) {
	return &ListAdaptationResultsOutput{}, nil
}
