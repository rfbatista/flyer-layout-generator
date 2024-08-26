package layoutgenerator

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type GetLayoytRequestInput struct {
	RequestID int32 `param:"request_id"`
}

type GetLayoytRequestOutput struct {
	Request entities.ReplicationBatch
}

func GetLayoutRequestUseCase(
	ctx context.Context,
	db *database.Queries,
	req GetLayoytRequestInput,
) (GetLayoytRequestOutput, error) {
	var out GetLayoytRequestOutput
	r, err := db.GetLayoutRequestByID(
		ctx,
		int64(req.RequestID),
	)
	if err != nil {
		return out, err
	}
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
	out.Request = ereq
	return out, nil
}
