package usecase

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/templates"
	"algvisual/internal/templates/usecase"
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

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
	IsAdaptation          bool
}

type CreateLayoutJobsOutput struct {
	Request entities.ReplicationBatch
	Jobs    []entities.LayoutJob
}

func CreateLayoutJobsUseCase(
	ctx context.Context,
	queries *database.Queries,
	dbx *pgxpool.Pool,
	req CreateLayoutJobsInput,
	tservice templates.TemplatesService,
) (*CreateLayoutJobsOutput, error) {
	tx, err := dbx.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	qtx := queries.WithTx(tx)
	var jobs []entities.LayoutJob
	for _, tid := range req.Templates {
		templateFound, getTemplErr := tservice.GetTemplateByID(ctx, usecase.GetTemplateByIdInput{
			TemplateID: tid,
		})
		if getTemplErr != nil {
			continue
		}
		templateDomain := templateFound.Data
		for _, grid := range templateDomain.Grids() {
			c, unmarshErr := json.Marshal(entities.LayoutRequestConfig{
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
			})
			if unmarshErr != nil {
				return nil, unmarshErr
			}
			job := entities.LayoutJob{
				BasedOnLayoutID: req.LayoutID,
				TemplateID:      tid,
				Config:          string(c),
			}
			id, jerr := qtx.CreateLayoutJob(ctx, database.CreateLayoutJobParams{
				BasedOnLayoutID: pgtype.Int4{Int32: job.BasedOnLayoutID, Valid: true},
				Config:          pgtype.Text{String: job.Config, Valid: true},
				TemplateID:      pgtype.Int4{Int32: job.TemplateID, Valid: job.TemplateID != 0},
			})
			if jerr != nil {
				return nil, jerr
			}
			job.ID = id
			jobs = append(jobs, job)
		}
	}
	_, err = qtx.UpdateLayoutRequest(ctx, database.UpdateLayoutRequestParams{
		DoAddTotal: true,
		Total:      pgtype.Int4{Int32: int32(len(jobs)), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return &CreateLayoutJobsOutput{
		Jobs: jobs,
	}, nil
}
