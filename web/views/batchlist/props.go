package batchlist

import (
	"algvisual/internal/database"
	"algvisual/internal/layoutgenerator"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Props(
	c echo.Context,
	db *database.Queries,
	log *zap.Logger,
) (PageProps, error) {
	var props PageProps
	var in layoutgenerator.ListLayoytRequestInput
	err := c.Bind(&in)
	if err != nil {
		return props, err
	}
	out, err := layoutgenerator.ListLayoutRequestUseCase(c.Request().Context(), db, in)
	if err != nil {
		return props, err
	}
	props.requests = out.Requests
	return props, nil
}
