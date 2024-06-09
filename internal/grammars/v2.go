package grammars

import (
	"algvisual/internal/entities"

	"go.uber.org/zap"
)

func RunV2(
	original entities.Layout,
	template entities.Template,
	gridX, gridY int32,
	log *zap.Logger,
) (entities.Layout, error) {
	var out entities.Layout
	return out, nil
}
