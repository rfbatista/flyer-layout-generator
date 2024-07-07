package designassets

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type GetDesignAssetsByDesignIdInput struct {
	DesignID int32
}

type GetDesignAssetsByDesignIdOutput struct {
	Data []entities.DesignAsset
}

func GetDesignAssetsByDesignIdUseCase(
	ctx context.Context,
	req GetDesignAssetsByDesignIdInput,
	db *database.Queries,
) (*GetDesignAssetsByDesignIdOutput, error) {
	rawAssets, err := db.GetDesignAssetByDesignID(
		ctx,
		pgtype.Int4{Int32: req.DesignID, Valid: true},
	)
	if err != nil {
		return nil, err
	}
	var assets []entities.DesignAsset
	for _, rawAsset := range rawAssets {
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
		assets = append(assets, asset)
	}
	return &GetDesignAssetsByDesignIdOutput{
		Data: assets,
	}, nil
}
