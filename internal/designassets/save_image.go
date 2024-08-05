package designassets

import (
	"algvisual/internal/infra/config"
	"algvisual/internal/shared"
	"context"
	"fmt"
	"io"
	"os"
)

type SaveImageInput struct {
	Name      string
	ImageFile io.Reader
}

type SaveImageOutput struct {
	ImageURL string `json:"image_url,omitempty"`
}

func SaveImageUseCase(
	ctx context.Context,
	cfg *config.AppConfig,
	req SaveImageInput,
) (*SaveImageOutput, error) {
	fullpath := fmt.Sprintf("%s/%s", cfg.ImagesFolderPath, req.Name)
	dst, err := os.Create(fullpath)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao iniciar file descriptor", err.Error())
		return nil, err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, req.ImageFile); err != nil {
		err = shared.WrapWithAppError(err, "Falha ao salvar imagem em disco", err.Error())
		return nil, err
	}
	return &SaveImageOutput{
		ImageURL: fmt.Sprintf("%s/%s", shared.UploadImageEndpoint, req.Name),
	}, nil
}
