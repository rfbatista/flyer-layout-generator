package usecases

import (
	"context"

	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/shared"
)

type GetPhotoshopByIdRequest struct {
	ID int32 `params:"photoshop_id" json:"id,omitempty"`
}

type GetPhotoshopByIdResult struct {
	Status string
	Data   entities.DesignFile
}

func GetPhotoshopByIdUseCase(
	ctx context.Context,
	req GetPhotoshopByIdRequest,
	queries *database.Queries,
	log *zap.Logger,
) (*GetPhotoshopByIdResult, error) {
	photoshop, err := queries.GetPhotoshop(ctx, req.ID)
	if err != nil {
		err = shared.WrapWithAppError(err, "failed to create template distortion", "")
		log.Error(err.Error())
		return nil, err
	}
	return &GetPhotoshopByIdResult{
		Status: "success",
		Data:   database.TodesignEntitie(photoshop),
	}, nil
}
