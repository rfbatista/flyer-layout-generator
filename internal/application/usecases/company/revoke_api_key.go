package company

import (
	"context"
)

type RevokeApiKeyInput struct {
}

type RevokeApiKeyOutput struct {
}

func RevokeApiKeyUseCase(ctx context.Context, req RevokeApiKeyInput) (*RevokeApiKeyOutput, error) {
	return &RevokeApiKeyOutput{}, nil
}
