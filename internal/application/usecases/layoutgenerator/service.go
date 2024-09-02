package layoutgenerator

import (
	"algvisual/internal/application/usecases/designassets"
	"algvisual/internal/application/usecases/renderer"
	"algvisual/internal/application/usecases/templates"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func NewLayoutGeneratorService(
	db *database.Queries,
	dbx *pgxpool.Pool,
	templateService templates.TemplatesService,
	log *zap.Logger,
	rendererService renderer.RendererService,
	pool *pgxpool.Pool,
	repo repositories.LayoutRepository,
	das *designassets.DesignAssetService,
	genLayout *GenerateLayoutUseCase,
) LayoutGeneratorService {
	return LayoutGeneratorService{
		db:              db,
		genLayout:       genLayout,
		templateService: templateService,
		log:             log,
		rendererService: rendererService,
		pool:            pool,
		repo:            repo,
		das:             das,
		dbx:             dbx,
	}
}

type LayoutGeneratorService struct {
	db              *database.Queries
	dbx             *pgxpool.Pool
	templateService templates.TemplatesService
	log             *zap.Logger
	rendererService renderer.RendererService
	pool            *pgxpool.Pool
	repo            repositories.LayoutRepository
	das             *designassets.DesignAssetService
	genLayout       *GenerateLayoutUseCase
}

func (l LayoutGeneratorService) GenerateNewLayout(
	ctx context.Context,
	in GenerateImageV2Input,
) (*GenerateImageV2Output, error) {
	return l.genLayout.Execute(
		ctx,
		in,
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
	in DeleteLayoutByIdInput,
) (*DeleteLayoutByIdOutput, error) {
	return DeleteLayoutByIdUseCase(ctx, in, l.repo)
}

func (l LayoutGeneratorService) ZipBatchImages(
	ctx context.Context,
	in CreateZipForBatchInput,
) (*CreateZipForBatchOutput, error) {
	return CreateZipForBatchUseCase(ctx, in, l.repo)
}
