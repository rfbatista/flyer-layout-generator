package iam

import (
	"context"
)

type ListUsersInput struct {
}

type ListUsersOutput struct {
}

func  ListUsersUseCase(ctx context.Context, req ListUsersInput) (*ListUsersOutput, error) {
  return &ListUsersOutput{}, nil
} 
