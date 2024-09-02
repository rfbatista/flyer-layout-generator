package iam

import (
	"context"
)

type RemoveUserInput struct {
}

type RemoveUserOutput struct {
}

func  RemoveUserUseCase(ctx context.Context, req RemoveUserInput) (*RemoveUserOutput, error) {
  return &RemoveUserOutput{}, nil
} 
