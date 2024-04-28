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

func NewCreateTemplateAPI(db *database.Queries, conn *pgx.Conn, log *zap.Logger) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.EndpointCreateTemplate.String())
	h.SetHandle(func(c echo.Context) error {
		var req usecases.CreateTemplateUseCaseRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := usecases.CreateTemplateUseCase(c.Request().Context(), conn, db, req, log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}

func NewListTemplatesAPI(
	queries *database.Queries,
	conn *pgx.Conn,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.EndpointListTemplate.String())
	h.SetHandle(func(c echo.Context) error {
		var req usecases.ListTemplatesUseCaseRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := usecases.ListTemplatesUseCase(c.Request().Context(), req, queries, log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	})
	return h
}
