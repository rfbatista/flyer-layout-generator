package advertisers

import (
	"context"
)

type GetAdvertiserByIdInput struct {
}

type GetAdvertiserByIdOutput struct {
}

func  GetAdvertiserByIdUseCase(ctx context.Context, req GetAdvertiserByIdInput) (*GetAdvertiserByIdOutput, error) {
  return &GetAdvertiserByIdOutput{}, nil
} 
