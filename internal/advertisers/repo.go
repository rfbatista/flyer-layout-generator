package advertisers

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func NewAdvertiserRepository() (AdvertiserRepository, error) {
	return AdvertiserRepository{}, nil
}

type AdvertiserRepository struct {
	db *database.Queries
}

func (c AdvertiserRepository) GetByID() {}

func (c AdvertiserRepository) Create(
	ctx context.Context,
	e entities.Advertiser,
) (entities.Advertiser, error) {
	id, err := c.db.CreateAdvertiser(ctx, database.CreateAdvertiserParams{
		Name:      e.Name,
		CompanyID: pgtype.Int4{Int32: e.CompanyID, Valid: e.CompanyID != int32(0)},
	})
	if err != nil {
		return e, err
	}
	e.ID = id
	return e, nil
}
