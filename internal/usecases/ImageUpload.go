package usecases

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/google/uuid"

	"algvisual/internal/infra"
	"algvisual/internal/shared"
)

type ImageUploadRequest struct {
	Filename string                `form:"filename" json:"filename,omitempty"`
	File     *multipart.FileHeader `form:"file"     json:"file,omitempty"`
}

type ImageUploadResult struct {
	Status   string `json:"status,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}

func ImageUploadUseCase(
	ctx context.Context,
	req ImageUploadRequest,
	cfg *infra.AppConfig,
) (*ImageUploadResult, error) {
	src, err := req.File.Open()
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao abrir arquivo para upload", err.Error())
		return nil, err
	}
	defer src.Close()
	identifier := uuid.New()
	name := fmt.Sprintf("%s::%s", identifier, req.Filename)
	fullpath := fmt.Sprintf("%s/%s", cfg.ImagesFolderPath, name)
	fmt.Println(fullpath)
	dst, err := os.Create(fullpath)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao iniciar file descriptor", err.Error())
		return nil, err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		err = shared.WrapWithAppError(err, "Falha ao salvar imagem em disco", err.Error())
		return nil, err
	}
	return &ImageUploadResult{
		Status:   "success",
		ImageUrl: fmt.Sprintf("%s/%s", shared.UploadImageEndpoint, name),
	}, nil
}
