package files

import (
	"algvisual/internal/database"
	"algvisual/internal/designprocessor"
	"context"

	"go.uber.org/zap"
)

func Props(ctx context.Context, queries *database.Queries, log *zap.Logger) (PageProps, error) {
	var props PageProps
	out, err := designprocessor.ListDesignFiles(
		ctx,
		designprocessor.ListDesignFilesRequest{Limit: 10, Skip: 0},
		queries,
		log,
	)
	if err != nil {
		return props, err
	}
	props.files = out.Data
	return props, nil
}
