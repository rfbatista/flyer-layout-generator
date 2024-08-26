package layoutgenerator

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type ListLayoytRequestInput struct {
	RequestID int32 `param:"request_id"`
	Limit     int32 `query:"limit"`
	Skip      int32 `query:"limit"`
}

type ListLayoytRequestOutput struct {
	Requests []entities.ReplicationBatch
}

func ListLayoutRequestUseCase(
	ctx context.Context,
	db *database.Queries,
	req ListLayoytRequestInput,
) (ListLayoytRequestOutput, error) {
	var out ListLayoytRequestOutput
	if req.Limit == 0 {
		req.Limit = 10
	}
	requests, err := db.ListLayoutRequests(
		ctx,
		database.ListLayoutRequestsParams{Limit: req.Limit, Offset: req.Skip},
	)
	if err != nil {
		return out, err
	}
	for _, r := range requests {
		jobs, err := db.GetRequestJobsByRequestID(ctx, pgtype.Int4{Int32: int32(r.ID), Valid: true})
		if err != nil {
			return out, err
		}
		var ejobs []entities.LayoutRequestJob
		for _, j := range jobs {
			ejobs = append(ejobs, mapper.LayoutRequestJobToDomain(j))
		}
		ereq := mapper.LayoutRequestToDomain(r)
		ereq.Jobs = ejobs
		out.Requests = append(out.Requests, ereq)
	}
	return out, nil
}
