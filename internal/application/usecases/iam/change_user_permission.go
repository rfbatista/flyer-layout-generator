package iam

import (
	"context"
)

type ChangeUserPermissionInput struct {
}

type ChangeUserPermissionOutput struct {
}

func  ChangeUserPermissionUseCase(ctx context.Context, req ChangeUserPermissionInput) (*ChangeUserPermissionOutput, error) {
  return &ChangeUserPermissionOutput{}, nil
} 
