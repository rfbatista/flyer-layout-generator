package api

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/infra"
	"algvisual/internal/shared"
	"algvisual/internal/usecases"
)

func NewGenerateDesignAPI(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
	client *infra.ImageGeneratorClient,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.CreateNewDesignEndpoint.String())
	h.SetHandle(func(c echo.Context) error {
		var req usecases.GenerateDesignRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := usecases.GenerateDesignUseCase(c.Request().Context(), req, client, queries)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}
