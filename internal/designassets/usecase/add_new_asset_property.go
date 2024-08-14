package usecase

import (
	"algvisual/database"
	"algvisual/internal/entities"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type AddNewAssetPropertyInput struct {
	AssetID int32  `param:"asset_id" json:"asset_id,omitempty"`
	Text    string `param:"text"     json:"text,omitempty"`
}

type AddNewAssetPropertyOutput struct {
	Data entities.DesignAsset
}

func AddNewAssetPropertyUseCase(
	c echo.Context,
	req AddNewAssetPropertyInput,
	db *database.Queries,
) (*AddNewAssetPropertyOutput, error) {
	ctx := c.Request().Context()
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
