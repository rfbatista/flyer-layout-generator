package designs

import (
	"algvisual/database"
	"context"

	"go.uber.org/zap"
)

type DesignService struct {
	db  *database.Queries
	log *zap.Logger
}

func (d DesignService) GetDesignByID(
	ctx context.Context,
	in GetDesignByIdRequest,
) (*GetDesignByIdResult, error) {
	return GetDesignByIdUseCase(ctx, in, d.db, d.log)
}
