package layoutgenerator

import (
	"algvisual/database"
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
) LayoutGeneratorService {
	return LayoutGeneratorService{
		db:              db,
		templateService: templateService,
		log:             log,
		rendererService: rendererService,
		pool:            pool,
	}
}

type LayoutGeneratorService struct {
	db              *database.Queries
	templateService templates.TemplatesService
	log             *zap.Logger
	rendererService renderer.RendererService
	pool            *pgxpool.Pool
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
	)
}

func (l LayoutGeneratorService) UpdateElementPosition(
	ctx context.Context,
	in UpdateLayoutElementPositionInput,
) (*UpdateLayoutElementPositionOutput, error) {
	return UpdateLayoutElementPositionUseCase(ctx, in, l.db, l.rendererService)
}

func (l LayoutGeneratorService) UpdateElementSize(
	ctx context.Context,
	in UpdateLayoutElementSizeInput,
) (*UpdateLayoutElementSizeOutput, error) {
	return UpdateLayoutElementSizeUseCase(ctx, in, l.db, l.rendererService)
}
