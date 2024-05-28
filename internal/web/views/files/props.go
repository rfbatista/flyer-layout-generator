package files

import (
	"algvisual/internal/database"
	"algvisual/internal/designprocessor"
	"context"

	"go.uber.org/zap"
)

func Props(ctx context.Context, queries *database.Queries, log *zap.Logger) (*PageProps, error) {
	out, err := designprocessor.ListPhotoshopFilesUseCase(ctx, designprocessor.ListPhotoshopFilesRequest{Limit: 10, Skip: 0}, queries, log)
	if err != nil {
		return nil, err
	}
	return &PageProps{
		files: out.Data,
	}, nil
}
