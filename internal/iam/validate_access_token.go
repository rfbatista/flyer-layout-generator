package iam

import (
	"context"
)

type ValidateAccessTokenInput struct {
	AccessToken []byte
}

type ValidateAccessTokenOutput struct{}

func ValidateAccessTokenUseCase(
	ctx context.Context,
	req ValidateAccessTokenInput,
) (*ValidateAccessTokenOutput, error) {
	return &ValidateAccessTokenOutput{}, nil
}
