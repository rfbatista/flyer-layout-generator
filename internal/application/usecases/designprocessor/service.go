package designprocessor

import (
	"algvisual/internal/infrastructure"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/storage"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewDesignProcessorService(
	db *database.Queries,
	log *zap.Logger,
	ph *infrastructure.PhotoshopProcessor,
	pool *pgxpool.Pool,
	fs storage.FileStorage,
) (*DesignProcessorService, error) {
	return &DesignProcessorService{db: db, log: log, ph: ph, pool: pool, fs: fs}, nil
}

type DesignProcessorService struct {
	db   *database.Queries
	log  *zap.Logger
	ph   *infrastructure.PhotoshopProcessor
	pool *pgxpool.Pool
	fs   storage.FileStorage
}

func (d DesignProcessorService) ListDesignFiles(
	ctx echo.Context,
	req ListDesignFilesRequest,
) (*ListDesignFilesResult, error) {
	return ListDesignFiles(ctx, req, d.db, d.log)
}

func (d DesignProcessorService) ProcessDesignFile(
	ctx echo.Context,
	req ProcessDesignFileRequest,
) (*ProcessDesignFileResult, error) {
	return ProcessDesignFileUseCase(ctx, req, d.ph.ProcessFile, d.log, d.db, d.pool)
}

func (d DesignProcessorService) ProcessDesignFileV2(
	ctx echo.Context,
	req ProcessDesignFileRequestv2,
) (*ProcessDesignFileResultv2, error) {
	return ProcessDesignFileUseCasev2(ctx, req, d.ph, d.log, d.db, d.pool)
}

func (d DesignProcessorService) UploadDesignFile(
	ctx echo.Context,
	req UploadDesignFileUseCaseRequest,
) (*UploadDesignFileUseCaseResult, error) {
	return UploadDesignFileUseCase(ctx, d.db, req, d.fs, d.log)
}
