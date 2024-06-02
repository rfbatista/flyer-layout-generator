package requestcreateimages

import (
	"algvisual/internal/database"
	"algvisual/internal/infra"
	"algvisual/internal/layoutgenerator"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type createImageRequest struct {
	RequestID string `param:"request_id" json:"request_id,omitempty"`
	DesignID  int32  `param:"design_id"  json:"design_id,omitempty"`
}

type result struct {
	IsError  bool
	ImageURL string
}

type createImageResult struct {
	Results []result
}

func createImages(
	ctx context.Context,
	req createImageRequest,
	db *database.Queries,
	client *infra.ImageGeneratorClient,
	log *zap.Logger,
	config infra.AppConfig,
	pool *pgxpool.Pool,
) (*createImageResult, error) {
	templates, err := db.GetTemplatesByRequestID(
		ctx,
		pgtype.Text{String: req.RequestID, Valid: true},
	)
	if err != nil {
		return nil, err
	}
	var results []result
	for _, t := range templates {
		log.Info("processando templates com id", zap.Int32("id", t.ID), zap.String("name", t.Name))
		res, err := layoutgenerator.GenerateDesignUseCasev2(
			ctx,
			layoutgenerator.GenerateDesignRequestv2{
				PhotoshopID: req.DesignID,
				TemplateID:  t.ID,
			},
			db,
			pool,
			config,
			log,
		)
		if err != nil {
			results = append(results, result{IsError: true})
		} else {
			results = append(results, result{IsError: false, ImageURL: res.Data.ImageURL})
			results = append(results, result{IsError: false, ImageURL: res.TwistedURL})
		}
	}
	return &createImageResult{Results: results}, nil
}
