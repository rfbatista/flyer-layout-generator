package files

import (
	"algvisual/internal/database"
	"algvisual/internal/designprocessor"
	"algvisual/internal/infra"
	"algvisual/internal/shared"
	"algvisual/web/components/notification"
	"algvisual/web/render"
	"fmt"
	"net/http"

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
			fmt.Sprintf("%s/web/views/files/index.js", infra.FindProjectRoot()),
		},
		Name: "files/editor",
	})
	if err != nil {
		panic(shared.WrapWithAppError(err, "failed to build editor new doc page", ""))
	}
	h := apitools.NewHandler()
	h.SetMethod(apitools.GET)
	h.SetPath("/")
	h.SetHandle(func(c echo.Context) error {
		props, err := Props(c.Request().Context(), queries, log)
		if err != nil {
			log.Error("failed to render home page", zap.Error(err))
			return err
		}
		return render.Render(c, http.StatusOK, Page(props, static.CSSName, static.JSName))
	})
	return h
}

func NewUploadDesignAPI(
	db *database.Queries,
	proc *infra.PhotoshopProcessor,
	storage infra.FileStorage,
	log *zap.Logger,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath(shared.PageUploadDesignFile.String())
	h.SetHandle(func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return shared.RenderComponent(
				shared.WithComponent(
					notification.FailureMessage(err.Error()), c,
				),
			)
		}
		src, err := file.Open()
		if err != nil {
			shared.ErrorNotification(c, err.Error())
			return c.NoContent(http.StatusBadRequest)
		}
		defer src.Close()
		req := designprocessor.UploadDesignFileUseCaseRequest{
			Filename: c.FormValue("filename"),
			File:     src,
		}
		out, err := designprocessor.UploadDesignFileUseCase(
			c.Request().Context(),
			db,
			req,
			storage.Upload,
			log,
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

func NewProcessDesignFile(
	db *database.Queries,
	proc *infra.PhotoshopProcessor,
	storage infra.FileStorage,
	log *zap.Logger,
	pool *pgxpool.Pool,
) apitools.Handler {
	h := apitools.NewHandler()
	h.SetMethod(apitools.POST)
	h.SetPath("/upload/design/:design_id/process")
	h.SetHandle(func(c echo.Context) error {
		var req designprocessor.ProcessDesignFileRequestv2
		err := c.Bind(&req)
		if err != nil {
			shared.ErrorNotification(c, err.Error())
			return c.NoContent(http.StatusBadRequest)
		}
		out, err := designprocessor.ProcessDesignFileUseCasev2(
			c.Request().Context(),
			req,
			proc,
			log,
			db,
			pool,
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
