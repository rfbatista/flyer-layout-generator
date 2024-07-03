package templates

import (
	"algvisual/database"
	"algvisual/internal/shared"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type CreateTemplateUseCaseRequest struct {
	Name      string `form:"name"   json:"name,omitempty"`
	Width     int    `form:"width"  json:"width,omitempty"`
	Height    int    `form:"height" json:"height,omitempty"`
	X         int    `              json:"x,omitempty"`
	Y         int    `              json:"y,omitempty"`
	RequestID string
	ProjectID int32
}

type CreateTemplateUseCaseResult struct {
	RequestID      string
	Template       database.Template            `json:"templates,omitempty"`
	Distortion     database.TemplatesDistortion `json:"distortion,omitempty"`
	SlotsPositions []database.TemplatesSlot     `json:"slots_positions,omitempty"`
}

func CreateTemplateUseCase(
	ctx context.Context,
	db *pgxpool.Pool,
	queries *database.Queries,
	req CreateTemplateUseCaseRequest,
	log *zap.Logger,
) (*CreateTemplateUseCaseResult, error) {
	tx, err := db.Begin(ctx)
	if err != nil {
		err = shared.WrapWithAppError(err, "cant start transaction", "")
		log.Error(err.Error())
		return nil, err
	}
	defer tx.Rollback(ctx)
	qtx := queries.WithTx(tx)
	var id string
	if req.RequestID == "" {
		uniqid, _ := uuid.NewRandom()
		id = uniqid.String()
	} else {
		id = req.RequestID
	}
	temp, err := qtx.CreateTemplateByProject(ctx, database.CreateTemplateByProjectParams{
		Name:      req.Name,
		Width:     pgtype.Int4{Int32: int32(req.Width), Valid: true},
		ProjectID: pgtype.Int4{Int32: int32(req.ProjectID), Valid: true},
		Height:    pgtype.Int4{Int32: int32(req.Height), Valid: true},
		RequestID: pgtype.Text{String: id, Valid: true},
	})
	if err != nil {
		err = shared.WrapWithAppError(err, "failed to create templates", "")
		log.Error(err.Error())
		return nil, err
	}
	dist, err := qtx.CreateTemplateDistortions(ctx, database.CreateTemplateDistortionsParams{
		X:          pgtype.Int4{Int32: int32(req.X), Valid: true},
		Y:          pgtype.Int4{Int32: int32(req.Y), Valid: true},
		TemplateID: temp.ID,
	})
	if err != nil {
		err = shared.WrapWithAppError(err, "failed to create templates distortion", "")
		log.Error(err.Error())
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		err = shared.WrapWithAppError(err, "failed to create templates distortion", "")
		log.Error(err.Error())
		return nil, err
	}
	return &CreateTemplateUseCaseResult{
		RequestID:  id,
		Template:   temp,
		Distortion: dist,
	}, nil
}
