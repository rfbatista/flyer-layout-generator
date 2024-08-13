package repository

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func NewRepository(db *database.Queries) LayoutRepository {
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
