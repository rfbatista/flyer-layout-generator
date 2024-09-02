package designs

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories/mapper"
	"algvisual/internal/shared"

	"github.com/labstack/echo/v4"
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
	c echo.Context,
	req GetDesignByIdRequest,
	queries *database.Queries,
	log *zap.Logger,
) (*GetDesignByIdResult, error) {
	ctx := c.Request().Context()
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
