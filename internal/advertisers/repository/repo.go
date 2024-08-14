package repository

import (
	"algvisual/database"
	"algvisual/internal/entities"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func NewAdvertiserRepository(db *database.Queries) (AdvertiserRepository, error) {
	return AdvertiserRepository{db: db}, nil
}

type AdvertiserRepository struct {
	db *database.Queries
}

func (c AdvertiserRepository) GetByID() {}

func (c AdvertiserRepository) Create(
	ctx echo.Context,
	e entities.Advertiser,
) (entities.Advertiser, error) {
	id, err := c.db.CreateAdvertiser(ctx.Request().Context(), database.CreateAdvertiserParams{
		Name:      e.Name,
		CompanyID: pgtype.Int4{Int32: e.CompanyID, Valid: e.CompanyID != int32(0)},
	})
	if err != nil {
		return e, err
	}
	e.ID = id
	return e, nil
}
