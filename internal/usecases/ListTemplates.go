package usecases

import (
	"context"

	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/shared"
)

type ListTemplatesUseCaseRequest struct {
	Limit int `query:"limit" json:"limit,omitempty"`
	Skip  int `query:"skip"  json:"skip,omitempty"`
}

type ListTemplatesUseCaseResult struct {
	Data []database.ListTemplatesRow `json:"data,omitempty"`
}

func ListTemplatesUseCase(
	ctx context.Context,
	req ListTemplatesUseCaseRequest,
	queries *database.Queries,
	log *zap.Logger,
) (*ListTemplatesUseCaseResult, error) {
	result, err := queries.ListTemplates(ctx, database.ListTemplatesParams{
		Limit:  int32(req.Limit),
		Offset: int32(req.Skip),
	})
	if err != nil {
		err = shared.WrapWithAppError(err, "failed to list templates", "")
		log.Error(err.Error())
		return nil, err
	}
	return &ListTemplatesUseCaseResult{
		Data: result,
	}, nil
}
