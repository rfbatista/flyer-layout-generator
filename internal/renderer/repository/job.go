package repository

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type RendererJobRepository struct {
	db *database.Queries
}

func (r RendererJobRepository) Create(
	ctx context.Context,
	job entities.RenderJob,
) (*entities.RenderJob, error) {
	e, err := r.db.CreateRendererJob(ctx, database.CreateRendererJobParams{
		LayoutID:     pgtype.Int4{Int32: job.LayoutID, Valid: job.LayoutID != 0},
		AdaptationID: pgtype.Int4{Int32: job.AdaptationID, Valid: job.AdaptationID != 0},
	})
	if err != nil {
		return nil, err
	}
	job.ID = e
	return &job, nil
}

type RenderJobListParams struct {
	AdaptationID   int64
	FilterByStatus bool
	Status         entities.RenderJobStatus
}

func (r RendererJobRepository) List(
	ctx context.Context,
	params RenderJobListParams,
) ([]entities.RenderJob, error) {
	var jobs []entities.RenderJob
	raw, err := r.db.ListRendererJobs(ctx, database.ListRendererJobsParams{
		AdaptationID: pgtype.Int4{
			Int32: int32(params.AdaptationID),
			Valid: params.AdaptationID != 0,
		},
		FilterByStatus: params.FilterByStatus,
		Status: database.NullRendererJobStatus{
			RendererJobStatus: mapper.RenderJobStatusToDatabase(params.Status),
			Valid:             params.FilterByStatus,
		},
	})
	if err != nil {
		return jobs, err
	}
	for _, r := range raw {
		jobs = append(jobs, mapper.RendererJobToDomain(r))
	}
	return jobs, nil
}

type RendererJobRepositoryUpdateParams struct {
	UpdateStatus       bool                     `json:"update_status,omitempty"`
	Status             entities.RenderJobStatus `json:"status,omitempty"`
	ImageDoUpdate      bool                     `json:"image_do_update,omitempty"`
	ImageID            int64                    `json:"image_id,omitempty"`
	StartedAtDoUpdate  bool                     `json:"started_at_do_update,omitempty"`
	StartedAt          time.Time                `json:"started_at,omitempty"`
	FinishedAtDoUpdate bool                     `json:"finished_at_do_update,omitempty"`
	FinishedAt         time.Time                `json:"finished_at,omitempty"`
	ErrorAtDoUpdate    bool                     `json:"error_at_do_update,omitempty"`
	ErrorAt            time.Time                `json:"error_at,omitempty"`
	StoppedAtDoUpdate  bool                     `json:"stopped_at_do_update,omitempty"`
	LogDoUpdate        bool                     `json:"log_do_update,omitempty"`
	Log                string                   `json:"log,omitempty"`
}

func (r RendererJobRepository) Update(
	ctx context.Context,
	e entities.RenderJob,
	params RendererJobRepositoryUpdateParams,
) (*entities.RenderJob, error) {
	raw, err := r.db.UpdateRendererJob(ctx, database.UpdateRendererJobParams{
		StatusDoUpdate: params.UpdateStatus,
		Status: database.NullRendererJobStatus{
			RendererJobStatus: mapper.RenderJobStatusToDatabase(params.Status),
			Valid:             params.UpdateStatus,
		},
		ImageDoUpdate:     params.ImageDoUpdate,
		ImageID:           pgtype.Int4{Int32: int32(params.ImageID), Valid: params.ImageDoUpdate},
		StartedAtDoUpdate: params.StartedAtDoUpdate,
		StartedAt: pgtype.Timestamp{
			Time:  params.StartedAt,
			Valid: params.StartedAtDoUpdate,
		},
		FinishedAtDoUpdate: params.FinishedAtDoUpdate,
		FinishedAt: pgtype.Timestamp{
			Time:  params.FinishedAt,
			Valid: params.FinishedAtDoUpdate,
		},
		ErrorAtDoUpdate: params.ErrorAtDoUpdate,
		ErrorAt:         pgtype.Timestamp{Time: params.ErrorAt, Valid: params.ErrorAtDoUpdate},
		LogDoUpdate:     params.LogDoUpdate,
		Log:             pgtype.Text{String: params.Log, Valid: params.LogDoUpdate},
	})
	if err != nil {
		return nil, err
	}
	ent := mapper.RendererJobToDomain(raw)
	return &ent, nil
}
