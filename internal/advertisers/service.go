package advertisers

import (
	"algvisual/internal/advertisers/repository"
	"algvisual/internal/advertisers/usecase"

	"github.com/labstack/echo/v4"
)

func NewAdvertiserService(repo repository.AdvertiserRepository) AdvertiserService {
	return AdvertiserService{repo: repo}
}

type AdvertiserService struct {
	repo repository.AdvertiserRepository
}

func (a AdvertiserService) Create(
	ctx echo.Context,
	req usecase.CreateAdvertiserInput,
) (*usecase.CreateAdvertiserOutput, error) {
	return usecase.CreateAdvertiserUseCase(ctx, req, a.repo)
}
