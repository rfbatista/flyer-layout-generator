package worker

import (
	"algvisual/database"
	"algvisual/internal/layoutgenerator"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func NewWorkerService(
	db *database.Queries,
	pool *pgxpool.Pool,
	layoutService layoutgenerator.LayoutGeneratorService,
	log *zap.Logger,
) WorkerService {
	return WorkerService{db: db, pool: pool, layoutService: layoutService, log: log}
}

type WorkerService struct {
	db            *database.Queries
	pool          *pgxpool.Pool
	layoutService layoutgenerator.LayoutGeneratorService
	log           *zap.Logger
}

func (w WorkerService) GenerateJob(
	ctx context.Context,
	in GenerateJobInput,
) (*GenerateJobOutput, error) {
	return GenerateJobUseCase(ctx, in, w.layoutService, w.pool, w.log, w.db)
}
