package layoutgenerator

import (
	"algvisual/internal/database"
	"algvisual/internal/infra"
	"algvisual/internal/mapper"
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
) error {
	ctx := context.TODO()
	layoutJobReq, err := queries.StartLayoutRequest(ctx, int64(req.ID))
	if err != nil {
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
		})
		if uerr != nil {
			return uerr
		}
		return err
	}
	job := mapper.LayoutRequestJobToDomain(layoutJobReq)
	jobReq := GenerateDesignRequestv2{
		PhotoshopID: layoutReq.DesignID.Int32,
		TemplateID:  layoutJobReq.TemplateID.Int32,
	}
	if job.Config != nil {
		jobReq.Config = *job.Config
	}
	out, err := GenerateDesignUseCasev2(
		ctx,
		jobReq,
		queries,
		db,
		config,
		log,
	)
	if err != nil {
		_, uerr := queries.UpdateLayoutRequestJob(ctx, database.UpdateLayoutRequestJobParams{
			DoAddLog:           true,
			Log:                pgtype.Text{String: err.Error(), Valid: true},
			DoAddErrorAt:       true,
			ErrorAt:            pgtype.Timestamp{Time: time.Now(), Valid: true},
			LayoutRequestJobID: layoutJobReq.ID,
		})
		if uerr != nil {
			return uerr
		}
		return err
	}

	_, uerr := queries.UpdateLayoutRequestJob(ctx, database.UpdateLayoutRequestJobParams{
		DoAddFinishedAt:    true,
		FinishedAt:         pgtype.Timestamp{Time: time.Now(), Valid: true},
		LayoutRequestJobID: layoutJobReq.ID,
		DoAddImageUrl:      true,
		ImageUrl:           pgtype.Text{String: out.Data.ImageURL, Valid: true},
	})
	if uerr != nil {
		return uerr
	}
	return nil
}
