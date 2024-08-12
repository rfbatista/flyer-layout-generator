package layoutgenerator

import (
	"algvisual/database"
	"algvisual/internal/infra/config"
	"algvisual/internal/layoutgenerator/usecase"
	"algvisual/internal/renderer"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewLayoutController(
	db *database.Queries,
	lservice LayoutGeneratorService,
	cfg *config.AppConfig,
	log *zap.Logger,
	pool *pgxpool.Pool,
	render renderer.RendererService,
) LayoutController {
	return LayoutController{db: db, layoutService: lservice, cfg: cfg, log: log, pool: pool, render: render}
}

type LayoutController struct {
	db            *database.Queries
	layoutService LayoutGeneratorService
	render        renderer.RendererService
	pool          *pgxpool.Pool
	cfg           *config.AppConfig
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
	e.DELETE("/api/v1/batch/:batch_id/layout/:layout_id", s.DeleteBatchLayout())
	return nil
}

func (s LayoutController) GetLayoutByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req GetLayoutByIDRequest
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := GetLayoutByIDUseCase(c.Request().Context(), s.db, req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s LayoutController) GenerateLayout() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req GenerateImage
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := GenerateImageUseCase(
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
		var req GenerateImageV2Input
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
		var req CreateLayoutRequestInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := CreateLayoutRequestUseCase(
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
		var req UpdateLayoutElementPositionInput
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
		var req UpdateLayoutElementSizeInput
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

func (s LayoutController) DeleteBatchLayout() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.DeleteLayoutByIdInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := s.layoutService.DeleteLayout(
			c.Request().Context(),
			req,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
