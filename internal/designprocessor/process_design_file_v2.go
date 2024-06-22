package designprocessor

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/infra"
	"algvisual/internal/ports"
	"algvisual/internal/shared"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type ProcessDesignFileRequestv2 struct {
	ID int32 `param:"design_id" json:"id,omitempty"`
}

type ProcessDesignFileResultv2 struct {
	Elements []entities.LayoutElement
}

func ProcessDesignFileUseCasev2(
	ctx context.Context,
	req ProcessDesignFileRequestv2,
	processorFile *infra.PhotoshopProcessor,
	log *zap.Logger,
	queries *database.Queries,
	db *pgxpool.Pool,
) (*ProcessDesignFileResultv2, error) {
	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	qtx := queries.WithTx(tx)
	design, err := qtx.Getdesign(ctx, req.ID)
	if err != nil {
		log.Error("falha buscar arquivo design", zap.Error(err))
		return nil, err
	}
	res, err := processorFile.ProcessFile(
		ports.ProcessFileInput{Filepath: design.FileUrl.String, ID: design.ID},
	)
	if err != nil {
		log.Error("falha ao processar arquivo photoshop", zap.Error(err))
		return nil, shared.WrapWithAppError(err, "falha ao processar arquivo photoshop", "")
	}
	if res.Error != "" {
		log.Error("falha ao processar arquivo photoshop", zap.String("error", res.Error))
		return nil, shared.NewAppError(500, "Falha ao processar o arquivo photoshop", res.Error)
	}
	log.Debug("image", zap.String("url", res.ImageUrl))
	photoshop, err := qtx.UpdateDesignByID(ctx, database.UpdateDesignByIDParams{
		DesignID:         req.ID,
		ImageUrlDoUpdate: res.ImageUrl != "",
		ImageUrl:         pgtype.Text{String: res.ImageUrl, Valid: res.ImageUrl != ""},
		WidthDoUpdate:    true,
		Width:            pgtype.Int4{Int32: res.Photoshop.Width, Valid: true},
		HeightDoUpdate:   true,
		Height:           pgtype.Int4{Int32: res.Photoshop.Height, Valid: true},
	})
	if err != nil {
		log.Error(
			"falhar ao salvar metadados do arquivo photoshop",
			zap.Int32("id", req.ID),
			zap.Error(err),
		)
		return nil, err
	}
	layout, err := qtx.CreateLayout(ctx, database.CreateLayoutParams{
		Width:    photoshop.Width,
		Height:   photoshop.Height,
		DesignID: pgtype.Int4{Int32: photoshop.ID, Valid: true},
	})
	if err != nil {
		log.Error("failed to create layout", zap.Error(err))
		return nil, err
	}
	for _, i := range res.Elements {
		_, err = qtx.CreateElement(ctx, database.CreateElementParams{
			DesignID:       photoshop.ID,
			LayoutID:       int32(layout.ID),
			LayerID:        pgtype.Text{String: i.LayerID, Valid: true},
			Name:           pgtype.Text{String: i.Name, Valid: true},
			Text:           pgtype.Text{String: i.Text, Valid: true},
			Xi:             pgtype.Int4{Int32: int32(i.Xi), Valid: true},
			Yi:             pgtype.Int4{Int32: int32(i.Yi), Valid: true},
			Xii:            pgtype.Int4{Int32: int32(i.Xii), Valid: true},
			Yii:            pgtype.Int4{Int32: int32(i.Yii), Valid: true},
			InnerXi:        pgtype.Int4{Int32: int32(i.InnerXi), Valid: true},
			InnerXii:       pgtype.Int4{Int32: int32(i.InnerXii), Valid: true},
			InnerYi:        pgtype.Int4{Int32: int32(i.InnerYi), Valid: true},
			InnerYii:       pgtype.Int4{Int32: int32(i.InnerYii), Valid: true},
			Kind:           pgtype.Text{String: i.Kind, Valid: true},
			IsGroup:        pgtype.Bool{Bool: i.IsGroup, Valid: true},
			GroupID:        pgtype.Int4{Int32: int32(i.GroupId), Valid: true},
			Level:          pgtype.Int4{Int32: int32(i.Level), Valid: true},
			ImageUrl:       pgtype.Text{String: i.ImageURL, Valid: true},
			Width:          pgtype.Int4{Int32: int32(i.FWidth), Valid: true},
			Height:         pgtype.Int4{Int32: int32(i.FHeight), Valid: true},
			ImageExtension: pgtype.Text{String: i.ImageExtension, Valid: true},
		})
		if err != nil {
			log.Error("failed to create element in database", zap.Error(err))
			return nil, err
		}
	}
	_, err = qtx.SetDesignAsProccessed(ctx, req.ID)
	if err != nil {
		log.Error("failed to set design as proccessed")
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		log.Error("failed to commit transaction")
		return nil, err
	}
	return &ProcessDesignFileResultv2{
		Elements: res.Elements,
	}, nil
}
