package renderer

import (
	"algvisual/internal/infra"
	"context"

	"go.uber.org/zap"
)

func NewRendererService(
	storage infra.FileStorage,
	cfg *infra.AppConfig,
	log *zap.Logger,
	t *TextDrawer,
) RendererService {
	return RendererService{
		storage:    storage,
		cfg:        cfg,
		log:        log,
		textDrawer: t,
	}
}

type RendererService struct {
	storage    infra.FileStorage
	cfg        *infra.AppConfig
	log        *zap.Logger
	textDrawer *TextDrawer
}

func (r RendererService) RenderPNGImage(
	ctx context.Context,
	req RenderPngImageInput,
) (*RenderPngImageOutput, error) {
	return RenderPngImageUseCase(ctx, req, r.storage, r.cfg, r.log, r.textDrawer)
}
