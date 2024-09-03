package controllers

import (
	"algvisual/internal/application/usecases/designassets"
	"algvisual/internal/application/usecases/layoutgenerator"
	"algvisual/internal/application/usecases/renderer"
	"algvisual/internal/infrastructure/cognito"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/middlewares"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewLayoutController(
	db *database.Queries,
	lservice layoutgenerator.LayoutGeneratorService,
	log *zap.Logger,
	pool *pgxpool.Pool,
	render renderer.RendererService,
	das *designassets.DesignAssetService,
	cfg config.AppConfig,
	cog *cognito.Cognito,
	getLayoutByAdaptation *layoutgenerator.GetLayoutByJobsUseCase,
) LayoutController {
	return LayoutController{
		db:             db,
		layoutService:  lservice,
		cfg:            cfg,
		log:            log,
		cog:            cog,
		pool:           pool,
		render:         render,
		das:            das,
		getLayoutByJob: getLayoutByAdaptation,
	}
}

type LayoutController struct {
	db             *database.Queries
	layoutService  layoutgenerator.LayoutGeneratorService
	render         renderer.RendererService
	pool           *pgxpool.Pool
	cfg            config.AppConfig
	cog            *cognito.Cognito
	log            *zap.Logger
	das            *designassets.DesignAssetService
	getLayoutByJob *layoutgenerator.GetLayoutByJobsUseCase
}

func (s LayoutController) Load(e *echo.Echo) error {
	e.GET(
		"/api/v1/layout/:layout_id",
		s.GetLayoutByID(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.GET(
		"/api/v1/layout/adaptation/:adaptation_id",
		s.GetLayoutByJob(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.POST(
		"/api/v2/layout/:layout_id/template/:template_id/generate",
		s.GenerateLayoutv2(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.POST(
		"/api/v1/project/design/:design_id/layout/:layout_id/generate",
		s.CreateLayoutRequest(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.PATCH(
		"/api/v1/layout/:layout_id/element/:element_id/position",
		s.UpdateLayoutElementPosition(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.PATCH(
		"/api/v1/layout/:layout_id/element/:element_id/size",
		s.UpdateLayoutElementSize(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.DELETE(
		"/api/v1/batch/:batch_id/layout/:layout_id",
		s.DeleteBatchLayout(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	e.GET(
		"/api/v1/batch/:batch_id/download",
		s.CreateZipBatch(),
		middlewares.NewAuthMiddleware(s.cog, s.cfg),
	)
	return nil
}

func (s LayoutController) GetLayoutByID() echo.HandlerFunc {
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
			s.das,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

// func (s LayoutController) GenerateLayout() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var req layoutgenerator.GenerateImage
// 		err := c.Bind(&req)
// 		if err != nil {
// 			return err
// 		}
// 		out, err := layoutgenerator.GenerateImageFromLayoutUseCase(
// 			c.Request().Context(),
// 			req,
// 			s.db,
// 			s.pool,
// 			s.cfg,
// 			s.log,
// 			s.render,
// 			s.das,
// 		)
// 		if err != nil {
// 			return err
// 		}
// 		return c.JSON(http.StatusOK, out)
// 	}
// }

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

func (s LayoutController) DeleteBatchLayout() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req layoutgenerator.DeleteLayoutByIdInput
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

func (s LayoutController) CreateZipBatch() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req layoutgenerator.CreateZipForBatchInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := s.layoutService.ZipBatchImages(
			c.Request().Context(),
			req,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (s LayoutController) GetLayoutByJob() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req layoutgenerator.GetLayoutByJobsInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := s.getLayoutByJob.Execute(
			c.Request().Context(),
			req,
		)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
