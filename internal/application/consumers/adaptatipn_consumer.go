package consumers

import (
	"algvisual/internal/application/usecases/layoutgenerator"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/shared"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type AdaptatipnBatchConsumer struct {
	log              *zap.Logger
	repo             *repositories.AdaptationBatchRepository
	templatesRepo    *repositories.TemplateRepository
	createLayoutJobs *layoutgenerator.CreateLayoutJobsUseCase
}

func NewAdaptationBatchConsumer(
	log *zap.Logger,
	repo *repositories.AdaptationBatchRepository,
	createLayoutJobs *layoutgenerator.CreateLayoutJobsUseCase,
	templatesRepo *repositories.TemplateRepository,
) (*AdaptatipnBatchConsumer, error) {
	return &AdaptatipnBatchConsumer{
		log:              log,
		repo:             repo,
		createLayoutJobs: createLayoutJobs,
		templatesRepo:    templatesRepo,
	}, nil
}

type AdaptatipnBatchStartedInput struct{}

type AdaptatipnBatchStartedOutput struct{}

func (a *AdaptatipnBatchConsumer) Execute(
	ctx context.Context,
	event *shared.ApplicationEvent,
) error {
	if event == nil {
		a.log.Error("missing event")
		return errors.New("missing event")
	}
	started := time.Now()
	var batch entities.AdaptationBatch
	err := json.Unmarshal([]byte(event.Body), &batch)
	if err != nil {
		return errors.New("error in unmarshal adaptation evento")
	}
	a.log.Info("event received")
	batchFound, err := a.repo.GetByID(ctx, batch.ID)
	if err != nil {
		msg := fmt.Sprintf("adaptation with id %d was not found", batch.ID)
		a.log.Error(msg)
		return errors.New(msg)
	}
	templates, err := a.templatesRepo.List(ctx, repositories.ListTemplatesParams{
		Limit:        10,
		FilterByType: true,
		Type:         entities.TemplateTypeAdaptation,
	})
	if err != nil {
		a.log.Error("failed to create layout jobs", zap.Error(err))
		return err
	}
	var ids []int32
	for _, t := range templates {
		ids = append(ids, t.ID)
	}
	if len(ids) == 0 {
		a.log.Warn("no adaptation template was found")
		return shared.NewInternalError("no adaptation template was found")
	}
	res, err := a.createLayoutJobs.Execute(ctx, layoutgenerator.CreateLayoutJobsInput{
		LayoutID:          batchFound.LayoutID,
		Templates:         ids,
		AdaptationBatchID: int32(batchFound.ID),
	})
	if err != nil {
		a.log.Error("failed to create layout jobs", zap.Error(err))
		return err
	}
	a.log.Debug("layout jobs created", zap.Int("jobs", len(res.Jobs)))
	batchFound.Status = entities.AdaptationBatchStatusStarted
	batchFound.StartedAt = started
	_, err = a.repo.Update(ctx, *batchFound, repositories.AdaptationBatchRepositoryUpdateParams{
		UpdateStatus:    true,
		UpdateStartedAt: true,
	})
	if err != nil {
		a.log.Error("failed to update adaptation", zap.Error(err))
	}
	return nil
}
