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

func NewAdaptationBatchRepository(db *database.Queries) (*AdaptationBatchRepository, error) {
	return &AdaptationBatchRepository{db: db}, nil
}

type AdaptationBatchRepository struct {
	db *database.Queries
}

type AdaptationBatchRepositoryGetByUserParams struct {
	FilterByPending         bool
	FilterByStarted         bool
	FilterByRenderingImages bool
}

func (a AdaptationBatchRepository) GetByID(
	ctx context.Context,
	id int64,
) (*entities.AdaptationBatch, error) {
	raw, err := a.db.GetAdaptationBatchByID(ctx, id)
	if err != nil {
		return nil, err
	}
	ent := mapper.AdaptationBatchToDomain(raw)
	return &ent, nil
}

func (a AdaptationBatchRepository) GetByUser(
	ctx context.Context,
	id int32,
	params AdaptationBatchRepositoryGetByUserParams,
) (*entities.AdaptationBatch, error) {
	raw, err := a.db.GetAdaptationBatchByUser(ctx, pgtype.Int4{Int32: id, Valid: id != 0})
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

func (a AdaptationBatchRepository) Create(
	ctx context.Context,
	b entities.AdaptationBatch,
) (*entities.AdaptationBatch, error) {
	id, err := a.db.CreateAdaptationBatch(ctx, database.CreateAdaptationBatchParams{
		LayoutID:  pgtype.Int4{Int32: b.LayoutID, Valid: b.LayoutID != 0},
		UserID:    pgtype.Int4{Int32: int32(b.UserID), Valid: b.UserID != 0},
		RequestID: pgtype.Int4{Int32: b.RequestID, Valid: b.RequestID != 0},
		Status: database.NullAdaptationBatchStatus{
			AdaptationBatchStatus: mapper.AdaptationBatchStatusToDatabase(b.Status),
			Valid:                 true,
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

type AdaptationBatchRepositoryUpdateParams struct {
	UpdateStatus     bool
	UpdateImageURL   bool
	UpdateStartedAt  bool
	UpdateFinishedAt bool
	UpdateErrorAt    bool
	UpdateStoppedAt  bool
}

func (a AdaptationBatchRepository) Update(
	ctx context.Context,
	b entities.AdaptationBatch,
	p AdaptationBatchRepositoryUpdateParams,
) (*entities.AdaptationBatch, error) {
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
	})
	if err != nil {
		return nil, err
	}
	ent := mapper.AdaptationBatchToDomain(raw)
	return &ent, nil
}

func (a AdaptationBatchRepository) CancelActiveAdaptations(
	ctx context.Context,
	id int32,
) ([]entities.AdaptationBatch, error) {
	var batches []entities.AdaptationBatch
	raw, err := a.db.CancelActiveAdaptationBatches(
		ctx,
		pgtype.Int4{Int32: id, Valid: id != 0},
	)
	if err != nil {
		return nil, err
	}
	for _, r := range raw {
		batches = append(batches, mapper.AdaptationBatchToDomain(r))
	}
	return batches, nil
}

func (a AdaptationBatchRepository) CloseActiveAdaptations(
	ctx context.Context,
	id int32,
) ([]entities.AdaptationBatch, error) {
	var batches []entities.AdaptationBatch
	raw, err := a.db.CloseActiveAdaptation(
		ctx,
		pgtype.Int4{Int32: id, Valid: id != 0},
	)
	if err != nil {
		return nil, err
	}
	for _, r := range raw {
		batches = append(batches, mapper.AdaptationBatchToDomain(r))
	}
	return batches, nil
}

func (a AdaptationBatchRepository) GetSummary(
	ctx context.Context,
	id int32,
) (*entities.AdaptationSummary, error) {
	raw, err := a.db.GetAdaptationSummary(
		ctx,
		pgtype.Int4{Int32: id, Valid: id != 0},
	)
	if err != nil {
		return nil, err
	}
	return &entities.AdaptationSummary{Total: int64(raw.Total), Done: int64(raw.Done)}, nil
}
