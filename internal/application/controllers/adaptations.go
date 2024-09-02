package controllers

import (
	"algvisual/internal/application/usecases/adaptations"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/cognito"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewAdaptationController(
	cfg config.AppConfig,
	cog *cognito.Cognito,
	getActiveAdaptation *adaptations.GetActiveAdaptationBatchUseCase,
	startAdaptation *adaptations.StartAdaptationUseCase,
) AdaptationController {
	return AdaptationController{
		cfg:                 cfg,
		cog:                 cog,
		getActiveAdaptation: getActiveAdaptation,
		startAdaptation:     startAdaptation,
	}
}

type AdaptationController struct {
	cfg                 config.AppConfig
	cog                 *cognito.Cognito
	getActiveAdaptation *adaptations.GetActiveAdaptationBatchUseCase
	startAdaptation     *adaptations.StartAdaptationUseCase
}

func (a AdaptationController) Load(e *echo.Echo) error {
	e.GET(
		"/api/v1/adaptation",
		a.GetActiveAdaptation(),
		middlewares.NewAuthMiddleware(a.cog, a.cfg),
	)
	e.POST(
		"/api/v1/adaptation/start",
		a.StartAdaptation(),
		middlewares.NewAuthMiddleware(a.cog, a.cfg),
	)
	// e.POST(
	// 	"/api/v1/adaptation/stop",
	// 	a.StopAdaptation(),
	// 	middlewares.NewAuthMiddleware(a.cog, a.cfg),
	// )
	// e.GET(
	// 	"/api/v1/adaptation/batch/:batch_id/result",
	// 	a.ListAdaptationResults(),
	// 	middlewares.NewAuthMiddleware(a.cog, a.cfg),
	// )
	// e.GET(
	// 	"/api/v1/adaptation/formats",
	// 	a.ListAdaptationTemplates(),
	// 	middlewares.NewAuthMiddleware(a.cog, a.cfg),
	// )
	return nil
}

func (a AdaptationController) GetActiveAdaptation() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req adaptations.GetActiveAdaptationBatchInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		session := c.Get("session").(*entities.UserSession)
		req.Session = *session
		out, err := a.getActiveAdaptation.Execute(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (a AdaptationController) StartAdaptation() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req adaptations.StartAdaptationInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		session := c.Get("session").(*entities.UserSession)
		req.Session = *session
		out, err := a.startAdaptation.Execute(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
