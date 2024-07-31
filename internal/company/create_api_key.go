package company

import (
	"context"
)

type CreateApiKeyInput struct {
}

type CreateApiKeyOutput struct {
}

func  CreateApiKeyUseCase(ctx context.Context, req CreateApiKeyInput) (*CreateApiKeyOutput, error) {
  return &CreateApiKeyOutput{}, nil
} 
