package api

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/shared"
	"algvisual/internal/usecases"
)

func NewRemoveComponentAPI(
	db *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.EndpointRemoveComponentElements.String())
	h.SetHandle(func(c echo.Context) error {
		var req usecases.RemoveComponentUseCaseRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := usecases.RemoveComponentUseCase(c.Request().Context(), db, req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}

func NewSetPhotoshopBackgroundAPI(
	db *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.EndpointSetPhotoshopBackground.String())
	h.SetHandle(func(c echo.Context) error {
		var req usecases.SetBackgroundUseCaseRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := usecases.SetBackgroundUseCase(c.Request().Context(), db, conn, req, log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}

func NewListComponentsByFileIDAPI(
	db *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.ListComponentByFileIDEndpoint.String())
	h.SetHandle(func(c echo.Context) error {
		var req usecases.ListComponentsByFileIdRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := usecases.ListComponentsByFileIdUseCase(c.Request().Context(), req, db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}
