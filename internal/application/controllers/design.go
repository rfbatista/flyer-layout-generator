package controllers

import (
	"algvisual/internal/application/usecases/designs"
	"algvisual/internal/infrastructure/cognito"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/middlewares"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewDesignController(
	db *database.Queries,
	pool *pgxpool.Pool,
	log *zap.Logger,
	cog *cognito.Cognito,
	cfg config.AppConfig,
) (*DesignController, error) {
	return &DesignController{
		db:   db,
		pool: pool,
		log:  log,
		cog:  cog,
		cfg:  cfg,
	}, nil
}

type DesignController struct {
	db   *database.Queries
	pool *pgxpool.Pool
	log  *zap.Logger
	cog  *cognito.Cognito
	cfg  config.AppConfig
}

func (d DesignController) Load(e *echo.Echo) error {
	e.GET(
		"/api/v1/design/:id/file",
		d.GetDesignFile(),
	)
	e.POST(
		"/api/v1/photoshop/:photoshop_id/components/remove",
		d.RemoveComponentAPI(),
		middlewares.NewAuthMiddleware(d.cog, d.cfg),
	)
	e.POST(
		"/api/v1/photoshop/:photoshop_id/background",
		d.SetPhotoshopBackgroundAPI(),
		middlewares.NewAuthMiddleware(d.cog, d.cfg),
	)
	e.GET(
		"/api/v1/file/:photoshop_id/components",
		d.ListComponentsByFileIDAPI(),
		middlewares.NewAuthMiddleware(d.cog, d.cfg),
	)
	e.POST(
		"/api/v1/editor/design/:design_id/layout/:layout_id/component",
		d.CreateComponent(),
		middlewares.NewAuthMiddleware(d.cog, d.cfg),
	)
	e.GET(
		"/api/v1/design/:design_id",
		d.GetDesignByID(),
		middlewares.NewAuthMiddleware(d.cog, d.cfg),
	)
	e.GET(
		"/api/v1/designs/project/:project_id",
		d.ListDesignsByProjectID(),
		middlewares.NewAuthMiddleware(d.cog, d.cfg),
	)
	e.POST(
		"/editor/design/:design_id/layout/:layout_id/component",
		d.CreateComponent(),
		middlewares.NewAuthMiddleware(d.cog, d.cfg),
	)
	return nil
}

func (d DesignController) CreateComponent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req designs.CreateComponentRequest
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := designs.CreateComponentUseCase(c, req, d.db, d.pool, d.log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (d DesignController) ListComponentsByFileIDAPI() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req designs.ListComponentsByFileIdRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := designs.ListComponentsByFileIdUseCase(c, req, d.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (d DesignController) SetPhotoshopBackgroundAPI() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req designs.SetBackgroundUseCaseRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := designs.SetBackgroundUseCase(c, d.db, d.pool, req, d.log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (d DesignController) RemoveComponentAPI() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req designs.RemoveComponentUseCaseRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := designs.RemoveComponentUseCase(c, d.db, req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (s DesignController) GetDesignByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req designs.GetDesignByIdRequest
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := designs.GetDesignByIdUseCase(c, req, s.db, s.log)
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
		out, err := designs.ListDesignByProjectIdUseCase(c, req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s DesignController) GetDesignFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			return err
		}
		design, err := s.db.Getdesign(c.Request().Context(), int32(idInt))
		if err != nil {
			return err
		}
		return c.File(design.FileUrl.String)
	}
}
