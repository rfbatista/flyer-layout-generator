package usecase

import (
	"algvisual/database"
	"algvisual/internal/shared"

	"github.com/labstack/echo/v4"
)

type ListComponentsByFileIdRequest struct {
	DesignID int32 `param:"photoshop_id"`
}

type ListComponentsByFileIdResult struct {
	Status string                     `json:"status,omitempty"`
	Data   []database.LayoutComponent `json:"data,omitempty"`
}

func ListComponentsByFileIdUseCase(
	c echo.Context,
	req ListComponentsByFileIdRequest,
	queries *database.Queries,
) (*ListComponentsByFileIdResult, error) {
	ctx := c.Request().Context()
	res, err := queries.GetComponentsByDesignID(ctx, req.DesignID)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao listar componentes por arquivo", "")
		return nil, err
	}
	return &ListComponentsByFileIdResult{
		Status: "success",
		Data:   res,
	}, nil
}
