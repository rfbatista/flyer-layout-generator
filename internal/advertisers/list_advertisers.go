package advertisers

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"
)

type ListAdvertisersInput struct {
	Page  int32 `query:"page"  json:"page,omitempty"`
	Limit int32 `query:"limit" json:"skip,omitempty"`
}

type ListAdvertisersOutput struct {
	Advertisers []entities.Advertiser `json:"advertisers,omitempty"`
}

func ListAdvertisersUseCase(
	ctx context.Context,
	req ListAdvertisersInput,
	db *database.Queries,
) (*ListAdvertisersOutput, error) {
	advertisersFromDb, err := db.ListAdvertisers(ctx, database.ListAdvertisersParams{
		Limit:  req.Limit,
		Offset: req.Page,
	})
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
