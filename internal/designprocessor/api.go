package designprocessor

import (
	"algvisual/database"
	"algvisual/internal/designassets"
	"algvisual/internal/designprocessor/usecase"
	"algvisual/internal/infra"
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
) DesignController {
	return DesignController{
		db:            db,
		ds:            ds,
		proc:          proc,
		storage:       storage,
		log:           log,
		pool:          pool,
		processorFile: processorFile,
	}
}

type DesignController struct {
	db            *database.Queries
	proc          *infra.PhotoshopProcessor
	storage       infra.FileStorage
	log           *zap.Logger
	pool          *pgxpool.Pool
	processorFile *infra.PhotoshopProcessor
	ds            *designassets.DesignAssetService
	dpr           *DesignProcessorService
}

func (s DesignController) Load(e *echo.Echo) error {
	e.POST("/api/v1/design", s.UploadDesign())
	e.POST("/api/v1/design/:design_id/process", s.ProcessDesginFileByID())
	e.GET("/api/v1/project/design/:design_id/last_request", s.GetLastRequestJob())
	e.GET("/api/v1/project/design/:design_id/layout/:request_id", s.GetLayoutByID())
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
