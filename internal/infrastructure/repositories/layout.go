package repositories

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func NewLayoutRepository(db *database.Queries) LayoutRepository {
	return LayoutRepository{db: db}
}

type LayoutRepository struct {
	db *database.Queries
}

func (lr LayoutRepository) GetLayoutByID(ctx context.Context, id int64) (*entities.Layout, error) {
	layout, err := lr.db.GetLayoutByID(ctx, id)
	if err != nil {
		return nil, err
	}
	entity := mapper.LayoutToDomain(layout)
	return &entity, nil
}

func (lr LayoutRepository) DeleteLayout(ctx context.Context, l entities.Layout) error {
	return lr.db.DeleteLayoutByID(ctx, int64(l.ID))
}

func (lr LayoutRepository) SoftDeleteLayout(ctx context.Context, l entities.Layout) error {
	return lr.db.SoftDeleteLayout(ctx, int64(l.ID))
}

func (lr LayoutRepository) GetLayoutByRequestID(
	ctx context.Context,
	id int64,
) ([]entities.Layout, error) {
	layout, err := lr.db.GetLayoutByRequestID(ctx, pgtype.Int4{Int32: int32(id), Valid: true})
	if err != nil {
		return nil, err
	}
	var entities []entities.Layout
	for _, lay := range layout {
		entity := mapper.LayoutToDomain(lay)
		entities = append(entities, entity)
	}
	return entities, nil
}

func (lr LayoutRepository) GetAdvertiserByBatchID(
	ctx context.Context,
	id int64,
) (*entities.Advertiser, error) {
	res, err := lr.db.GetAdvertiserByBatchID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &entities.Advertiser{
		ID:         res.ID.Int64,
		Name:       res.Name.String,
		CompanyID:  res.CompanyID.Int32,
		CreatedAt:  &res.CreatedAt.Time,
		UpdatedAt:  &res.UpdatedAt.Time,
		DeleteedAt: &res.DeletedAt.Time,
	}, nil
}

func (lr LayoutRepository) GetClientByBatchID(
	ctx context.Context,
	id int64,
) (*entities.Client, error) {
	res, err := lr.db.GetClientByBatchID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &entities.Client{
		ID:         res.ID.Int64,
		Name:       res.Name.String,
		CompanyID:  res.CompanyID.Int32,
		CreatedAt:  &res.CreatedAt.Time,
		UpdatedAt:  &res.UpdatedAt.Time,
		DeleteedAt: &res.DeletedAt.Time,
	}, nil
}

func (lr LayoutRepository) ListLayoutsByJob(
	ctx context.Context,
	adaptationID int64,
) ([]entities.Layout, error) {
	res, err := lr.db.ListLayoutFromAdaptation(
		ctx,
		pgtype.Int4{Int32: int32(adaptationID), Valid: adaptationID != 0},
	)
	if err != nil {
		return nil, err
	}
	var list []entities.Layout
	for _, l := range res {
		list = append(list, mapper.LayoutToDomain(l))
	}
	return list, nil
}

func (lr LayoutRepository) ListLayoutsByProjectID(
	ctx context.Context,
	id int32,
) ([]entities.Layout, error) {
	res, err := lr.db.ListProjectLayoutsByProject(
		ctx,
		pgtype.Int4{Int32: int32(id), Valid: id != 0},
	)
	if err != nil {
		return nil, err
	}
	var list []entities.Layout
	for _, l := range res {
		list = append(list, mapper.LayoutToDomain(l))
	}
	return list, nil
}

func (lr LayoutRepository) SaveProjectLayout(
	ctx context.Context,
	ent entities.ProjectLayout,
) (*entities.ProjectLayout, error) {
	_, err := lr.db.SaveProjectLayout(
		ctx,
		database.SaveProjectLayoutParams{
			LayoutID:  pgtype.Int4{Int32: ent.LayoutID, Valid: ent.LayoutID != 0},
			ProjectID: pgtype.Int4{Int32: ent.ProjectID, Valid: ent.ProjectID != 0},
			DesignID:  pgtype.Int4{Int32: ent.DesignID, Valid: ent.DesignID != 0},
		},
	)
	if err != nil {
		return nil, err
	}
	return &ent, nil
}
