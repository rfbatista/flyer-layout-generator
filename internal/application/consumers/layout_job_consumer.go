package consumers

import (
	"algvisual/internal/application/usecases/layoutgenerator"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/shared"
	"context"
	"encoding/json"
	"errors"
	"time"

	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type LayoutJobConsumer struct {
	log       *zap.Logger
	repo      *repositories.LayoutJobRepository
	genLayout *layoutgenerator.GenerateLayoutUseCase
}

func NewLayoutJobConsumer(
	log *zap.Logger,
	genLayout *layoutgenerator.GenerateLayoutUseCase,
	repo *repositories.LayoutJobRepository,
) (*LayoutJobConsumer, error) {
	return &LayoutJobConsumer{
		log:       log,
		genLayout: genLayout,
		repo:      repo,
	}, nil
}

func (l *LayoutJobConsumer) Execute(
	ctx context.Context,
	event *shared.ApplicationEvent,
) error {
	if event == nil {
		l.log.Error("missing event")
		return errors.New("missing event")
	}
	// started := time.Now()
	var job entities.LayoutJob
	err := json.Unmarshal([]byte(event.Body), &job)
	if err != nil {
		return errors.New("error in unmarshal layout job event")
	}
	l.log.Info("event received")
	jobFound, err := l.repo.GetByID(ctx, job.ID)
	if err != nil || jobFound == nil {
		l.log.Error("failed to find layout by id", zap.Error(err))
		return multierr.Append(err, shared.NewInternalError("failed to find layout job by id"))
	}
	jobFound.Status = entities.LayoutJobStatusStarted
	jobFound.StartedAt = time.Now()
	_, err = l.repo.Update(ctx, *jobFound, repositories.UpdateLayoutJobByParams{
		StatusDoUpdate:    true,
		StartedAtDoUpdate: true,
	})
	if err != nil {
		l.log.Error("failed to update layout job", zap.Error(err))
		return multierr.Append(err, shared.NewInternalError("failed to update layout job"))
	}
	l.log.Debug("executing layout generation")
	lay, err := l.genLayout.Execute(ctx, layoutgenerator.GenerateImageV2Input{
		LayoutID:   jobFound.BasedOnLayoutID,
		TemplateID: jobFound.TemplateID,
		SlotsX:     jobFound.Config.SlotsX,
		SlotsY:     jobFound.Config.SlotsY,
	})
	l.log.Debug("layout generation execution finished")
	if err != nil {
		l.log.Error("failed to generate layout")
		jobFound.Status = entities.LayoutJobStatusError
		jobFound.ErrorAt = time.Now()
		jobFound.Log = err.Error()
		_, err = l.repo.Update(ctx, *jobFound, repositories.UpdateLayoutJobByParams{
			StatusDoUpdate:  true,
			LogDoUpdate:     true,
			ErrorAtDoUpdate: true,
		})
		if err != nil {
			l.log.Error("failed to update layout job", zap.Error(err))
			return multierr.Append(err, shared.NewInternalError("failed to update layout job"))
		}
		// se aconteceu algum error na geracao ja salvamos as infos no job
		return nil
	}
	l.log.Debug("layout generated with success")
	jobFound.Status = entities.LayoutJobStatusFinished
	jobFound.FinishedAt = time.Now()
	jobFound.CreatedLayoutID = lay.Layout.ID
	_, err = l.repo.Update(ctx, *jobFound, repositories.UpdateLayoutJobByParams{
		StatusDoUpdate:          true,
		FinishedAtDoUpdate:      true,
		CreatedLayoutIDDoUpdate: true,
	})
	if err != nil {
		l.log.Error("failed to update layout job", zap.Error(err))
		return multierr.Append(err, shared.NewInternalError("failed to update layout job"))
	}
	return nil
}
