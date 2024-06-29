package designs

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/layoutgenerator"
	"algvisual/internal/mapper"
	"algvisual/internal/shared"
	"context"

	"go.uber.org/zap"
)

type GetDesignByIdRequest struct {
	ID int32 `params:"design_id" json:"id,omitempty"`
}

type GetDesignByIdResult struct {
	Status string
	Data   entities.DesignFile
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
	if design.LayoutID.Valid {
		out, err := layoutgenerator.GetLayoutByIDUseCase(
			ctx,
			queries,
			layoutgenerator.GetLayoutByIDRequest{
				LayoutID: design.ID,
			},
		)
		if err != nil {
			err = shared.WrapWithAppError(err, "failed to get layoutby id", "")
			log.Error(err.Error())
			return nil, err
		}
		desgnEntities.Layout = out.Layout
	}
	return &GetDesignByIdResult{
		Status: "success",
		Data:   desgnEntities,
	}, nil
}
