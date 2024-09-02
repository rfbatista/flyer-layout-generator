package designassets

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories/mapper"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type CreateAlternativeAssetInput struct {
	AssetID int32
	Text    string
}

type CreateAlternativeAssetOutput struct {
	Data database.DesignAsset
}

func CreateAlternativeAssetUseCase(
	c echo.Context,
	req CreateAlternativeAssetInput,
	db *database.Queries,
	log *zap.Logger,
) (*CreateAlternativeAssetOutput, error) {
	ctx := c.Request().Context()
	asset, err := GetDesignAssetByIdUseCase(ctx, GetDesignAssetByIdInput{ID: req.AssetID}, db)
	if err != nil {
		return nil, err
	}
	designAsset, err := db.CreateDesignAsset(
		ctx,
		database.CreateDesignAssetParams{
			ProjectID: pgtype.Int4{Int32: asset.Data.ProjectID, Valid: true},
			DesignID:  pgtype.Int4{Int32: asset.Data.DesignID, Valid: true},
			Name:      req.Text,
			Type: database.NullDesignAssetType{
				DesignAssetType: mapper.DesignAssetTypeToDB(
					entities.DesignAssetTypeText,
				),
				Valid: true,
			},
			Width:  pgtype.Int4{Int32: asset.Data.Width, Valid: true},
			Height: pgtype.Int4{Int32: asset.Data.Height, Valid: true},
		},
	)
	if err != nil {
		log.Error("failed to create design asset property", zap.Error(err))
		return nil, err
	}
	for _, p := range asset.Data.Properties {
		if p.Key == "text" {
			err := db.CreateDesignAssetProperty(ctx, database.CreateDesignAssetPropertyParams{
				AssetID: pgtype.Int4{Int32: designAsset.ID, Valid: true},
				Key:     p.Key,
				Value:   req.Text,
			})
			if err != nil {
				log.Error("failed to create design asset property", zap.Error(err))
				return nil, err
			}
			continue
		}
		err := db.CreateDesignAssetProperty(ctx, database.CreateDesignAssetPropertyParams{
			AssetID: pgtype.Int4{Int32: designAsset.ID, Valid: true},
			Key:     p.Key,
			Value:   p.Value,
		})
		if err != nil {
			log.Error("failed to create design asset property", zap.Error(err))
			return nil, err
		}
	}
	return &CreateAlternativeAssetOutput{
		Data: designAsset,
	}, nil
}
