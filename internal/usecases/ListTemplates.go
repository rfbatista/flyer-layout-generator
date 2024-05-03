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
	Data []database.Template `json:"data,omitempty"`
}

func ListTemplatesUseCase(
	ctx context.Context,
	req ListTemplatesUseCaseRequest,
	queries *database.Queries,
	log *zap.Logger,
) (*ListTemplatesUseCaseResult, error) {
	limit := req.Limit
	if limit == 0 {
		limit = 10
	}
	result, err := queries.ListTemplates(ctx, database.ListTemplatesParams{
		Limit:  int32(limit),
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
