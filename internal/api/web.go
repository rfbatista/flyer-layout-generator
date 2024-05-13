package api

import (
	"algvisual/web/views/request/requestuploadfile"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/shared"
	"algvisual/web/views/createtemplatepage"
	"algvisual/web/views/defineelements"
	"algvisual/web/views/home"
	"algvisual/web/views/proccessdesign"
)

func NewPageCreateTemplate(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath("/criar")
	h.SetHandle(func(c echo.Context) error {
		component := createtemplatepage.CreateTemplatePage()
		w := c.Response().Writer
		err := component.Render(c.Request().Context(), w)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		return nil
	})
	return h
}

func NewPageHome(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.PageHome.String())
	h.SetHandle(func(c echo.Context) error {
		component := home.HomePage()
		w := c.Response().Writer
		err := component.Render(context.WithValue(c.Request().Context(), "page", shared.PageHome.String()), w)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		return nil
	})
	return h
}

func NewPageDefineElements(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.PageDefineElements.String())
	h.SetHandle(func(c echo.Context) error {
		component := defineelements.DefineElementsPage()
		w := c.Response().Writer
		err := component.Render(c.Request().Context(), w)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		return nil
	})
	return h
}

func NewPageProccessDesign(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.PageProccessDesing.String())
	h.SetHandle(func(c echo.Context) error {
		ctx := c.Request().Context()
		d, err := queries.Listdesign(ctx, database.ListdesignParams{Offset: 0, Limit: 100})
		component := proccessdesign.Page(d)
		w := c.Response().Writer
		err = component.Render(c.Request().Context(), w)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		return nil
	})
	return h
}

func NewPageRequestUploadFile() apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath(shared.PageRequestUploadFile.String())
	h.SetHandle(func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "page", shared.PageRequestUploadFile.String())
		return shared.RenderComponent(
			requestuploadfile.PageRequestUploadFile(),
			ctx,
		)
	})
	return h
}
