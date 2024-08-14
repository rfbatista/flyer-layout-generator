package designassets

import (
	"algvisual/database"
	"algvisual/internal/designassets/usecase"
	"algvisual/internal/infra/config"
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
	req usecase.AddNewAssetPropertyInput,
) (*usecase.AddNewAssetPropertyOutput, error) {
	return usecase.AddNewAssetPropertyUseCase(ctx, req, d.db)
}

func (d DesignAssetService) CreateAlternativeAsset(
	ctx echo.Context,
	req usecase.CreateAlternativeAssetInput,
) (*usecase.CreateAlternativeAssetOutput, error) {
	return usecase.CreateAlternativeAssetUseCase(ctx, req, d.db, d.log)
}

func (d DesignAssetService) GetDesignAssetByID(
	ctx context.Context,
	req usecase.GetDesignAssetByIdInput,
) (*usecase.GetDesignAssetByIdOutput, error) {
	return usecase.GetDesignAssetByIdUseCase(ctx, req, d.db)
}

func (d DesignAssetService) GetDesignAssetByProjectID(
	ctx echo.Context,
	req usecase.GetDesignAssetByProjectIdInput,
) (*usecase.GetDesignAssetByProjectIdOutput, error) {
	return usecase.GetDesignAssetByProjectIdUseCase(ctx, req, d.db)
}

func (d DesignAssetService) GetDesignAssetByDesignID(
	ctx context.Context,
	req usecase.GetDesignAssetsByDesignIdInput,
) (*usecase.GetDesignAssetsByDesignIdOutput, error) {
	return usecase.GetDesignAssetsByDesignIdUseCase(ctx, req, d.db)
}

func (d DesignAssetService) ImageUpload(
	ctx echo.Context,
	req usecase.ImageUploadRequest,
) (*usecase.ImageUploadResult, error) {
	return usecase.ImageUploadUseCase(ctx, req, d.cfg)
}

func (d DesignAssetService) ListGeneratedImages(
	ctx echo.Context,
	req usecase.ListGeneratedImagesRequest,
) (*usecase.ListGeneratedImagesResult, error) {
	return usecase.ListGeneratedImagesUseCase(ctx, req, d.db, d.log)
}

func (d DesignAssetService) SaveImage(
	ctx echo.Context,
	req usecase.SaveImageInput,
) (*usecase.SaveImageOutput, error) {
	return usecase.SaveImageUseCase(ctx, &d.cfg, req)
}
