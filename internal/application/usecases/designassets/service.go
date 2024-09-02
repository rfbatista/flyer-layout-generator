package designassets

import (
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/database"
	"context"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewDesignAssetService(
	db *database.Queries,
	log *zap.Logger,
	cfg config.AppConfig,
) (*DesignAssetService, error) {
	return &DesignAssetService{db: db, log: log, cfg: cfg}, nil
}

type DesignAssetService struct {
	db  *database.Queries
	log *zap.Logger
	cfg config.AppConfig
}

func (d DesignAssetService) AddNewAssetProperty(
	ctx echo.Context,
	req AddNewAssetPropertyInput,
) (*AddNewAssetPropertyOutput, error) {
	return AddNewAssetPropertyUseCase(ctx, req, d.db)
}

func (d DesignAssetService) CreateAlternativeAsset(
	ctx echo.Context,
	req CreateAlternativeAssetInput,
) (*CreateAlternativeAssetOutput, error) {
	return CreateAlternativeAssetUseCase(ctx, req, d.db, d.log)
}

func (d DesignAssetService) GetDesignAssetByID(
	ctx context.Context,
	req GetDesignAssetByIdInput,
) (*GetDesignAssetByIdOutput, error) {
	return GetDesignAssetByIdUseCase(ctx, req, d.db)
}

func (d DesignAssetService) GetDesignAssetByProjectID(
	ctx echo.Context,
	req GetDesignAssetByProjectIdInput,
) (*GetDesignAssetByProjectIdOutput, error) {
	return GetDesignAssetByProjectIdUseCase(ctx, req, d.db)
}

func (d DesignAssetService) GetDesignAssetByDesignID(
	ctx context.Context,
	req GetDesignAssetsByDesignIdInput,
) (*GetDesignAssetsByDesignIdOutput, error) {
	return GetDesignAssetsByDesignIdUseCase(ctx, req, d.db)
}

func (d DesignAssetService) ImageUpload(
	ctx echo.Context,
	req ImageUploadRequest,
) (*ImageUploadResult, error) {
	return ImageUploadUseCase(ctx, req, d.cfg)
}

func (d DesignAssetService) ListGeneratedImages(
	ctx echo.Context,
	req ListGeneratedImagesRequest,
) (*ListGeneratedImagesResult, error) {
	return ListGeneratedImagesUseCase(ctx, req, d.db, d.log)
}

func (d DesignAssetService) SaveImage(
	ctx echo.Context,
	req SaveImageInput,
) (*SaveImageOutput, error) {
	return SaveImageUseCase(ctx, &d.cfg, req)
}
