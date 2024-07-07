package api

import (
	"algvisual/database"
	"algvisual/internal/infra"
	"algvisual/internal/layoutgenerator"
	"algvisual/internal/renderer"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewLayoutController(
	db *database.Queries,
	lservice layoutgenerator.LayoutGeneratorService,
	cfg *infra.AppConfig,
	log *zap.Logger,
	pool *pgxpool.Pool,
	render renderer.RendererService,
) LayoutController {
	return LayoutController{db: db, layoutService: lservice, cfg: cfg, log: log, pool: pool, render: render}
}

type LayoutController struct {
	db            *database.Queries
	layoutService layoutgenerator.LayoutGeneratorService
	render        renderer.RendererService
	pool          *pgxpool.Pool
	cfg           *infra.AppConfig
	log           *zap.Logger
}

func (s LayoutController) Load(e *echo.Echo) error {
	e.GET("/api/v1/layout/:layout_id", s.GetLayoutByID())
	e.POST(
		"/api/v1/design/:design_id/layout/:layout_id/template/:template_id/generate",
		s.GenerateLayout(),
	)
	e.POST("/api/v2/layout/:layout_id/template/:template_id/generate", s.GenerateLayoutv2())
	e.POST("/api/v1/project/design/:design_id/layout/:layout_id/generate", s.CreateLayoutRequest())
	e.PATCH("/api/v1/layout/:layout_id/element/:element_id/position", s.UpdateLayoutElementPosition())
	e.PATCH("/api/v1/layout/:layout_id/element/:element_id/size", s.UpdateLayoutElementSize())
	return nil
}

func (s LayoutController) GetLayoutByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req layoutgenerator.GetLayoutByIDRequest
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := layoutgenerator.GetLayoutByIDUseCase(c.Request().Context(), s.db, req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s LayoutController) GenerateLayout() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req layoutgenerator.GenerateImage
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := layoutgenerator.GenerateImageUseCase(
			c.Request().Context(),
			req,
			s.db,
			s.pool,
			*s.cfg,
			s.log,
			s.render,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s LayoutController) GenerateLayoutv2() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req layoutgenerator.GenerateImageV2Input
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := s.layoutService.GenerateNewLayout(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s LayoutController) CreateLayoutRequest() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req layoutgenerator.CreateLayoutRequestInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := layoutgenerator.CreateLayoutRequestUseCase(
			c.Request().Context(),
			s.db,
			s.pool,
			req,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s LayoutController) UpdateLayoutElementPosition() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req layoutgenerator.UpdateLayoutElementPositionInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := s.layoutService.UpdateElementPosition(
			c.Request().Context(),
			req,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s LayoutController) UpdateLayoutElementSize() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req layoutgenerator.UpdateLayoutElementSizeInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := s.layoutService.UpdateElementSize(
			c.Request().Context(),
			req,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
