package designprocessor

import (
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/middlewares"
	"algvisual/internal/infrastructure/storage"
	"algvisual/internal/shared"
	"io"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
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
	c echo.Context,
	db *database.Queries,
	req UploadDesignFileUseCaseRequest,
	upload storage.FileStorage,
	log *zap.Logger,
) (*UploadDesignFileUseCaseResult, error) {
	session := c.(*middlewares.ApplicationContext)
	ctx := c.Request().Context()
	name := req.Filename
	if name == "" {
		id := uuid.New()
		name = id.String()
	}
	url, err := upload.Upload(req.File, name)
	if err != nil {
		log.Error("falha ao fazer upload do arquivo photoshop", zap.Error(err))
		return nil, shared.WrapWithAppError(err, "falha ao processar arquivo photoshop", "")
	}
	design, err := db.Createdesign(ctx, database.CreatedesignParams{
		Name:      name,
		FileUrl:   pgtype.Text{String: url, Valid: true},
		CompanyID: pgtype.Int4{Int32: int32(session.UserSession().CompanyID), Valid: true},
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
