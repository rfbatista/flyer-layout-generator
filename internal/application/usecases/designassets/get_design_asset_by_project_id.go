package designassets

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories/mapper"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type GetDesignAssetByProjectIdInput struct {
	ProjectID int32 `json:"project_id,omitempty" param:"project_id"`
}

type GetDesignAssetByProjectIdOutput struct {
	Data []entities.DesignAsset `json:"data,omitempty"`
}

func GetDesignAssetByProjectIdUseCase(
	ctx echo.Context,
	req GetDesignAssetByProjectIdInput,
	db *database.Queries,
) (*GetDesignAssetByProjectIdOutput, error) {
	var assets []entities.DesignAsset
	rawAssets, err := db.GetDesignAssetByProjectID(
		ctx.Request().Context(),
		pgtype.Int4{Int32: req.ProjectID, Valid: true},
	)
	if err != nil {
		return nil, err
	}
	for _, rawAsset := range rawAssets {
		asset := mapper.DesignAssetToDomain(rawAsset)
		rawProperties, err := db.GetDesignAssetPropertyByAssetID(
			ctx.Request().Context(),
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
	return &GetDesignAssetByProjectIdOutput{
		Data: assets,
	}, nil
}
