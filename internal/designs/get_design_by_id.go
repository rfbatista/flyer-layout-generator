package designs

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"algvisual/internal/shared"
	"context"

	"go.uber.org/zap"
)

type GetDesignByIdRequest struct {
	ID int32 `param:"design_id" json:"id,omitempty"`
}

type GetDesignByIdResult struct {
	Status string              `json:"status,omitempty"`
	Data   entities.DesignFile `json:"data,omitempty"`
}

func GetDesignByIdUseCase(
	ctx context.Context,
	req GetDesignByIdRequest,
	queries *database.Queries,
	log *zap.Logger,
) (*GetDesignByIdResult, error) {
	design, err := queries.Getdesign(ctx, req.ID)
	if err != nil {
		err = shared.WrapWithAppError(err, "failed to get design by id", "")
		log.Error(err.Error())
		return nil, err
	}
	desgnEntities := mapper.DesignFileToDomain(design)
	return &GetDesignByIdResult{
		Status: "success",
		Data:   desgnEntities,
	}, nil
}
