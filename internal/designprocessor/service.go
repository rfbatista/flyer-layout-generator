package designprocessor

import (
	"algvisual/database"
	"algvisual/internal/designprocessor/usecase"
	"algvisual/internal/infra"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewDesignProcessorService(
	db *database.Queries,
	log *zap.Logger,
	ph *infra.PhotoshopProcessor,
	pool *pgxpool.Pool,
	fs infra.FileStorage,
) (*DesignProcessorService, error) {
	return &DesignProcessorService{db: db, log: log, ph: ph, pool: pool, fs: fs}, nil
}

type DesignProcessorService struct {
	db   *database.Queries
	log  *zap.Logger
	ph   *infra.PhotoshopProcessor
	pool *pgxpool.Pool
	fs   infra.FileStorage
}

func (d DesignProcessorService) ListDesignFiles(
	ctx echo.Context,
	req usecase.ListDesignFilesRequest,
) (*usecase.ListDesignFilesResult, error) {
	return usecase.ListDesignFiles(ctx, req, d.db, d.log)
}

func (d DesignProcessorService) ProcessDesignFile(
	ctx echo.Context,
	req usecase.ProcessDesignFileRequest,
) (*usecase.ProcessDesignFileResult, error) {
	return usecase.ProcessDesignFileUseCase(ctx, req, d.ph.ProcessFile, d.log, d.db, d.pool)
}

func (d DesignProcessorService) ProcessDesignFileV2(
	ctx echo.Context,
	req usecase.ProcessDesignFileRequestv2,
) (*usecase.ProcessDesignFileResultv2, error) {
	return usecase.ProcessDesignFileUseCasev2(ctx, req, d.ph, d.log, d.db, d.pool)
}

func (d DesignProcessorService) UploadDesignFile(
	ctx echo.Context,
	req usecase.UploadDesignFileUseCaseRequest,
) (*usecase.UploadDesignFileUseCaseResult, error) {
	return usecase.UploadDesignFileUseCase(ctx, d.db, req, d.fs.Upload, d.log)
}
