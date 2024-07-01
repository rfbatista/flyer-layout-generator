package layoutgenerator

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type GetLastLayoutRequestInput struct {
	DesignID int32 `json:"design_id,omitempty" param:"design_id"`
}

type GetLastLayoutRequestOutput struct {
	Data entities.LayoutRequest `json:"data,omitempty"`
}

func GetLastLayoutRequestUseCase(
	ctx context.Context,
	req GetLastLayoutRequestInput,
	db *database.Queries,
) (*GetLastLayoutRequestOutput, error) {
	out, err := db.GetLastLayoutRequest(ctx, pgtype.Int4{Int32: req.DesignID, Valid: true})
	if err != nil {
		return nil, err
	}
	layoutRequest := mapper.LayoutRequestToDomain(out)
	rawJobs, err := db.GetRequestJobsByRequestID(
		ctx,
		pgtype.Int4{Int32: layoutRequest.ID, Valid: true},
	)
	if err != nil {
		return nil, err
	}
	var jobs []entities.LayoutRequestJob
	for _, j := range rawJobs {
		jobs = append(jobs, mapper.LayoutRequestJobToDomain(j))
	}
	layoutRequest.Jobs = jobs
	return &GetLastLayoutRequestOutput{
		Data: layoutRequest,
	}, nil
}
