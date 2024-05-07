package usecases

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/ports"
	"algvisual/internal/shared"
)

type UploadPhotoshopFileUseCaseRequest struct {
	Filename string    `json:"filename,omitempty"`
	File     io.Reader `json:"file,omitempty"`
}

type UploadPhotoshopFileUseCaseResult struct {
	Photoshop database.Design          `json:"photoshop,omitempty"`
	Elements  []database.DesignElement `json:"elements,omitempty"`
}

func UploadPhotoshopFileUseCase(
	ctx context.Context,
	db *database.Queries,
	req UploadPhotoshopFileUseCaseRequest,
	upload ports.StorageUpload,
	processorFile ports.PhotoshopProcessorServiceProcessFile,
	log *zap.Logger,
) (*UploadPhotoshopFileUseCaseResult, error) {
	name := req.Filename
	if name == "" {
		id := uuid.New()
		name = id.String()
	}
	url, err := upload(req.File, name)
	if err != nil {
		return nil, err
	}
	res, err := processorFile(ports.ProcessFileInput{Filepath: url})
	if err != nil {
		log.Error("falha ao processar arquivo photoshop", zap.Error(err))
		return nil, shared.WrapWithAppError(err, "falha ao processar arquivo photoshop", "")
	}
	if res.Error != "" {
		log.Error("falha ao processar arquivo photoshop", zap.String("error", res.Error))
		return nil, shared.NewAppError(500, "Falha ao processar o arquivo photoshop", res.Error)
	}
	photoshop, err := db.CreatePhotoshop(ctx, database.CreatePhotoshopParams{
		Width:    pgtype.Int4{Int32: res.Photoshop.Width, Valid: res.Photoshop.Width != 0},
		Height:   pgtype.Int4{Int32: res.Photoshop.Height, Valid: res.Photoshop.Height != 0},
		FileUrl:  pgtype.Text{String: url, Valid: true},
		ImageUrl: pgtype.Text{String: res.Photoshop.ImagePath, Valid: true},
		ImageExtension: pgtype.Text{
			String: res.Photoshop.ImageExtension,
			Valid:  res.Photoshop.ImageExtension != "",
		},
		Name: name,
	})
	if err != nil {
		log.Error("falhar ao salvar metadados do arquivo photoshop", zap.Error(err))
		return nil, err
	}
	var elements []database.DesignElement
	for _, i := range res.Elements {
		c, err := db.CreateElement(ctx, database.CreateElementParams{
			PhotoshopID:    photoshop.ID,
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
	return &UploadPhotoshopFileUseCaseResult{
		Elements:  elements,
		Photoshop: photoshop,
	}, nil
}
