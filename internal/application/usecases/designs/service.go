package designs

import (
	"algvisual/internal/infrastructure/database"

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
	in GetDesignByIdRequest,
) (*GetDesignByIdResult, error) {
	return GetDesignByIdUseCase(ctx, in, d.db, d.log)
}

func (d DesignService) CreateComponent(
	ctx echo.Context,
	req CreateComponentRequest,
) (*CreateComponentResult, error) {
	return CreateComponentUseCase(ctx, req, d.db, d.pool, d.log)
}

func (d DesignService) GetComponentsByDesignID(
	ctx echo.Context,
	req GetComponentsByDesignIdRequest,
) (*GetComponentsByDesignIdResult, error) {
	return GetComponentsByDesignIdUseCase(ctx, req, d.db)
}

func (d DesignService) ListComponentsByFieldID(
	ctx echo.Context,
	req ListComponentsByFileIdRequest,
) (*ListComponentsByFileIdResult, error) {
	return ListComponentsByFileIdUseCase(ctx, req, d.db)
}

func (d DesignService) ListDesignByProjectID(
	ctx echo.Context,
	req ListDesignByProjectIdInput,
) (*ListDesignByProjectIdOutput, error) {
	return ListDesignByProjectIdUseCase(ctx, req, d.db)
}

func (d DesignService) ListDesignElements(
	ctx echo.Context,
	req ListDesignElementsUseCaseRequest,
) (*ListDesignElementsUseCaseResult, error) {
	return ListDesignElementsUseCase(ctx, req, d.db)
}

func (d DesignService) RemoveComponent(
	ctx echo.Context,
	req RemoveComponentUseCaseRequest,
) (*RemoveComponentUseCaseResult, error) {
	return RemoveComponentUseCase(ctx, d.db, req)
}

func (d DesignService) SetBackground(
	ctx echo.Context,
	req SetBackgroundUseCaseRequest,
) (*SetBackgroundUseCaseResult, error) {
	return SetBackgroundUseCase(ctx, d.db, d.pool, req, d.log)
}
