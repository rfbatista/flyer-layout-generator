package advertisers

import (
	"context"
)

func NewAdvertiserService(repo AdvertiserRepository) AdvertiserService {
	return AdvertiserService{repo: repo}
}

type AdvertiserService struct {
	repo AdvertiserRepository
}

func (a AdvertiserService) Create(
	ctx context.Context,
	req CreateAdvertiserInput,
) (*CreateAdvertiserOutput, error) {
	return CreateAdvertiserUseCase(ctx, req, a.repo)
}
