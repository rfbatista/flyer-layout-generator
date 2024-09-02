package designassets

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type GetDesignAssetByIdInput struct {
	ID int32
}

type GetDesignAssetByIdOutput struct {
	Data entities.DesignAsset `json:"data,omitempty"`
}

func GetDesignAssetByIdUseCase(
	ctx context.Context,
	req GetDesignAssetByIdInput,
	db *database.Queries,
) (*GetDesignAssetByIdOutput, error) {
	rawAsset, err := db.GetDesignAssetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	asset := mapper.DesignAssetToDomain(rawAsset)
	rawProperties, err := db.GetDesignAssetPropertyByAssetID(
		ctx,
		pgtype.Int4{Int32: asset.ID, Valid: true},
	)
	if err != nil {
		return nil, err
	}
	var properties []entities.DesignAssetPropertyData
	for _, r := range rawProperties {
		properties = append(
			properties,
			entities.DesignAssetPropertyData{Key: r.Key, Value: r.Value},
		)
	}
	asset.Properties = properties
	return &GetDesignAssetByIdOutput{
		Data: asset,
	}, nil
}
