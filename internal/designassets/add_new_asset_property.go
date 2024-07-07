package designassets

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type AddNewAssetPropertyInput struct {
	AssetID int32  `param:"asset_id" json:"asset_id,omitempty"`
	Text    string `param:"text"     json:"text,omitempty"`
}

type AddNewAssetPropertyOutput struct {
	Data entities.DesignAsset
}

func AddNewAssetPropertyUseCase(
	ctx context.Context,
	req AddNewAssetPropertyInput,
	db *database.Queries,
) (*AddNewAssetPropertyOutput, error) {
	asset, err := GetDesignAssetByIdUseCase(ctx, GetDesignAssetByIdInput{ID: req.AssetID}, db)
	if err != nil {
		return nil, err
	}
	err = db.CreateDesignAssetProperty(ctx, database.CreateDesignAssetPropertyParams{
		AssetID: pgtype.Int4{Int32: req.AssetID, Valid: true},
		Key:     "text",
		Value:   req.Text,
	})
	if err != nil {
		return nil, err
	}
	return &AddNewAssetPropertyOutput{Data: asset.Data}, nil
}
