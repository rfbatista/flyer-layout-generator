package controllers

import (
	"algvisual/internal/application/usecases/designassets"
	"algvisual/internal/application/usecases/designprocessor"
	"algvisual/internal/application/usecases/layoutgenerator"
	"algvisual/internal/infrastructure"
	"algvisual/internal/infrastructure/cognito"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/middlewares"
	"algvisual/internal/infrastructure/storage"
	"algvisual/internal/shared"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewDesignProcessorController(
	db *database.Queries,
	proc *infrastructure.PhotoshopProcessor,
	storage storage.FileStorage,
	log *zap.Logger,
	pool *pgxpool.Pool,
	processorFile *infrastructure.PhotoshopProcessor,
	ds *designassets.DesignAssetService,
	cfg config.AppConfig,
	cog *cognito.Cognito,
	dpr *designprocessor.DesignProcessorService,
) DesignProcessorController {
	return DesignProcessorController{
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

type DesignProcessorController struct {
	db            *database.Queries
	ds            *designassets.DesignAssetService
	proc          *infrastructure.PhotoshopProcessor
	storage       storage.FileStorage
	log           *zap.Logger
	cfg           config.AppConfig
	pool          *pgxpool.Pool
	processorFile *infrastructure.PhotoshopProcessor
	dpr           *designprocessor.DesignProcessorService
	cog           *cognito.Cognito
}

func (s DesignProcessorController) Load(e *echo.Echo) error {
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

func (s DesignProcessorController) UploadDesign() echo.HandlerFunc {
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
		var req designprocessor.UploadDesignFileUseCaseRequest
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

func (s DesignProcessorController) ProcessDesginFileByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req designprocessor.ProcessDesignFileRequestv2
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

func (s DesignProcessorController) GetLastRequestJob() echo.HandlerFunc {
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

func (s DesignProcessorController) GetLayoutByID() echo.HandlerFunc {
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
