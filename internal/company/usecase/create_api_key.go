package usecase

import (
	"context"
)

type CreateApiKeyInput struct {
	CompanyID int32
}

type CreateApiKeyOutput struct{}

func CreateApiKeyUseCase(ctx context.Context, req CreateApiKeyInput) (*CreateApiKeyOutput, error) {
	return &CreateApiKeyOutput{}, nil
}
