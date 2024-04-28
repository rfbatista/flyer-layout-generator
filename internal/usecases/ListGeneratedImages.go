package usecases

import (
	"context"

	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/shared"
)

type ListGeneratedImagesRequest struct {
	Limit int `query:"limit" json:"limit,omitempty"`
	Skip  int `query:"skip"  json:"skip,omitempty"`
}

type ListGeneratedImagesResult struct {
	Data []database.Image
}

func ListGeneratedImagesUseCase(
	ctx context.Context,
	req ListGeneratedImagesRequest,
	queries *database.Queries,
	log *zap.Logger,
) (*ListGeneratedImagesResult, error) {
	result, err := queries.ListImagesGenerated(ctx, database.ListImagesGeneratedParams{
		Limit:  int32(req.Limit),
		Offset: int32(req.Skip),
	})
	if err != nil {
		err = shared.WrapWithAppError(err, "failed to list templates", "")
		log.Error(err.Error())
		return nil, err
	}
	return &ListGeneratedImagesResult{
		Data: result,
	}, nil
}
