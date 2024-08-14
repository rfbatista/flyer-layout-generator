package layoutgenerator

import (
	"algvisual/database"
	"algvisual/internal/designassets"
	"algvisual/internal/layoutgenerator/repository"
	"algvisual/internal/layoutgenerator/usecase"
	"algvisual/internal/renderer"
	"algvisual/internal/templates"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func NewLayoutGeneratorService(
	db *database.Queries,
	templateService templates.TemplatesService,
	log *zap.Logger,
	rendererService renderer.RendererService,
	pool *pgxpool.Pool,
	repo repository.LayoutRepository,
	das *designassets.DesignAssetService,
) LayoutGeneratorService {
	return LayoutGeneratorService{
		db:              db,
		templateService: templateService,
		log:             log,
		rendererService: rendererService,
		pool:            pool,
		repo:            repo,
		das:             das,
	}
}

type LayoutGeneratorService struct {
	db              *database.Queries
	templateService templates.TemplatesService
	log             *zap.Logger
	rendererService renderer.RendererService
	pool            *pgxpool.Pool
	repo            repository.LayoutRepository
	das             *designassets.DesignAssetService
}

func (l LayoutGeneratorService) GenerateNewLayout(
	ctx context.Context,
	in GenerateImageV2Input,
) (*GenerateImageV2Output, error) {
	return GenerateImageV2UseCase(
		ctx,
		in,
		l.db,
		l.templateService,
		l.log,
		l.rendererService,
		l.pool,
		l.das,
	)
}

func (l LayoutGeneratorService) UpdateElementPosition(
	ctx context.Context,
	in UpdateLayoutElementPositionInput,
) (*UpdateLayoutElementPositionOutput, error) {
	return UpdateLayoutElementPositionUseCase(ctx, in, l.db, l.rendererService, l.das)
}

func (l LayoutGeneratorService) UpdateElementSize(
	ctx context.Context,
	in UpdateLayoutElementSizeInput,
) (*UpdateLayoutElementSizeOutput, error) {
	return UpdateLayoutElementSizeUseCase(ctx, in, l.db, l.rendererService, l.das)
}

func (l LayoutGeneratorService) DeleteLayout(
	ctx context.Context,
	in usecase.DeleteLayoutByIdInput,
) (*usecase.DeleteLayoutByIdOutput, error) {
	return usecase.DeleteLayoutByIdUseCase(ctx, in, l.repo)
}

func (l LayoutGeneratorService) ZipBatchImages(
	ctx context.Context,
	in usecase.CreateZipForBatchInput,
) (*usecase.CreateZipForBatchOutput, error) {
	return usecase.CreateZipForBatchUseCase(ctx, in, l.repo)
}
