package usecases

import (
	"context"
	"errors"
	"io"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"

	"algvisual/database"
	"algvisual/infra"
)

type SavePhotoshopFileUseCaseRequest struct {
	Filename string
	File     io.Reader
}

func SavePhotoshopFileUseCase(
	ctx context.Context,
	db *database.Queries,
	req SavePhotoshopFileUseCaseRequest,
	storage infra.Storage,
	processor *infra.PhotoshopProcessor,
	log *zap.Logger,
) (*[]database.PhotoshopElement, error) {
	name := req.Filename
	if name == "" {
		id := uuid.New()
		name = id.String()
	}
	url, err := storage.Upload(req.File, name)
	if err != nil {
		return nil, err
	}
	res, err := processor.ProcessFile(infra.ProcessFileInput{Filepath: url})
	if err != nil {
		err = errors.Join(err, errors.New("falha ao processar arquivo photoshop"))
		log.Error("falha ao processar arquivo photoshop", zap.Error(err))
		return nil, err
	}
	photoshop, err := db.CreatePhotoshop(ctx, database.CreatePhotoshopParams{
		FileUrl: pgtype.Text{String: url, Valid: true},
		Name:    name,
	})
	if err != nil {
		log.Error("falhar ao salvar metadados do arquivo photoshop", zap.Error(err))
		return nil, err
	}
	var elements []database.PhotoshopElement
	for _, i := range res.Elements {
		c, err := db.CreateElement(ctx, database.CreateElementParams{
			PhotoshopID: photoshop.ID,
			LayerID:     pgtype.Text{String: i.LayerID, Valid: true},
			Name:        pgtype.Text{String: i.Name, Valid: true},
			Text:        pgtype.Text{String: i.Text, Valid: true},
			Xi:          pgtype.Int4{Int32: int32(i.Xi), Valid: true},
			Yi:          pgtype.Int4{Int32: int32(i.Yi), Valid: true},
			Xii:         pgtype.Int4{Int32: int32(i.Xii), Valid: true},
			Yii:         pgtype.Int4{Int32: int32(i.Yii), Valid: true},
			Kind:        pgtype.Text{String: i.Kind, Valid: true},
			IsGroup:     pgtype.Bool{Bool: i.IsGroup, Valid: true},
			GroupID:     pgtype.Int4{Int32: int32(i.GroupId), Valid: true},
			Level:       pgtype.Int4{Int32: int32(i.Level), Valid: true},
			ImageUrl:    pgtype.Text{String: i.Image, Valid: true},
			Width:       pgtype.Int4{Int32: int32(i.Width), Valid: true},
			Height:      pgtype.Int4{Int32: int32(i.Height), Valid: true},
		})
		if err != nil {
			return nil, err
		}
		elements = append(elements, c)
	}
	return &elements, nil
}
