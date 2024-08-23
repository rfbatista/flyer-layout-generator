package repositories

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
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
	Status     entities.AdaptationBatchStatus
	DoByStatus bool
}

func (a AdaptationBatchRepository) GetByUser(
	ctx context.Context,
	user entities.User,
	params AdaptationBatchRepositoryGetByUserParams,
) (*entities.AdaptationBatch, error) {
	_, err := a.db.GetAdaptationBatchByUser(ctx, database.GetAdaptationBatchByUserParams{
		UserID:         pgtype.Int4{Int32: user.ID, Valid: user.ID != 0},
		FilterByStatus: params.DoByStatus,
		Status: database.NullAdaptationBatchStatus{
			AdaptationBatchStatus: mapper.AdaptationBatchStatusToDatabase(params.Status),
			Valid:                 params.DoByStatus,
		},
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (a AdaptationBatchRepository) Create(
	ctx context.Context,
	b entities.AdaptationBatch,
) (*entities.AdaptationBatch, error) {
	id, err := a.db.CreateAdaptationBatch(ctx, database.CreateAdaptationBatchParams{
		LayoutID:   pgtype.Int4{Int32: b.LayoutID, Valid: b.LayoutID != 0},
		DesignID:   pgtype.Int4{Int32: b.DesignID, Valid: b.DesignID != 0},
		RequestID:  pgtype.Int4{Int32: b.RequestID, Valid: b.RequestID != 0},
		TemplateID: pgtype.Int4{Int32: b.TemplateID, Valid: b.TemplateID != 0},
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
	b.ID = id
	return &b, nil
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
	})
	if err != nil {
		return nil, err
	}
	ent := mapper.AdaptationBatchToDomain(raw)
	return &ent, nil
}
