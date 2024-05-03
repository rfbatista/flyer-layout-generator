package usecases

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/shared"
)

type CreateTemplateUseCaseRequest struct {
	Name           string                            `form:"name"            json:"name,omitempty"`
	Width          int                               `form:"width"           json:"width,omitempty"`
	Height         int                               `form:"height"          json:"height,omitempty"`
	Type           entities.TemplateType             `form:"type"            json:"type,omitempty"`
	SlotsPositions []entities.TemplateSlotsPositions `form:"slots_positions" json:"slots_positions,omitempty"`
	Distortion     entities.TemplateDistortion       `form:"distortion"      json:"distortion,omitempty"`
	X              int                               `                       json:"x,omitempty"`
	Y              int                               `                       json:"y,omitempty"`
}

type CreateTemplateUseCaseResult struct {
	Template       database.Template            `json:"template,omitempty"`
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
	tempType := &database.NullTemplateType{}
	err = tempType.Scan(req.Type.String())
	if err != nil {
		err = shared.WrapWithAppError(err, "invalid template type provided", "")
		log.Error(err.Error())
		return nil, err
	}
	temp, err := qtx.CreateTemplate(ctx, database.CreateTemplateParams{
		Name:   req.Name,
		Type:   *tempType,
		Width:  pgtype.Int4{Int32: int32(req.Width), Valid: true},
		Height: pgtype.Int4{Int32: int32(req.Height), Valid: true},
	})
	if err != nil {
		err = shared.WrapWithAppError(err, "failed to create template", "")
		log.Error(err.Error())
		return nil, err
	}
	dist, err := qtx.CreateTemplateDistortions(ctx, database.CreateTemplateDistortionsParams{
		X:          pgtype.Int4{Int32: int32(req.X), Valid: true},
		Y:          pgtype.Int4{Int32: int32(req.Y), Valid: true},
		TemplateID: temp.ID,
	})
	if err != nil {
		err = shared.WrapWithAppError(err, "failed to create template distortion", "")
		log.Error(err.Error())
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		err = shared.WrapWithAppError(err, "failed to create template distortion", "")
		log.Error(err.Error())
		return nil, err
	}
	return &CreateTemplateUseCaseResult{
		Template:   temp,
		Distortion: dist,
	}, nil
}
