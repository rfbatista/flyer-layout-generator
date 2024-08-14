package designs

import (
	"algvisual/database"
	"algvisual/internal/designs/usecase"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewDesignService(
	db *database.Queries,
	log *zap.Logger,
	pool *pgxpool.Pool,
) (*DesignService, error) {
	return &DesignService{
		db:   db,
		log:  log,
		pool: pool,
	}, nil
}

type DesignService struct {
	db   *database.Queries
	pool *pgxpool.Pool
	log  *zap.Logger
}

func (d DesignService) GetDesignByID(
	ctx echo.Context,
	in usecase.GetDesignByIdRequest,
) (*usecase.GetDesignByIdResult, error) {
	return usecase.GetDesignByIdUseCase(ctx, in, d.db, d.log)
}

func (d DesignService) CreateComponent(
	ctx echo.Context,
	req usecase.CreateComponentRequest,
) (*usecase.CreateComponentResult, error) {
	return usecase.CreateComponentUseCase(ctx, req, d.db, d.pool, d.log)
}

func (d DesignService) GetComponentsByDesignID(
	ctx echo.Context,
	req usecase.GetComponentsByDesignIdRequest,
) (*usecase.GetComponentsByDesignIdResult, error) {
	return usecase.GetComponentsByDesignIdUseCase(ctx, req, d.db)
}

func (d DesignService) ListComponentsByFieldID(
	ctx echo.Context,
	req usecase.ListComponentsByFileIdRequest,
) (*usecase.ListComponentsByFileIdResult, error) {
	return usecase.ListComponentsByFileIdUseCase(ctx, req, d.db)
}

func (d DesignService) ListDesignByProjectID(
	ctx echo.Context,
	req usecase.ListDesignByProjectIdInput,
) (*usecase.ListDesignByProjectIdOutput, error) {
	return usecase.ListDesignByProjectIdUseCase(ctx, req, d.db)
}

func (d DesignService) ListDesignElements(
	ctx echo.Context,
	req usecase.ListDesignElementsUseCaseRequest,
) (*usecase.ListDesignElementsUseCaseResult, error) {
	return usecase.ListDesignElementsUseCase(ctx, req, d.db)
}

func (d DesignService) RemoveComponent(
	ctx echo.Context,
	req usecase.RemoveComponentUseCaseRequest,
) (*usecase.RemoveComponentUseCaseResult, error) {
	return usecase.RemoveComponentUseCase(ctx, d.db, req)
}

func (d DesignService) SetBackground(
	ctx echo.Context,
	req usecase.SetBackgroundUseCaseRequest,
) (*usecase.SetBackgroundUseCaseResult, error) {
	return usecase.SetBackgroundUseCase(ctx, d.db, d.pool, req, d.log)
}
