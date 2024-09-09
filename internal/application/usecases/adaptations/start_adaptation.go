package adaptations

import (
	"algvisual/internal/application/errors"
	"algvisual/internal/application/usecases/layoutgenerator"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/queue"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/shared"
	"context"
	"time"

	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type StartAdaptationUseCase struct {
	repo          *repositories.JobRepository
	jrepo         *repositories.JobRepository
	event         *queue.SQS
	cfg           config.AppConfig
	log           *zap.Logger
	createLayouts *layoutgenerator.CreateLayoutJobsUseCase
	templatesRepo *repositories.TemplateRepository
}

func NewStartAdaptationUseCase(
	repo *repositories.JobRepository,
	event *queue.SQS,
	cfg config.AppConfig,
	log *zap.Logger,
	createLayouts *layoutgenerator.CreateLayoutJobsUseCase,
	templatesRepo *repositories.TemplateRepository,
	jrepo *repositories.JobRepository,
) (*StartAdaptationUseCase, error) {
	if repo == nil {
		return nil, shared.NewInjectionError("missing adaptation batch repository")
	}
	if event == nil {
		return nil, shared.NewInjectionError("missing sqs")
	}
	return &StartAdaptationUseCase{
		jrepo:         jrepo,
		createLayouts: createLayouts,
		templatesRepo: templatesRepo,
		repo:          repo,
		event:         event,
		cfg:           cfg,
		log:           log,
	}, nil
}

type StartAdaptationInput struct {
	LayoutID   int32                `json:"layout_id,omitempty"`
	Session    entities.UserSession `json:"session,omitempty"`
	Priorities []string             `json:"priorities,omitempty"`
}

type StartAdaptationOutput struct {
	Data entities.Job `json:"data"`
}

func (s StartAdaptationUseCase) Execute(
	ctx context.Context,
	req StartAdaptationInput,
) (*StartAdaptationOutput, error) {
	session := req.Session
	_, err := s.repo.CancelActiveAdaptations(ctx, session.UserID, entities.JobTypeAdaptation)
	if err != nil {
		return nil, shared.NewError(
			errors.CANT_CANCEL_ADAPTATIONS,
			"fail to cancel adaptation",
			err.Error(),
		)
	}
	_, err = s.repo.CloseActiveAdaptations(ctx, session.UserID, entities.JobTypeAdaptation)
	if err != nil {
		return nil, shared.NewError(
			errors.CANT_CLOSE_ADAPTATIONS,
			"fail to close adaptation",
			err.Error(),
		)
	}
	adaptation := entities.Job{
		UserID:   int64(session.UserID),
		LayoutID: req.LayoutID,
		Status:   entities.AdaptationBatchStatusPending,
		Type:     entities.JobTypeAdaptation,
	}
	created, err := s.repo.Create(ctx, adaptation)
	if err != nil {
		s.log.Error("failed to create adaptation", zap.Error(err))
		return nil, multierr.Append(err, shared.NewInternalError("failed to create adaptation"))
	}
	templates, err := s.templatesRepo.List(ctx, repositories.ListTemplatesParams{
		Limit:              10,
		FilterByType:       true,
		Type:               entities.TemplateTypeAdaptation,
		AddPublicTemplates: false,
	})
	if err != nil {
		s.log.Error("failed to create layout jobs", zap.Error(err))
		return nil, err
	}
	var templatesId []int32
	for _, t := range templates {
		templatesId = append(templatesId, t.ID)
	}
	started := time.Now()
	_, err = s.createLayouts.Execute(ctx, layoutgenerator.CreateLayoutJobsInput{
		Templates: templatesId,
		LayoutID:  req.LayoutID,
		JobID:     int32(created.ID),
		Priority:  req.Priorities,
	})
	if err != nil {
		return nil, err
	}
	created.Status = entities.AdaptationBatchStatusStarted
	created.StartedAt = started
	_, err = s.jrepo.Update(ctx, *created, repositories.JobRepositoryUpdateParams{
		UpdateStatus:    true,
		UpdateStartedAt: true,
	})
	if err != nil {
		s.log.Error("failed to update adaptation", zap.Error(err))
	}
	return &StartAdaptationOutput{
		Data: *created,
	}, nil
}
