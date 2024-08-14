package designprocessor

import (
	"algvisual/database"
	"algvisual/internal/designassets"
	"algvisual/internal/designprocessor/usecase"
	"algvisual/internal/infra"
	"algvisual/internal/infra/cognito"
	"algvisual/internal/infra/config"
	"algvisual/internal/infra/middlewares"
	"algvisual/internal/layoutgenerator"
	"algvisual/internal/shared"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewDesignController(
	db *database.Queries,
	proc *infra.PhotoshopProcessor,
	storage infra.FileStorage,
	log *zap.Logger,
	pool *pgxpool.Pool,
	processorFile *infra.PhotoshopProcessor,
	ds *designassets.DesignAssetService,
	cfg config.AppConfig,
	cog *cognito.Cognito,
	dpr *DesignProcessorService,
) DesignController {
	return DesignController{
		db:            db,
		ds:            ds,
		proc:          proc,
		storage:       storage,
		log:           log,
		cfg:           cfg,
		pool:          pool,
		processorFile: processorFile,
		dpr:           dpr,
		cog:           cog,
	}
}

type DesignController struct {
	db            *database.Queries
	ds            *designassets.DesignAssetService
	proc          *infra.PhotoshopProcessor
	storage       infra.FileStorage
	log           *zap.Logger
	cfg           config.AppConfig
	pool          *pgxpool.Pool
	processorFile *infra.PhotoshopProcessor
	dpr           *DesignProcessorService
	cog           *cognito.Cognito
}

func (s DesignController) Load(e *echo.Echo) error {
	e.POST(
		"/api/v1/design",
		s.UploadDesign(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.POST(
		"/api/v1/design/:design_id/process",
		s.ProcessDesginFileByID(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.GET(
		"/api/v1/project/design/:design_id/last_request",
		s.GetLastRequestJob(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.GET(
		"/api/v1/project/design/:design_id/layout/:request_id",
		s.GetLayoutByID(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	return nil
}

func (s DesignController) UploadDesign() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			shared.ErrorNotification(c, err.Error())
			return c.NoContent(http.StatusBadRequest)
		}
		defer src.Close()
		var req usecase.UploadDesignFileUseCaseRequest
		err = c.Bind(&req)
		if err != nil {
			return err
		}
		req.File = src
		out, err := s.dpr.UploadDesignFile(
			c,
			req,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s DesignController) ProcessDesginFileByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.ProcessDesignFileRequestv2
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := s.dpr.ProcessDesignFileV2(
			c,
			req,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s DesignController) GetLastRequestJob() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req layoutgenerator.GetLastLayoutRequestInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := layoutgenerator.GetLastLayoutRequestUseCase(
			c.Request().Context(),
			req,
			s.db,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s DesignController) GetLayoutByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req layoutgenerator.GetLayoutByIDRequest
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := layoutgenerator.GetLayoutByIDUseCase(
			c.Request().Context(),
			s.db,
			req,
			s.ds,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
