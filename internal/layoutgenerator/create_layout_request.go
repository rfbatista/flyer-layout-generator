package layoutgenerator

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateLayoutRequestInput struct {
	DesignID              int32    `form:"design_id"               json:"design_id,omitempty"`
	LayoutID              int32    `form:"layout_id"               json:"layout_id,omitempty"`
	LimitSizerPerElement  bool     `form:"limit_sizer_per_element" json:"limit_sizer_per_element,omitempty"`
	AnchorElements        bool     `form:"anchor_elements"         json:"anchor_elements,omitempty"`
	MinimiumComponentSize int32    `form:"minimium_component_size" json:"minimium_component_size,omitempty"`
	MinimiumTextSize      int32    `form:"minimium_text_size"      json:"minimium_text_size,omitempty"`
	Templates             []int32  `form:"templates[]"             json:"templates,omitempty"`
	Padding               int32    `form:"padding"                 json:"padding,omitempty"`
	Priority              []string `form:"priority[]"              json:"priority,omitempty"`
	KeepProportions       bool     `form:"keep_proportions"        json:"keep_proportions,omitempty"`
}

type CreateLayoutRequestOutput struct {
	Request entities.LayoutRequest
	Jobs    []entities.LayoutRequestJob
}

func CreateLayoutRequestUseCase(
	ctx context.Context,
	queries *database.Queries,
	dbx *pgxpool.Pool,
	req CreateLayoutRequestInput,
) (*CreateLayoutRequestOutput, error) {
	tx, err := dbx.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	qtx := queries.WithTx(tx)
	layoutConfig := entities.LayoutRequestConfig{
		Padding:    req.Padding,
		Priorities: entities.ListToPrioritiesMap(req.Priority),
	}
	rawConfig, err := json.Marshal(layoutConfig)
	if err != nil {
		return nil, err
	}
	layoutRes, err := qtx.CreateLayoutRequest(
		ctx,
		database.CreateLayoutRequestParams{
			DesignID: pgtype.Int4{Int32: req.DesignID, Valid: true},
			LayoutID: pgtype.Int4{Int32: req.LayoutID, Valid: true},
			Config:   pgtype.Text{String: string(rawConfig), Valid: true},
		},
	)
	if err != nil {
		return nil, err
	}
	var jobs []entities.LayoutRequestJob
	for _, tid := range req.Templates {
		templateFound, getTemplErr := qtx.GetTemplateByID(ctx, tid)
		if getTemplErr != nil {
			continue
		}
		templateDomain := mapper.TemplateToDomain(templateFound)
		for _, grid := range templateDomain.Grids() {
			c, unmarshErr := json.Marshal(entities.LayoutRequestConfig{
				LimitSizerPerElement:  req.LimitSizerPerElement,
				AnchorElements:        req.AnchorElements,
				ShowGrid:              false,
				MinimiumComponentSize: req.MinimiumComponentSize,
				MinimiumTextSize:      req.MinimiumComponentSize,
				Grid:                  grid,
				Padding:               10,
				KeepProportions:       req.KeepProportions,
				SlotsX:                grid.SlotsX,
				SlotsY:                grid.SlotsY,
				Priorities:            entities.NewLayoutRequestConfigPriority(req.Priority),
			})
			if unmarshErr != nil {
				return nil, unmarshErr
			}
			job, jerr := qtx.CreateLayoutRequestJob(ctx, database.CreateLayoutRequestJobParams{
				RequestID:  pgtype.Int4{Int32: int32(layoutRes.ID), Valid: true},
				TemplateID: pgtype.Int4{Int32: tid, Valid: true},
				DesignID:   pgtype.Int4{Int32: req.DesignID, Valid: true},
				Config:     pgtype.Text{String: string(c), Valid: true},
			})
			if jerr != nil {
				return nil, jerr
			}
			jobs = append(jobs, mapper.LayoutRequestJobToDomain(job))
		}
	}
	_, err = qtx.UpdateLayoutRequest(ctx, database.UpdateLayoutRequestParams{
		DoAddTotal:      true,
		Total:           pgtype.Int4{Int32: int32(len(jobs)), Valid: true},
		LayoutRequestID: layoutRes.ID,
	})
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return &CreateLayoutRequestOutput{
		Request: mapper.LayoutRequestToDomain(layoutRes),
		Jobs:    jobs,
	}, nil
}
