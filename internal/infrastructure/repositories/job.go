package repositories

import (
	"algvisual/internal/application/errors"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories/mapper"
	"algvisual/internal/shared"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func NewJobRepository(db *database.Queries) (*JobRepository, error) {
	return &JobRepository{db: db}, nil
}

type JobRepository struct {
	db *database.Queries
}

type JobRepositoryGetByUserParams struct {
	Type                    entities.JobType
	FilterByPending         bool
	FilterByStarted         bool
	FilterByRenderingImages bool
}

func (a JobRepository) GetByID(
	ctx context.Context,
	id int64,
) (*entities.Job, error) {
	raw, err := a.db.GetAdaptationBatchByID(ctx, id)
	if err != nil {
		return nil, err
	}
	ent := mapper.AdaptationBatchToDomain(raw)
	return &ent, nil
}

func (a JobRepository) GetByUser(
	ctx context.Context,
	userID int32,
	params JobRepositoryGetByUserParams,
) (*entities.Job, error) {
	raw, err := a.db.GetAdaptationBatchByUser(ctx, database.GetAdaptationBatchByUserParams{
		UserID: pgtype.Int4{Int32: userID, Valid: userID != 0},
		Type: database.NullJobType{
			JobType: mapper.JobTypeToDatabase(params.Type),
			Valid:   params.Type != "",
		},
	})
	if err != nil {
		return nil, err
	}
	if raw == nil {
		return nil, shared.NewError(errors.NO_ADAPTATION_FOUND, "adaptaiton not found", "")
	}
	if len(raw) == 0 {
		return nil, shared.NewError(errors.NO_ADAPTATION_FOUND, "adaptation not found", "")
	}
	ent := mapper.AdaptationBatchToDomain(raw[0])
	return &ent, nil
}

func (a JobRepository) Create(
	ctx context.Context,
	b entities.Job,
) (*entities.Job, error) {
	id, err := a.db.CreateAdaptationBatch(ctx, database.CreateAdaptationBatchParams{
		LayoutID:  pgtype.Int4{Int32: b.LayoutID, Valid: b.LayoutID != 0},
		UserID:    pgtype.Int4{Int32: int32(b.UserID), Valid: b.UserID != 0},
		RequestID: pgtype.Int4{Int32: b.RequestID, Valid: b.RequestID != 0},
		Status: database.NullAdaptationBatchStatus{
			AdaptationBatchStatus: mapper.AdaptationBatchStatusToDatabase(b.Status),
			Valid:                 true,
		},
		Type: database.NullJobType{
			JobType: mapper.JobTypeToDatabase(b.Type),
			Valid:   b.Type != "",
		},
		StartedAt:  pgtype.Timestamp{Time: b.StartedAt, Valid: !b.StartedAt.IsZero()},
		FinishedAt: pgtype.Timestamp{Time: b.FinishedAt, Valid: !b.FinishedAt.IsZero()},
		ErrorAt:    pgtype.Timestamp{Time: b.ErrorAt, Valid: !b.ErrorAt.IsZero()},
		StoppedAt:  pgtype.Timestamp{Time: b.StoppedAt, Valid: !b.StoppedAt.IsZero()},
		UpdatedAt:  pgtype.Timestamp{Time: b.UpdatedAt, Valid: !b.UpdatedAt.IsZero()},
		Log:        pgtype.Text{String: b.Log, Valid: b.Log == ""},
	})
	if err != nil {
		return nil, err
	}
	return a.GetByID(ctx, id)
}

type JobRepositoryUpdateParams struct {
	UpdateStatus            bool
	UpdateImageURL          bool
	UpdateStartedAt         bool
	UpdateFinishedAt        bool
	UpdateErrorAt           bool
	UpdateStoppedAt         bool
	UpdateCleanedDuplicates bool
}

func (a JobRepository) Update(
	ctx context.Context,
	b entities.Job,
	p JobRepositoryUpdateParams,
) (*entities.Job, error) {
	raw, err := a.db.UpdateAdaptationBatch(ctx, database.UpdateAdaptationBatchParams{
		StatusDoUpdate: p.UpdateStatus,
		Status: database.NullAdaptationBatchStatus{
			AdaptationBatchStatus: mapper.AdaptationBatchStatusToDatabase(b.Status),
			Valid:                 p.UpdateStatus,
		},
		StartedAtDoUpdate:  p.UpdateStartedAt,
		StartedAt:          pgtype.Timestamp{Time: b.StartedAt, Valid: p.UpdateStartedAt},
		FinishedAtDoUpdate: p.UpdateFinishedAt,
		FinishedAt:         pgtype.Timestamp{Time: b.FinishedAt, Valid: p.UpdateFinishedAt},
		ErrorAtDoUpdate:    p.UpdateErrorAt,
		ErrorAt:            pgtype.Timestamp{Time: b.ErrorAt, Valid: p.UpdateErrorAt},
		StoppedAtDoUpdate:  p.UpdateStoppedAt,
		StoppedAt:          pgtype.Timestamp{Time: b.StoppedAt, Valid: p.UpdateStoppedAt},
		Log:                pgtype.Text{String: b.Log, Valid: b.Log != ""},
		AdaptationID:       pgtype.Int8{Int64: b.ID, Valid: b.ID != 0},
		RemovedDuplicates: pgtype.Bool{
			Bool:  p.UpdateCleanedDuplicates,
			Valid: p.UpdateCleanedDuplicates,
		},
		RemovedDuplicatesDoUpdate: p.UpdateCleanedDuplicates,
	})
	if err != nil {
		return nil, err
	}
	ent := mapper.AdaptationBatchToDomain(raw)
	return &ent, nil
}

func (a JobRepository) CancelActiveAdaptations(
	ctx context.Context,
	id int32,
	jobType entities.JobType,
) ([]entities.Job, error) {
	var batches []entities.Job
	raw, err := a.db.CancelActiveAdaptationBatches(
		ctx,
		database.CancelActiveAdaptationBatchesParams{
			UserID: pgtype.Int4{Int32: id, Valid: id != 0},
			Type: database.NullJobType{
				JobType: mapper.JobTypeToDatabase(jobType),
				Valid:   jobType != "",
			},
		},
	)
	if err != nil {
		return nil, err
	}
	for _, r := range raw {
		batches = append(batches, mapper.AdaptationBatchToDomain(r))
	}
	return batches, nil
}

func (a JobRepository) CloseActiveAdaptations(
	ctx context.Context,
	userID int32,
	jobType entities.JobType,
) ([]entities.Job, error) {
	var batches []entities.Job
	raw, err := a.db.CloseActiveAdaptation(
		ctx,
		database.CloseActiveAdaptationParams{
			UserID: pgtype.Int4{Int32: userID, Valid: userID != 0},
		},
	)
	if err != nil {
		return nil, err
	}
	for _, r := range raw {
		batches = append(batches, mapper.AdaptationBatchToDomain(r))
	}
	return batches, nil
}

func (a JobRepository) GetSummary(
	ctx context.Context,
	id int32,
) (*entities.JobSummary, error) {
	raw, err := a.db.GetJobSummary(
		ctx,
		pgtype.Int4{Int32: id, Valid: id != 0},
	)
	if err != nil {
		return nil, err
	}
	return &entities.JobSummary{Total: int64(raw.Total), Done: int64(raw.Done)}, nil
}
