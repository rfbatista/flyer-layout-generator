package layoutgenerator

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/repositories/mapper"

	"go.uber.org/zap"
)

type GenerateImageFromLayoutInput struct {
	Layout        entities.Layout
	DesignFileURL string
}

type GenerateImageFromLayoutOutput struct {
	ImageURL string
}

func GenerateImageFromLayoutUseCase(
	log *zap.Logger,
	config config.AppConfig,
	req GenerateImageFromLayoutInput,
) (GenerateImageFromLayoutOutput, error) {
	var out GenerateImageFromLayoutOutput
	res, err := GenerateImageFromPranchetaV2(GenerateImageRequestV2{
		DesignFile: req.DesignFileURL,
		Prancheta:  mapper.LayoutToDto(req.Layout),
	}, log, config)
	if err != nil {
		return out, err
	}
	out.ImageURL = res.ImageURL
	return out, nil
}
