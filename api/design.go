package api

import (
	"algvisual/internal/database"
	"algvisual/internal/designprocessor"
	"algvisual/internal/designs"
	"algvisual/internal/infra"
	"algvisual/internal/shared"
	"algvisual/web/components/notification"
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
) DesignController {
	return DesignController{
		db:            db,
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
}

func (s DesignController) Load(e *echo.Echo) error {
	e.POST("/api/v1/design", s.UploadDesign())
	e.GET("/api/v1/designs/project/:project_id", s.ListDesignsByProjectID())
	e.POST("/api/v1/design/:design_id/process", s.ProcessDesginFileByID())
	e.GET("/api/v1/design/:design_id", s.GetDesignByID())
	e.POST("/editor/design/:design_id/layout/:layout_id/component", s.CreateComponent())
	return nil
}

func (s DesignController) UploadDesign() echo.HandlerFunc {
	return func(c echo.Context) error {
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
		var req designprocessor.UploadDesignFileUseCaseRequest
		err = c.Bind(&req)
		if err != nil {
			return err
		}
		req.File = src
		out, err := designprocessor.UploadDesignFileUseCase(
			c.Request().Context(),
			s.db,
			req,
			s.storage.Upload,
			s.log,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s DesignController) GetDesignByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req designs.GetDesignByIdRequest
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := designs.GetDesignByIdUseCase(c.Request().Context(), req, s.db, s.log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s DesignController) ListDesignsByProjectID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req designs.ListDesignByProjectIdInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := designs.ListDesignByProjectIdUseCase(c.Request().Context(), req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s DesignController) ProcessDesginFileByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req designprocessor.ProcessDesignFileRequestv2
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := designprocessor.ProcessDesignFileUseCasev2(
			c.Request().Context(),
			req,
			s.processorFile,
			s.log,
			s.db,
			s.pool,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s DesignController) CreateComponent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req designs.CreateComponentRequest
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := designs.CreateComponentUseCase(
			c.Request().Context(),
			req,
			s.db,
			s.pool,
			s.log,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
