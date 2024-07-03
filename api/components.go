package api

import (
	"algvisual/internal/designs"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"

	"algvisual/database"
	"algvisual/internal/shared"
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
		var req designs.RemoveComponentUseCaseRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := designs.RemoveComponentUseCase(c.Request().Context(), db, req)
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
		var req designs.SetBackgroundUseCaseRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := designs.SetBackgroundUseCase(c.Request().Context(), db, conn, req, log)
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
		var req designs.ListComponentsByFileIdRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := designs.ListComponentsByFileIdUseCase(c.Request().Context(), req, db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}

func CreateComponent(db *database.Queries, tx *pgxpool.Pool, log *zap.Logger) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath("/api/v1/editor/design/:design_id/layout/:layout_id/component")
	h.SetHandle(func(c echo.Context) error {
		var req designs.CreateComponentRequest
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := designs.CreateComponentUseCase(c.Request().Context(), req, db, tx, log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	})
	return h
}
