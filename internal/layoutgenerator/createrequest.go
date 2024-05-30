package layoutgenerator

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateLayoutRequestInput struct {
	DesignID              int32   `form:"design_id"               json:"design_id,omitempty"`
	LimitSizerPerElement  bool    `form:"limit_sizer_per_element" json:"limit_sizer_per_element,omitempty"`
	AnchorElements        bool    `form:"anchor_elements"         json:"anchor_elements,omitempty"`
	ShowGrid              bool    `form:"show_grid"               json:"show_grid,omitempty"`
	MinimiumComponentSize int32   `form:"minimium_component_size" json:"minimium_component_size,omitempty"`
	MinimiumTextSize      int32   `form:"minimium_text_size"      json:"minimium_text_size,omitempty"`
	Templates             []int32 `form:"templates[]"               json:"templates,omitempty"`
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
	c, err := json.Marshal(entities.LayoutRequestConfig{
		LimitSizerPerElement:  req.LimitSizerPerElement,
		AnchorElements:        req.AnchorElements,
		ShowGrid:              req.ShowGrid,
		MinimiumComponentSize: req.MinimiumComponentSize,
		MinimiumTextSize:      req.MinimiumComponentSize,
	})
	if err != nil {
		return nil, err
	}
	layoutRes, err := qtx.CreateLayoutRequest(
		ctx,
		database.CreateLayoutRequestParams{
			DesignID: pgtype.Int4{Int32: req.DesignID, Valid: true},
			Config:   pgtype.Text{String: string(c), Valid: true},
		},
	)
	if err != nil {
		return nil, err
	}
	var jobs []entities.LayoutRequestJob
	for _, tid := range req.Templates {
		job, jerr := qtx.CreateLayoutRequestJob(ctx, database.CreateLayoutRequestJobParams{
			RequestID:  pgtype.Int4{Int32: int32(layoutRes.ID), Valid: true},
			TemplateID: pgtype.Int4{Int32: tid, Valid: true},
		})
		if jerr != nil {
			return nil, jerr
		}
		jobs = append(jobs, mapper.LayoutRequestJobToDomain(job))
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
