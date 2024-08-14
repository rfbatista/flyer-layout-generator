package usecase

import (
	"context"
)

type GetClientByIdInput struct {
}

type GetClientByIdOutput struct {
}

func GetClientByIdUseCase(ctx context.Context, req GetClientByIdInput) (*GetClientByIdOutput, error) {
	return &GetClientByIdOutput{}, nil
}
