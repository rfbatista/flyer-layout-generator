package usecase

import (
	"algvisual/database"
	"algvisual/internal/shared"
	"fmt"
	"mime/multipart"

	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type TemplatesCsvUploadRequest struct {
	Filename  string          `form:"filename"   json:"filename,omitempty"`
	File      *multipart.File `form:"file"       json:"file,omitempty"`
	ProjectID int32           `form:"project_id" json:"project_id,omitempty"`
}

type templateCsvData struct {
	Width   int32  `csv:"width"   json:"width,omitempty"`
	Height  int32  `csv:"height"  json:"height,omitempty"`
	Name    string `csv:"name"    json:"name,omitempty"`
	SlotsX  int32  `csv:"slots_x" json:"slots_x,omitempty"`
	SlostsY int32  `csv:"slots_y" json:"slosts_y,omitempty"`
}

type TemplatesCsvUploadResult struct {
	RequestID string
	Templates []database.Template
}

func TemplatesCsvUploadUseCase(
	c echo.Context,
	req TemplatesCsvUploadRequest,
	pool *pgxpool.Pool,
	db *database.Queries,
	log *zap.Logger,
) (*TemplatesCsvUploadResult, error) {
	var data []templateCsvData
	err := gocsv.UnmarshalMultipartFile(req.File, &data)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falhar ao mapear csv para estrutura", "")
		return nil, err
	}

	uniqid, _ := uuid.NewRandom()
	var templates []database.Template
	for _, t := range data {
		templ, err := CreateTemplateUseCase(c, pool, db, CreateTemplateUseCaseRequest{
			Name:      t.Name,
			Width:     int(t.Width),
			Height:    int(t.Height),
			X:         int(t.SlotsX),
			Y:         int(t.SlostsY),
			RequestID: uniqid.String(),
			ProjectID: req.ProjectID,
		}, log)
		if err != nil {
			err = shared.WrapWithAppError(
				err,
				fmt.Sprintf("Falha ao criar templates com o nome: %s", t.Name),
				"",
			)
			return nil, err
		}
		templates = append(templates, templ.Template)
	}
	return &TemplatesCsvUploadResult{
		RequestID: uniqid.String(),
		Templates: templates,
	}, nil
}
