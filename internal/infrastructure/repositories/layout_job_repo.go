package repositories

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories/mapper"
	"context"
	"encoding/json"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func NewLayoutJobRepository(db *database.Queries) (*LayoutJobRepository, error) {
	return &LayoutJobRepository{db: db}, nil
}

type LayoutJobRepository struct {
	db *database.Queries
}

func (l LayoutJobRepository) Create(
	ctx context.Context,
	job entities.LayoutJob,
) (*entities.LayoutJob, error) {
	rconfig, unmarshErr := json.Marshal(job.Config)
	if unmarshErr != nil {
		rconfig = []byte{}
	}
	id, err := l.db.CreateLayoutJob(ctx, database.CreateLayoutJobParams{
		BasedOnLayoutID:   pgtype.Int4{Int32: job.BasedOnLayoutID, Valid: job.BasedOnLayoutID != 0},
		AdaptationBatchID: pgtype.Int4{Int32: job.AdaptationID, Valid: job.AdaptationID != 0},
		TemplateID:        pgtype.Int4{Int32: job.TemplateID, Valid: job.TemplateID != 0},
		Status:            mapper.LayoutJobStatusToDatabase(job.Status),
		UserID:            pgtype.Int4{Int32: job.UserID, Valid: job.UserID != 0},
		StartedAt:         pgtype.Timestamp{Time: job.StartedAt, Valid: !job.StartedAt.IsZero()},
		FinishedAt:        pgtype.Timestamp{Time: job.FinishedAt, Valid: !job.FinishedAt.IsZero()},
		ErrorAt:           pgtype.Timestamp{Time: job.ErrorAt, Valid: !job.ErrorAt.IsZero()},
		UpdatedAt:         pgtype.Timestamp{Time: job.UpdatedAt, Valid: !job.UpdatedAt.IsZero()},
		Log:               pgtype.Text{String: job.Log, Valid: job.Log != ""},
		Config:            pgtype.Text{String: string(rconfig), Valid: true},
	})
	if err != nil {
		return nil, err
	}
	job.ID = id
	return &job, nil
}

func (l LayoutJobRepository) GetByID(ctx context.Context, id int64) (*entities.LayoutJob, error) {
	raw, err := l.db.GetLayoutJobByID(ctx, id)
	if err != nil {
		return nil, err
	}
	job := mapper.LayoutJobToDomain(raw)
	return &job, nil
}

type UpdateLayoutJobByIDParams struct {
	StatusDoUpdate        bool                     `json:"status_do_update,omitempty"`
	Status                entities.LayoutJobStatus `json:"status,omitempty"`
	StartedAtDoUpdate     bool                     `json:"started_at_do_update,omitempty"`
	StartedAt             time.Time                `json:"started_at,omitempty"`
	FinishedAtDoUpdate    bool                     `json:"finished_at_do_update,omitempty"`
	FinishedAt            time.Time                `json:"finished_at,omitempty"`
	ErrorAtDoUpdate       bool                     `json:"error_at_do_update,omitempty"`
	ErrorAt               time.Time                `json:"error_at,omitempty"`
	Log                   string                   `json:"log,omitempty"`
	CreatedLayoutDoUpdate bool                     `json:"created_layout_do_update,omitempty"`
	CreatedLayoutID       int32                    `json:"created_layout_id,omitempty"`
}

func (l LayoutJobRepository) UpdateLayoutJobByID(
	ctx context.Context,
	id int64,
	param UpdateLayoutJobByIDParams,
) (*entities.LayoutJob, error) {
	raw, err := l.db.UpdateLayoutJob(ctx, database.UpdateLayoutJobParams{
		StatusDoUpdate: param.StatusDoUpdate,
		Status: database.NullLayoutJobStatus{
			LayoutJobStatus: mapper.LayoutJobStatusToDatabase(param.Status),
			Valid:           param.StatusDoUpdate,
		},
		StartedAtDoUpdate:  param.StartedAtDoUpdate,
		StartedAt:          pgtype.Timestamp{Time: param.StartedAt, Valid: param.StartedAtDoUpdate},
		FinishedAtDoUpdate: param.FinishedAtDoUpdate,
		FinishedAt: pgtype.Timestamp{
			Time:  param.FinishedAt,
			Valid: param.FinishedAtDoUpdate,
		},
		ErrorAtDoUpdate:       param.ErrorAtDoUpdate,
		ErrorAt:               pgtype.Timestamp{Time: param.ErrorAt, Valid: param.ErrorAtDoUpdate},
		Log:                   pgtype.Text{String: param.Log, Valid: param.Log != ""},
		CreatedLayoutDoUpdate: param.CreatedLayoutDoUpdate,
		CreatedLayoutID: pgtype.Int4{
			Int32: param.CreatedLayoutID,
			Valid: param.CreatedLayoutDoUpdate,
		},
	})
	if err != nil {
		return nil, err
	}
	job := mapper.LayoutJobToDomain(raw)
	return &job, nil
}

type UpdateLayoutJobByParams struct {
	StatusDoUpdate          bool `json:"status_do_update,omitempty"`
	StartedAtDoUpdate       bool `json:"started_at_do_update,omitempty"`
	FinishedAtDoUpdate      bool `json:"finished_at_do_update,omitempty"`
	ErrorAtDoUpdate         bool `json:"error_at_do_update,omitempty"`
	CreatedLayoutIDDoUpdate bool
	LogDoUpdate             bool
}

func (l LayoutJobRepository) Update(
	ctx context.Context,
	job entities.LayoutJob,
	param UpdateLayoutJobByParams,
) (*entities.LayoutJob, error) {
	raw, err := l.db.UpdateLayoutJob(ctx, database.UpdateLayoutJobParams{
		ID:             pgtype.Int8{Int64: job.ID, Valid: job.ID != 0},
		StatusDoUpdate: param.StatusDoUpdate,
		Status: database.NullLayoutJobStatus{
			LayoutJobStatus: mapper.LayoutJobStatusToDatabase(job.Status),
			Valid:           param.StatusDoUpdate,
		},
		StartedAtDoUpdate:  param.StartedAtDoUpdate,
		StartedAt:          pgtype.Timestamp{Time: job.StartedAt, Valid: param.StartedAtDoUpdate},
		FinishedAtDoUpdate: param.FinishedAtDoUpdate,
		FinishedAt: pgtype.Timestamp{
			Time:  job.FinishedAt,
			Valid: param.FinishedAtDoUpdate,
		},
		ErrorAtDoUpdate: param.ErrorAtDoUpdate,
		ErrorAt:         pgtype.Timestamp{Time: job.ErrorAt, Valid: param.ErrorAtDoUpdate},
		CreatedLayoutID: pgtype.Int4{
			Int32: job.CreatedLayoutID,
			Valid: param.CreatedLayoutIDDoUpdate,
		},
		CreatedLayoutDoUpdate: param.CreatedLayoutIDDoUpdate,
		LogDoUpdated:          param.LogDoUpdate,
		Log:                   pgtype.Text{String: job.Log, Valid: param.LogDoUpdate},
	})
	if err != nil {
		return nil, err
	}
	jobCreated := mapper.LayoutJobToDomain(raw)
	return &jobCreated, nil
}

func (l LayoutJobRepository) CancelLayoutJobs(
	ctx context.Context,
	id int64,
) (*entities.LayoutJob, error) {
	return nil, nil
}
