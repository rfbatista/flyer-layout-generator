package advertisers

import (
	"algvisual/internal/repositories"
	"context"
)

func NewAdvertiserService() {}

type AdvertiserService struct {
	repo repositories.AdvertiserRepository
}

func (a AdvertiserService) Create(
	ctx context.Context,
	req CreateAdvertiserInput,
) (*CreateAdvertiserOutput, error) {
	return CreateAdvertiserUseCase(ctx, req, a.repo)
}
