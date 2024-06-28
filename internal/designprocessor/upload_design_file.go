package designprocessor

import (
	"algvisual/internal/database"
	"algvisual/internal/ports"
	"algvisual/internal/shared"
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
)

type UploadDesignFileUseCaseRequest struct {
	Filename  string    `form:"filename"   json:"filename,omitempty"`
	File      io.Reader `form:"file"       json:"file,omitempty"`
	ProjectID int32     `form:"project_id" json:"project_id,omitempty"`
}

type UploadDesignFileUseCaseResult struct {
	Design database.Design `json:"photoshop,omitempty"`
}

func UploadDesignFileUseCase(
	ctx context.Context,
	db *database.Queries,
	req UploadDesignFileUseCaseRequest,
	upload ports.StorageUpload,
	log *zap.Logger,
) (*UploadDesignFileUseCaseResult, error) {
	name := req.Filename
	if name == "" {
		id := uuid.New()
		name = id.String()
	}
	url, err := upload(req.File, name)
	if err != nil {
		log.Error("falha ao fazer upload do arquivo photoshop", zap.Error(err))
		return nil, shared.WrapWithAppError(err, "falha ao processar arquivo photoshop", "")
	}
	design, err := db.Createdesign(ctx, database.CreatedesignParams{
		Name:      name,
		FileUrl:   pgtype.Text{String: url, Valid: true},
		ProjectID: pgtype.Int4{Int32: req.ProjectID, Valid: true},
	})
	log.Info(design.FileUrl.String, zap.String("url", url))
	if err != nil {
		log.Error("falha ao processar arquivo photoshop", zap.Error(err))
		return nil, shared.WrapWithAppError(err, "falha ao processar arquivo photoshop", "")
	}
	return &UploadDesignFileUseCaseResult{
		Design: design,
	}, nil
}
