package controllers

import (
	"algvisual/internal/application/usecases/adaptations"
	"algvisual/internal/application/usecases/replications"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/cognito"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewJobsController(
	cfg config.AppConfig,
	cog *cognito.Cognito,
	getActiveAdaptation *adaptations.GetActiveAdaptationBatchUseCase,
	startAdaptation *adaptations.StartAdaptationUseCase,
	getActiveReplication *replications.GetActiveReplicationUseCase,
	startReplication *replications.StartReplicationUseCase,
) JobsController {
	return JobsController{
		cfg:                  cfg,
		cog:                  cog,
		getActiveAdaptation:  getActiveAdaptation,
		startAdaptation:      startAdaptation,
		getActiveReplication: getActiveReplication,
		startReplication:     startReplication,
	}
}

type JobsController struct {
	cfg                  config.AppConfig
	cog                  *cognito.Cognito
	getActiveAdaptation  *adaptations.GetActiveAdaptationBatchUseCase
	startAdaptation      *adaptations.StartAdaptationUseCase
	getActiveReplication *replications.GetActiveReplicationUseCase
	startReplication     *replications.StartReplicationUseCase
}

func (a JobsController) Load(e *echo.Echo) error {
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
	e.GET(
		"/api/v1/replication",
		a.GetActiveReplication(),
		middlewares.NewAuthMiddleware(a.cog, a.cfg),
	)
	e.POST(
		"/api/v1/replication/start",
		a.StartReplication(),
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

func (a JobsController) GetActiveAdaptation() echo.HandlerFunc {
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

func (a JobsController) StartAdaptation() echo.HandlerFunc {
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

func (a JobsController) StartReplication() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req replications.StartReplicationInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		session := c.Get("session").(*entities.UserSession)
		req.Session = *session
		out, err := a.startReplication.Execute(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}

func (a JobsController) GetActiveReplication() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req replications.GetActiveReplicationInput
		err := c.Bind(&req)
		if err != nil {
			return err
		}
		session := c.Get("session").(*entities.UserSession)
		req.Session = *session
		out, err := a.getActiveReplication.Execute(c.Request().Context(), req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, out)
	}
}
