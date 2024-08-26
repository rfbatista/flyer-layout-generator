package usecase

import (
	"algvisual/internal/adaptations/errors"
	"algvisual/internal/adaptations/repositories"
	"algvisual/internal/entities"
	"algvisual/internal/infra/sqs"
	"algvisual/internal/shared"

	"github.com/labstack/echo/v4"
)

type StartAdaptationInput struct {
	LayoutID int32 `json:"design_id,omitempty"`
	Session  entities.UserSession
}

type StartAdaptationOutput struct {
	Data entities.AdaptationBatch
}

func StartAdaptationUseCase(
	ctx echo.Context,
	req StartAdaptationInput,
	repo *repositories.AdaptationBatchRepository,
	event *sqs.SQS,
) (*StartAdaptationOutput, error) {
	session := req.Session
	_, err := repo.CancelActiveAdaptations(ctx.Request().Context(), session.UserID)
	if err != nil {
		return nil, shared.NewError(
			errors.CANT_CANCEL_ADAPTATIONS,
			"falha ao cancelar adaptações",
			err.Error(),
		)
	}
	adaptation := entities.AdaptationBatch{
		UserID: int64(session.UserID),
		Status: entities.AdaptationBatchStatusPending,
	}
	created, err := repo.Create(ctx.Request().Context(), adaptation)
	if err != nil {
		return nil, err
	}
	err = event.PublishAdaptation(*created)
	if err != nil {
		return nil, err
	}
	return &StartAdaptationOutput{
		Data: *created,
	}, nil
}
