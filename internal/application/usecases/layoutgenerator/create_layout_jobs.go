package layoutgenerator

import (
	"algvisual/internal/application/usecases/templates"
	usecase "algvisual/internal/application/usecases/templates"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/queue"
	"algvisual/internal/infrastructure/repositories"
	"algvisual/internal/shared"
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type CreateLayoutJobsUseCase struct {
	queries  *database.Queries
	dbx      *pgxpool.Pool
	tservice templates.TemplatesService
	log      *zap.Logger
	repo     *repositories.LayoutJobRepository
	queue    *queue.SQS
	cfg      config.AppConfig
}

func NewCreateLayoutJobsUsecase(
	queries *database.Queries,
	dbx *pgxpool.Pool,
	tservice templates.TemplatesService,
	log *zap.Logger,
	repo *repositories.LayoutJobRepository,
	queue *queue.SQS,
	cfg config.AppConfig,
) (*CreateLayoutJobsUseCase, error) {
	return &CreateLayoutJobsUseCase{
		queries:  queries,
		dbx:      dbx,
		tservice: tservice,
		log:      log,
		repo:     repo,
		queue:    queue,
		cfg:      cfg,
	}, nil
}

type CreateLayoutJobsInput struct {
	LayoutID              int32    `form:"layout_id"               json:"layout_id,omitempty"`
	LimitSizerPerElement  bool     `form:"limit_sizer_per_element" json:"limit_sizer_per_element,omitempty"`
	AnchorElements        bool     `form:"anchor_elements"         json:"anchor_elements,omitempty"`
	MinimiumComponentSize int32    `form:"minimium_component_size" json:"minimium_component_size,omitempty"`
	MinimiumTextSize      int32    `form:"minimium_text_size"      json:"minimium_text_size,omitempty"`
	Templates             []int32  `form:"templates[]"             json:"templates,omitempty"`
	Padding               int32    `form:"padding"                 json:"padding,omitempty"`
	Priority              []string `form:"priority[]"              json:"priority,omitempty"`
	KeepProportions       bool     `form:"keep_proportions"        json:"keep_proportions,omitempty"`
	IsAdaptation          bool     `                               json:"is_adaptation,omitempty"`
	AdaptationBatchID     int32    `                               json:"adaptation_batch_id,omitempty"`
}

type CreateLayoutJobsOutput struct {
	Request entities.ReplicationBatch
	Jobs    []entities.LayoutJob
}

func (c *CreateLayoutJobsUseCase) Execute(
	ctx context.Context,
	req CreateLayoutJobsInput,
) (*CreateLayoutJobsOutput, error) {
	c.log.Info("creating layout jobs")
	tx, err := c.dbx.Begin(ctx)
	if err != nil {
		return nil, multierr.Append(err, shared.NewInternalError("failed to start transaction"))
	}
	defer tx.Rollback(ctx)
	var jobs []entities.LayoutJob
	for _, tid := range req.Templates {
		templateFound, getTemplErr := c.tservice.GetTemplateByID(ctx, usecase.GetTemplateByIdInput{
			TemplateID: tid,
		})
		if getTemplErr != nil {
			c.log.Error("failed to find template", zap.Int("template_id", int(tid)))
			continue
		}
		templateDomain := templateFound.Data
		for _, grid := range templateDomain.Grids() {
			jobConfig := entities.LayoutJobConfig{
				LimitSizerPerElement:  req.LimitSizerPerElement,
				AnchorElements:        req.AnchorElements,
				ShowGrid:              false,
				Priorities:            entities.ListToPrioritiesMap(req.Priority),
				MinimiumComponentSize: req.MinimiumComponentSize,
				MinimiumTextSize:      req.MinimiumComponentSize,
				Grid:                  grid,
				Padding:               10,
				KeepProportions:       req.KeepProportions,
				SlotsX:                grid.SlotsX,
				SlotsY:                grid.SlotsY,
				// Priorities:            entities.NewLayoutRequestConfigPriority(req.Priority),
			}
			job := entities.LayoutJob{
				BasedOnLayoutID: req.LayoutID,
				TemplateID:      templateDomain.ID,
				Config:          jobConfig,
				Status:          entities.LayoutJobStatusPending,
				AdaptationID:    req.AdaptationBatchID,
			}
			jobCreated, jerr := c.repo.Create(ctx, job)
			if jerr != nil {
				return nil, multierr.Append(
					jerr,
					shared.NewInternalError("failed to create layout job"),
				)
			}
			jobs = append(jobs, *jobCreated)
		}
	}
	var messages []queue.ApplicationEvent

	for _, job := range jobs {
		raw, jerr := json.Marshal(job)
		if jerr != nil {
			c.log.Error("failed to marshal job", zap.Error(jerr))
			return nil, multierr.Append(jerr, shared.NewInternalError("failed to marshal job"))
		}
		messages = append(messages, queue.ApplicationEvent{
			ID:   fmt.Sprintf("%d", job.ID),
			Body: string(raw),
		})
	}
	c.log.Debug("publishing layout job batch")
	err = c.queue.PublishBatch(c.cfg.SQSConfig.LayoutJobQueue, messages)
	if err != nil {
		c.log.Error("failed to publish job batch", zap.Error(err))
		return nil, multierr.Append(err, shared.NewInternalError("failed to publish job batch"))
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, multierr.Append(err, shared.NewInternalError("failed to commit transaction"))
	}
	return &CreateLayoutJobsOutput{
		Jobs: jobs,
	}, nil
}
