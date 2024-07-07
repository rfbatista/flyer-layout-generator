package layoutgenerator

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/infra"
	"algvisual/internal/mapper"
	"algvisual/internal/renderer"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type StartRequestJobInput struct {
	ID int32
}

type StartRequestJobOutput struct{}

func StartRequestJobUseCase(
	client *infra.ImageGeneratorClient,
	queries *database.Queries,
	db *pgxpool.Pool,
	config infra.AppConfig,
	log *zap.Logger,
	req StartRequestJobInput,
	render renderer.RendererService,
) error {
	ctx := context.TODO()
	layoutJobReq, err := queries.StartLayoutRequest(ctx, int64(req.ID))
	if err != nil {
		log.Error("failed to start layout request", zap.Error(err))
		return err
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
			return uerr
		}
		return err
	}
	job := mapper.LayoutRequestJobToDomain(layoutJobReq)
	jobReq := GenerateImage{
		DesignID:   layoutReq.DesignID.Int32,
		TemplateID: layoutJobReq.TemplateID.Int32,
		LayoutID:   layoutReq.LayoutID.Int32,
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
	out, err := GenerateImageUseCase(
		ctx,
		jobReq,
		queries,
		db,
		config,
		log,
		render,
	)
	if err != nil {
		log.Error("failed to generate design", zap.Error(err))
		_, uerr := queries.UpdateLayoutRequestJob(ctx, database.UpdateLayoutRequestJobParams{
			DoAddLog:           true,
			Log:                pgtype.Text{String: err.Error(), Valid: true},
			DoAddErrorAt:       true,
			ErrorAt:            pgtype.Timestamp{Time: time.Now(), Valid: true},
			LayoutRequestJobID: layoutJobReq.ID,
			DoAddStatus:        true,
			Status:             pgtype.Text{String: entities.RequestStatusError.String()},
		})
		if uerr != nil {
			return uerr
		}
		return err
	}

	var layoutId int32
	if out.Layout != nil {
		layoutId = out.Layout.ID
	}

	_, uerr := queries.UpdateLayoutRequestJob(ctx, database.UpdateLayoutRequestJobParams{
		DoAddFinishedAt:    true,
		FinishedAt:         pgtype.Timestamp{Time: time.Now(), Valid: true},
		LayoutRequestJobID: layoutJobReq.ID,
		DoAddImageUrl:      true,
		ImageUrl:           pgtype.Text{String: out.Data.ImageURL, Valid: true},
		DoAddStatus:        true,
		Status:             pgtype.Text{String: entities.RequestStatusFinished.String()},
		DoAddLayoutID:      true,
		LayoutID:           pgtype.Int4{Int32: layoutId, Valid: true},
	})
	if uerr != nil {
		return uerr
	}
	err = queries.SetJobDoneForRequest(ctx, layoutReq.ID)
	if err != nil {
		return err
	}
	return nil
}
