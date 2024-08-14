package usecase

import (
	"context"
)

type GetCurrentApiKeyInput struct {
}

type GetCurrentApiKeyOutput struct {
}

func GetCurrentApiKeyUseCase(ctx context.Context, req GetCurrentApiKeyInput) (*GetCurrentApiKeyOutput, error) {
	return &GetCurrentApiKeyOutput{}, nil
}
