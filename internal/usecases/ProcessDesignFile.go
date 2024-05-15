package usecases

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/ports"
	"algvisual/internal/shared"
)

type ProcessDesignFileRequest struct {
	ID int32 `param:"id" json:"id,omitempty"`
}

type ProcessDesignFileResult struct{}

func ProcessDesignFileUseCase(
	ctx context.Context,
	req ProcessDesignFileRequest,
	processorFile ports.PhotoshopProcessorServiceProcessFile,
	log *zap.Logger,
	queries *database.Queries,
	db *pgxpool.Pool,
) (*ProcessDesignFileResult, error) {
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
	res, err := processorFile(
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
	photoshop, err := qtx.UpdateDesignByID(ctx, database.UpdateDesignByIDParams{
		DesignID:         req.ID,
		ImageUrlDoUpdate: res.ImageUrl != "",
		ImageUrl:         pgtype.Text{String: res.ImageUrl, Valid: res.ImageUrl != ""},
		WidthDoUpdate:    true,
		Width:            design.Width,
		HeightDoUpdate:   true,
		Height:           design.Height,
	})
	if err != nil {
		log.Error(
			"falhar ao salvar metadados do arquivo photoshop",
			zap.Int32("id", req.ID),
			zap.Error(err),
		)
		return nil, err
	}
	var elements []database.DesignElement
	for _, i := range res.Elements {
		c, err := qtx.CreateElement(ctx, database.CreateElementParams{
			DesignID:       photoshop.ID,
			LayerID:        pgtype.Text{String: i.LayerID, Valid: true},
			Name:           pgtype.Text{String: i.Name, Valid: true},
			Text:           pgtype.Text{String: i.Text, Valid: true},
			Xi:             pgtype.Int4{Int32: int32(i.Xi), Valid: true},
			Yi:             pgtype.Int4{Int32: int32(i.Yi), Valid: true},
			Xii:            pgtype.Int4{Int32: int32(i.Xii), Valid: true},
			Yii:            pgtype.Int4{Int32: int32(i.Yii), Valid: true},
			Kind:           pgtype.Text{String: i.Kind, Valid: true},
			IsGroup:        pgtype.Bool{Bool: i.IsGroup, Valid: true},
			GroupID:        pgtype.Int4{Int32: int32(i.GroupId), Valid: true},
			Level:          pgtype.Int4{Int32: int32(i.Level), Valid: true},
			ImageUrl:       pgtype.Text{String: i.Image, Valid: true},
			Width:          pgtype.Int4{Int32: int32(i.Width), Valid: true},
			Height:         pgtype.Int4{Int32: int32(i.Height), Valid: true},
			ImageExtension: pgtype.Text{String: i.ImageExtension, Valid: true},
		})
		if err != nil {
			return nil, err
		}
		elements = append(elements, c)
	}
	_, err = qtx.SetDesignAsProccessed(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return &ProcessDesignFileResult{}, nil
}
