package home

import (
	"algvisual/internal/database"
	"algvisual/internal/layoutgenerator"
	"algvisual/internal/shared"
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"
)

func NewPageHome(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.PageHome.String())
	h.SetHandle(func(c echo.Context) error {
		props, err := Props(c.Request().Context(), queries, log)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		component := HomePage(props)
		w := c.Response().Writer
		err = component.Render(
			context.WithValue(c.Request().Context(), "page", shared.PageHome.String()),
			w,
		)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		return nil
	})
	return h
}

func CreateRequest(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
	db *pgxpool.Pool,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.PageHomeCreateRequest.String())
	h.SetHandle(func(c echo.Context) error {
		var req layoutgenerator.CreateLayoutRequestInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := layoutgenerator.CreateLayoutRequestUseCase(c.Request().Context(), queries, db, req)
		if err != nil {
			shared.Error(c, err.Error())
			return c.NoContent(http.StatusBadRequest)
		}
		shared.Success(c, "sucesso")
		return c.JSON(http.StatusOK, out)
	})
	return h
}
