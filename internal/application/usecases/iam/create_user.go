package iam

import (
	"context"
)

type CreateUserInput struct {
}

type CreateUserOutput struct {
}

func  CreateUserUseCase(ctx context.Context, req CreateUserInput) (*CreateUserOutput, error) {
  return &CreateUserOutput{}, nil
} 
