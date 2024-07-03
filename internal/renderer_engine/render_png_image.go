package renderer_engine

import (
	"algvisual/internal/entities"
	"context"
)

type RenderPngImageInput struct {
	Layout entities.Layout
}

type RenderPngImageOutput struct {
	Data struct {
		ImageUrl string `json:"image_url,omitempty"`
	} `json:"data,omitempty"`
}

func RenderPngImageUseCase(
	ctx context.Context,
	req RenderPngImageInput,
) (*RenderPngImageOutput, error) {
	return &RenderPngImageOutput{}, nil
}
