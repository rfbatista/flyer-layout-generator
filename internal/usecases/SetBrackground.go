package usecases

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/shared"
)

type SetBackgroundUseCaseRequest struct {
	PhotoshopID int32   `params:"PhotoshopID" json:"photoshop_id,omitempty"`
	Elements    []int32 `                     json:"elements,omitempty"     body:"elements"`
}

type SetBackgroundUseCaseResult struct {
	Data []database.PhotoshopElement
}

func SetBackgroundUseCase(
	ctx context.Context,
	queries *database.Queries,
	db *pgx.Conn,
	req SetBackgroundUseCaseRequest,
	log *zap.Logger,
) (*SetBackgroundUseCaseResult, error) {
	tx, err := db.Begin(ctx)
	if err != nil {
		err = shared.WrapWithAppError(err, "cant start transaction", "")
		log.Error(err.Error())
		return nil, err
	}
	defer tx.Rollback(ctx)
	qtx := queries.WithTx(tx)
	comp, err := qtx.GetPhotoshopBackgroundComponent(ctx, int32(req.PhotoshopID))
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao procurar plano de fundo existente", "")
		log.Error(err.Error())
		return nil, err
	}
	if comp.ID == 0 {
		comp, err = qtx.CreateComponent(ctx, database.CreateComponentParams{
			PhotoshopID: req.PhotoshopID,
			Width:       pgtype.Int4{Int32: 0, Valid: true},
			Height:      pgtype.Int4{Int32: 0, Valid: true},
			Type: database.NullComponentType{
				ComponentType: database.ComponentTypeBackground,
				Valid:         true,
			},
		})
		if err != nil {
			err = shared.WrapWithAppError(err, "Falha ao criar plano de fundo", "")
			log.Error(err.Error())
			return nil, err
		}
	}
	elUpdated, err := qtx.UpdateManyPhotoshopElement(
		ctx,
		database.UpdateManyPhotoshopElementParams{
			PhotoshopID:         req.PhotoshopID,
			ComponentIDDoUpdate: true,
			ComponentID:         comp.ID,
			Ids:                 req.Elements,
		},
	)
	if err != nil {
		return nil, shared.WrapWithAppError(err, "Falha ao criar atualizar elementos do plano de fundo", "")
	}
	return &SetBackgroundUseCaseResult{
		Data: elUpdated,
	}, nil
}
