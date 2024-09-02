package advertisers

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/middlewares"
	"algvisual/internal/infrastructure/repositories/mapper"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type ListAdvertisersInput struct {
	Page  int32 `query:"page"  json:"page,omitempty"`
	Limit int32 `query:"limit" json:"skip,omitempty"`
}

type ListAdvertisersOutput struct {
	Advertisers []entities.Advertiser `json:"advertisers,omitempty"`
}

func ListAdvertisersUseCase(
	ctx echo.Context,
	req ListAdvertisersInput,
	db *database.Queries,
) (*ListAdvertisersOutput, error) {
	cc := ctx.(*middlewares.ApplicationContext)
	advertisersFromDb, err := db.ListAdvertisers(
		ctx.Request().Context(),
		database.ListAdvertisersParams{
			Limit:     req.Limit,
			Offset:    req.Page,
			CompanyID: pgtype.Int4{Int32: int32(cc.UserSession().CompanyID), Valid: true},
		},
	)
	if err != nil {
		return nil, err
	}
	var advertisers []entities.Advertiser
	for _, ad := range advertisersFromDb {
		advertisers = append(advertisers, mapper.AdvertiserToDomain(ad))
	}
	return &ListAdvertisersOutput{
		Advertisers: advertisers,
	}, nil
}
