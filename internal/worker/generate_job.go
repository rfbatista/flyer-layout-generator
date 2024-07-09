package worker

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/layoutgenerator"
	"algvisual/internal/mapper"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type GenerateJobInput struct {
	ID int32
}

type GenerateJobOutput struct{}

func GenerateJobUseCase(
	ctx context.Context,
	req GenerateJobInput,
	ls layoutgenerator.LayoutGeneratorService,
	db *pgxpool.Pool,
	log *zap.Logger,
	queries *database.Queries,
) (*GenerateJobOutput, error) {
	layoutJobReq, err := queries.StartLayoutRequest(ctx, int64(req.ID))
	if err != nil {
		log.Error("failed to start layout request", zap.Error(err))
		return nil, err
	}
	layoutReq, err := queries.GetLayoutRequestByID(ctx, int64(layoutJobReq.RequestID.Int32))
	if err != nil {
		log.Error("não foi possivel encontrar a requisição para gerar o layout", zap.Error(err))
		_, uerr := queries.UpdateLayoutRequestJob(ctx, database.UpdateLayoutRequestJobParams{
			DoAddLog:           true,
			Log:                pgtype.Text{String: err.Error(), Valid: true},
			DoAddErrorAt:       true,
			ErrorAt:            pgtype.Timestamp{Time: time.Now(), Valid: true},
			LayoutRequestJobID: layoutJobReq.ID,
			DoAddStatus:        true,
			Status:             pgtype.Text{String: entities.RequestStatusRunning.String()},
		})
		if uerr != nil {
			return nil, uerr
		}
		return nil, err
	}
	layoutRequest := mapper.LayoutRequestToDomain(layoutReq)
	job := mapper.LayoutRequestJobToDomain(layoutJobReq)
	jobReq := layoutgenerator.GenerateImageV2Input{
		RequestID:        int32(layoutReq.ID),
		TemplateID:       layoutJobReq.TemplateID.Int32,
		LayoutID:         layoutReq.LayoutID.Int32,
		LayoutPriorities: layoutRequest.Config.Priorities,
	}
	if job.Config != nil {
		jobReq.ShowGrid = job.Config.ShowGrid
		jobReq.SlotsX = job.Config.SlotsX
		jobReq.SlotsY = job.Config.SlotsY
		jobReq.Padding = job.Config.Padding
	}
	defer func() {
		log.Debug("closing worker")
		if r := recover(); r != nil {
			errp, ok := r.(error)
			if ok {
				log.Error("panic error in worker", zap.Error(errp))
			} else {
				log.Error("unknown panic error in worker")
			}
			_, erru := queries.UpdateLayoutRequestJob(ctx, database.UpdateLayoutRequestJobParams{
				DoAddLog:           true,
				Log:                pgtype.Text{String: errp.Error(), Valid: true},
				DoAddErrorAt:       true,
				ErrorAt:            pgtype.Timestamp{Time: time.Now(), Valid: true},
				LayoutRequestJobID: layoutJobReq.ID,
				DoAddStatus:        true,
				Status:             pgtype.Text{String: entities.RequestStatusError.String()},
			})
			if erru != nil {
				log.Warn(
					"failed to update layout request status to error after a panic",
					zap.Int("job id", int(layoutJobReq.ID)),
					zap.Error(erru),
				)
			}
			err = queries.SetJobDoneForRequest(ctx, layoutReq.ID)
			if err != nil {
				log.Warn("failed to update layout request done to error after a panic")
			}
		}
	}()
	out, err := ls.GenerateNewLayout(ctx, jobReq)
	if err != nil {
		log.Error("failed to generate design", zap.Error(err))
		_, uerr := queries.UpdateLayoutRequestJob(ctx, database.UpdateLayoutRequestJobParams{
			DoAddLog:           true,
			Log:                pgtype.Text{String: err.Error(), Valid: true},
			DoAddFinishedAt:    true,
			FinishedAt:         pgtype.Timestamp{Time: time.Now(), Valid: true},
			DoAddErrorAt:       true,
			ErrorAt:            pgtype.Timestamp{Time: time.Now(), Valid: true},
			LayoutRequestJobID: layoutJobReq.ID,
			DoAddStatus:        true,
			Status:             pgtype.Text{String: entities.RequestStatusError.String()},
		})
		if uerr != nil {
			return nil, uerr
		}
		err = queries.SetJobDoneForRequest(ctx, layoutReq.ID)
		if uerr != nil {
			return nil, uerr
		}
		return nil, err
	}

	_, uerr := queries.UpdateLayoutRequestJob(ctx, database.UpdateLayoutRequestJobParams{
		DoAddFinishedAt:    true,
		FinishedAt:         pgtype.Timestamp{Time: time.Now(), Valid: true},
		LayoutRequestJobID: layoutJobReq.ID,
		DoAddImageUrl:      true,
		ImageUrl:           pgtype.Text{String: out.ImageURL, Valid: true},
		DoAddStatus:        true,
		Status:             pgtype.Text{String: entities.RequestStatusFinished.String()},
		DoAddLayoutID:      true,
		LayoutID:           pgtype.Int4{Int32: out.Layout.ID, Valid: true},
	})
	if uerr != nil {
		return nil, uerr
	}
	err = queries.SetJobDoneForRequest(ctx, layoutReq.ID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
