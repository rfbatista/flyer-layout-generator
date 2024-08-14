package designs

import (
	"algvisual/database"
	"algvisual/internal/designs/usecase"
	"algvisual/internal/infra/cognito"
	"algvisual/internal/infra/config"
	"algvisual/internal/infra/middlewares"
	"net/http"

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
		var req usecase.CreateComponentRequest
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := usecase.CreateComponentUseCase(c, req, d.db, d.pool, d.log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (d DesignController) ListComponentsByFileIDAPI() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.ListComponentsByFileIdRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := usecase.ListComponentsByFileIdUseCase(c, req, d.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (d DesignController) SetPhotoshopBackgroundAPI() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.SetBackgroundUseCaseRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := usecase.SetBackgroundUseCase(c, d.db, d.pool, req, d.log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (d DesignController) RemoveComponentAPI() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.RemoveComponentUseCaseRequest
		err := c.Bind(&req)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		result, err := usecase.RemoveComponentUseCase(c, d.db, req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (s DesignController) GetDesignByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.GetDesignByIdRequest
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := usecase.GetDesignByIdUseCase(c, req, s.db, s.log)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s DesignController) ListDesignsByProjectID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.ListDesignByProjectIdInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := usecase.ListDesignByProjectIdUseCase(c, req, s.db)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
