package requestcreateimages

import (
	"algvisual/internal/database"
	"algvisual/internal/infra"
	"algvisual/internal/usecases"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
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
		log.Info("processando template com id", zap.Int32("id", t.ID), zap.String("name", t.Name))
		res, err := usecases.GenerateDesignUseCasev2(ctx, usecases.GenerateDesignRequestv2{
			PhotoshopID: req.DesignID,
			TemplateID:  t.ID,
		}, client, db, config, log)
		if err != nil {
			results = append(results, result{IsError: true})
		} else {
			results = append(results, result{IsError: false, ImageURL: res.Data.ImageURL})
		}
	}
	return &createImageResult{Results: results}, nil
}
