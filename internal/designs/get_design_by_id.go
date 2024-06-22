package designs

import (
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"

	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/shared"
)

type GetDesignByIdRequest struct {
	ID int32 `params:"photoshop_id" json:"id,omitempty"`
}

type GetPhotoshopByIdResult struct {
	Status string
	Data   entities.DesignFile
}

func GetPhotoshopByIdUseCase(
	ctx context.Context,
	req GetDesignByIdRequest,
	queries *database.Queries,
	log *zap.Logger,
) (*GetPhotoshopByIdResult, error) {
	photoshop, err := queries.Getdesign(ctx, req.ID)
	if err != nil {
		err = shared.WrapWithAppError(err, "failed to create templates distortion", "")
		log.Error(err.Error())
		return nil, err
	}
	return &GetPhotoshopByIdResult{
		Status: "success",
		Data:   mapper.TodesignEntitie(photoshop),
	}, nil
}
