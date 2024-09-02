package advertisers

import (
	"algvisual/internal/infrastructure/repositories"

	"github.com/labstack/echo/v4"
)

func NewAdvertiserService(repo repositories.AdvertiserRepository) AdvertiserService {
	return AdvertiserService{repo: repo}
}

type AdvertiserService struct {
	repo repositories.AdvertiserRepository
}

func (a AdvertiserService) Create(
	ctx echo.Context,
	req CreateAdvertiserInput,
) (*CreateAdvertiserOutput, error) {
	return CreateAdvertiserUseCase(ctx, req, a.repo)
}
