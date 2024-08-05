package layoutgenerator

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/infra/config"
	"algvisual/internal/mapper"
	"context"
)

type ListRequestJobsInput struct {
	RequestID int32
	Limit     int32
	Skip      int32
}

type ListRequestJobsOutput struct {
	Data []entities.LayoutRequestJob
}

func ListRequestJobsUseCase(
	ctx context.Context,
	db *database.Queries,
	req ListRequestJobsInput,
) (*ListRequestJobsOutput, error) {
	if req.Limit == 0 {
		req.Limit = 10
	}
	jobs, err := db.ListLayoutRequestJobs(
		ctx,
		database.ListLayoutRequestJobsParams{Limit: req.Limit, Offset: req.Skip},
	)
	if err != nil {
		return nil, err
	}
	var domainJobs []entities.LayoutRequestJob
	for _, j := range jobs {
		domainJobs = append(domainJobs, mapper.LayoutRequestJobToDomain(j))
	}
	return &ListRequestJobsOutput{
		Data: domainJobs,
	}, nil
}

func ListLayoutRequestJobsNotStartedUseCase(
	ctx context.Context,
	db *database.Queries,
	config *config.AppConfig,
) (*ListRequestJobsOutput, error) {
	jobs, err := db.ListLayoutRequestJobsNotStarted(ctx, config.MaxWorkers)
	if err != nil {
		return nil, err
	}
	var domainJobs []entities.LayoutRequestJob
	for _, j := range jobs {
		domainJobs = append(domainJobs, mapper.LayoutRequestJobToDomain(j))
	}
	return &ListRequestJobsOutput{
		Data: domainJobs,
	}, nil
}
