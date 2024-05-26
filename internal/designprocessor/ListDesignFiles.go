package designprocessor

import (
	"context"

	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/shared"
)

type ListPhotoshopFilesRequest struct {
	Limit int `query:"limit" json:"limit,omitempty"`
	Skip  int `query:"skip"  json:"skip,omitempty"`
}

type ListPhotoshopFilesResult struct {
	Status string            `json:"status,omitempty"`
	Data   []database.Design `json:"data,omitempty"`
}

func ListPhotoshopFilesUseCase(
	ctx context.Context,
	req ListPhotoshopFilesRequest,
	queries *database.Queries,
	log *zap.Logger,
) (*ListPhotoshopFilesResult, error) {
	limit := req.Limit
	if limit == 0 {
		limit = 10
	}
	files, err := queries.Listdesign(ctx, database.ListdesignParams{
		Offset: int32(req.Skip),
		Limit:  int32(limit),
	})
	if err != nil {
		log.Error("failed to list photoshop files", zap.Error(err))
		return nil, shared.WrapWithAppError(err, "Falha ao listar aquivos do Photoshop", "")
	}
	return &ListPhotoshopFilesResult{Data: files, Status: "success"}, nil
}
