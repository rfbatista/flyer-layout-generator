package adaptations

import (
	"algvisual/internal/adaptations/usecase"
	"algvisual/internal/entities"
	"algvisual/internal/infra/cognito"
	"algvisual/internal/infra/config"
	"algvisual/internal/infra/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewAdaptationController(
	service *AdaptationService,
	cfg config.AppConfig,
	cog *cognito.Cognito,
) AdaptationController {
	return AdaptationController{
		service: service,
		cfg:     cfg,
		cog:     cog,
	}
}

type AdaptationController struct {
	service *AdaptationService
	cfg     config.AppConfig
	cog     *cognito.Cognito
}

func (a AdaptationController) Load(e *echo.Echo) error {
	e.POST(
		"/api/v1/adaptation/start",
		a.StartAdaptationBatch(),
		middlewares.NewAuthMiddleware(a.cog, a.cfg),
	)
	e.POST(
		"/api/v1/adaptation/stop",
		a.StopAdaptation(),
		middlewares.NewAuthMiddleware(a.cog, a.cfg),
	)
	e.GET(
		"/api/v1/adaptation/batch",
		a.GetActiveAdaptation(),
		middlewares.NewAuthMiddleware(a.cog, a.cfg),
	)
	e.GET(
		"/api/v1/adaptation/batch/:batch_id/result",
		a.ListAdaptationResults(),
		middlewares.NewAuthMiddleware(a.cog, a.cfg),
	)
	e.GET(
		"/api/v1/adaptation/formats",
		a.ListAdaptationTemplates(),
		middlewares.NewAuthMiddleware(a.cog, a.cfg),
	)
	return nil
}

func (a AdaptationController) StartAdaptationBatch() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.StartAdaptationInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		session := c.Get("session").(*entities.UserSession)
		req.Session = *session
		out, err := a.service.StartAdaptation(c, req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (a AdaptationController) GetActiveAdaptation() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.GetActiveAdaptationBatchInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		session := c.Get("session").(*entities.UserSession)
		req.Session = *session
		out, err := a.service.GetActiveAdaptation(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (a AdaptationController) ListAdaptationTemplates() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.ListAdaptationTemplatesInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := a.service.ListAdatptationTemplates(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (a AdaptationController) ListAdaptationResults() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.ListAdaptationResultsInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		out, err := a.service.ListAdaptationResults(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (a AdaptationController) StopAdaptation() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req usecase.StopAdaptationBatchInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		session := c.Get("session").(*entities.UserSession)
		req.Session = *session
		out, err := a.service.StopAdaptation(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
