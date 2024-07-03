package generate

import (
	"algvisual/database"
	"algvisual/internal/designs"
	"algvisual/internal/infra"
	"algvisual/internal/layoutgenerator"
	"algvisual/internal/shared"
	"algvisual/web/components/notification"
	"algvisual/web/render"
	"algvisual/web/views/generate/components/editor"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/rfbatista/apitools"
	"go.uber.org/zap"
)

func NewPage(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
	bundler *infra.Bundler,
) apitools.Handler {
	static, err := bundler.AddPage(infra.BundlerPageParams{
		EntryPoints: []string{
			fmt.Sprintf("%s/web/views/generate/index.js", infra.FindProjectRoot()),
		},
		Name: "editor/newdoc",
	})
	if err != nil {
		panic(shared.WrapWithAppError(err, "failed to build editor new doc page", ""))
	}
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath("/editor/:design_id/:layout_id")
	h.SetHandle(func(c echo.Context) error {
		var req request
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		props, err := Props(c.Request().Context(), queries, log, req)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		return render.Render(c, http.StatusOK, Page(props, static.CSSName, static.JSName))
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
	h.SetPath("/request/batch")
	h.SetHandle(func(c echo.Context) error {
		var req layoutgenerator.CreateLayoutRequestInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := layoutgenerator.CreateLayoutRequestUseCase(
			c.Request().Context(),
			queries,
			db,
			req,
		)
		if err != nil {
			shared.ErrorNotification(c, err.Error())
			return c.NoContent(http.StatusBadRequest)
		}
		shared.SuccessNotification(c, "sucesso")
		return c.JSON(http.StatusOK, out)
	})
	return h
}

func CreateImage(
	queries *database.Queries,
	conn *pgxpool.Pool,
	log *zap.Logger,
	db *pgxpool.Pool,
	config *infra.AppConfig,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath("/editor/create/image")
	h.SetHandle(func(c echo.Context) error {
		var req GenerateImage
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		in := layoutgenerator.GenerateImage{
			PhotoshopID: req.PhotoshopID,
			TemplateID:  req.TemplateID[0],
			LayoutID:    req.LayoutID,
			SlotsX:      req.SlotsX,
			SlotsY:      req.SlotsY,
			Priorities:  req.Priorities,
			Padding:     10,
			ShowGrid:    true,
		}
		out, err := layoutgenerator.GenerateImageUseCase(
			c.Request().Context(),
			in,
			queries,
			db,
			*config,
			log,
		)
		if err != nil {
			shared.ErrorNotification(c, err.Error())
			return c.NoContent(http.StatusBadRequest)
		}
		shared.SuccessNotification(c, "sucesso")
		j, err := json.Marshal(out.Layout)
		if err != nil {
			shared.ErrorNotification(c, err.Error())
			return c.NoContent(http.StatusBadRequest)
		}
		return render.Render(
			c,
			http.StatusOK,
			editor.Editor(editor.EditorProps{
				Layout:     *out.Layout,
				Layoutjson: string(j),
			}),
		)
	})
	return h
}

func CreateComponent(db *database.Queries, tx *pgxpool.Pool, log *zap.Logger) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath("/editor/design/:design_id/layout/:layout_id/component")
	h.SetHandle(func(c echo.Context) error {
		var req designs.CreateComponentRequest
		err := c.Bind(&req)
		if err != nil {
			return shared.RenderComponent(
				shared.WithComponent(
					notification.FailureMessage(err.Error()), c,
				),
			)
		}
		_, err = designs.CreateComponentUseCase(c.Request().Context(), req, db, tx, log)
		if err != nil {
			return err
		}
		c.Response().
			Header().
			Set("HX-Redirect", shared.PageDefineComponents.Replace([]string{strconv.Itoa(int(req.DesignID))}))
		return c.NoContent(http.StatusOK)
	})
	return h
}
