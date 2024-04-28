package api

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/shared"
	"algvisual/internal/usecases"
)

func NewListGeneratedImagesAPI(
	queries *database.Queries,
	conn *pgx.Conn,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.EndpointListImagesGenerated.String())
	h.SetHandle(func(c echo.Context) error {
		var req usecases.ListGeneratedImagesRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := usecases.ListGeneratedImagesUseCase(c.Request().Context(), req, queries, log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}
