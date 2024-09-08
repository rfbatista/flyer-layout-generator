package replications

import (
	"algvisual/internal/application/errors"
	"algvisual/internal/application/usecases/layoutgenerator"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/queue"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/shared"
	"context"

	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type StartReplicationUseCase struct {
	repo          *repositories.JobRepository
	event         *queue.SQS
	cfg           config.AppConfig
	log           *zap.Logger
	createLayouts *layoutgenerator.CreateLayoutJobsUseCase
}

func NewStartReplicationUseCase(
	repo *repositories.JobRepository,
	event *queue.SQS,
	cfg config.AppConfig,
	log *zap.Logger,
	createLayouts *layoutgenerator.CreateLayoutJobsUseCase,
) (*StartReplicationUseCase, error) {
	if repo == nil {
		return nil, shared.NewInjectionError("missing adaptation batch repository")
	}
	if event == nil {
		return nil, shared.NewInjectionError("missing sqs")
	}
	return &StartReplicationUseCase{
		repo:          repo,
		event:         event,
		cfg:           cfg,
		log:           log,
		createLayouts: createLayouts,
	}, nil
}

type StartReplicationInput struct {
	LayoutID  int32                `json:"layout_id,omitempty"`
	Templates []int32              `json:"templates,omitempty"`
	Session   entities.UserSession `json:"session,omitempty"`
}

type StartReplicationOutput struct {
	Data entities.Job `json:"data"`
}

func (s StartReplicationUseCase) Execute(
	ctx context.Context,
	req StartReplicationInput,
) (*StartReplicationOutput, error) {
	session := req.Session
	_, err := s.repo.CancelActiveAdaptations(ctx, session.UserID, entities.JobTypeReplication)
	if err != nil {
		return nil, shared.NewError(
			errors.CANT_CANCEL_ADAPTATIONS,
			"fail to cancel replication",
			err.Error(),
		)
	}
	_, err = s.repo.CloseActiveAdaptations(ctx, session.UserID, entities.JobTypeReplication)
	if err != nil {
		return nil, shared.NewError(
			errors.CANT_CLOSE_ADAPTATIONS,
			"fail to close replication",
			err.Error(),
		)
	}
	replication := entities.Job{
		UserID:   int64(session.UserID),
		LayoutID: req.LayoutID,
		Status:   entities.AdaptationBatchStatusPending,
		Type:     entities.JobTypeReplication,
	}
	created, err := s.repo.Create(ctx, replication)
	if err != nil {
		s.log.Error("failed to create replication", zap.Error(err))
		return nil, multierr.Append(err, shared.NewInternalError("failed to create adaptation"))
	}
	_, err = s.createLayouts.Execute(ctx, layoutgenerator.CreateLayoutJobsInput{
		Templates: req.Templates,
		LayoutID:  req.LayoutID,
		JobID:     int32(created.ID),
	})
	if err != nil {
		return nil, err
	}
	return &StartReplicationOutput{
		Data: *created,
	}, nil
}
