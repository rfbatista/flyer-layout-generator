package advertisers

import (
	"algvisual/internal/entities"
	"context"
)

type CreateAdvertiserInput struct {
	Name      string `json:"name,omitempty"`
	CompanyID int32
}

type CreateAdvertiserOutput struct {
	Data entities.Advertiser `json:"data,omitempty"`
}

func CreateAdvertiserUseCase(
	ctx context.Context,
	req CreateAdvertiserInput,
	repo AdvertiserRepository,
) (*CreateAdvertiserOutput, error) {
	e := entities.Advertiser{
		Name:      req.Name,
		CompanyID: req.CompanyID,
	}
	eCreated, err := repo.Create(ctx, e)
	if err != nil {
		return nil, err
	}
	return &CreateAdvertiserOutput{
		Data: eCreated,
	}, nil
}
