package usecases

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/shared"
)

type CreateComponentRequest struct {
	Type        string  `json:"type,omitempty"`
	ElementsID  []int32 `json:"elements_id,omitempty"`
	Color       string  `json:"color,omitempty"`
	ComponentID string  `json:"component_id,omitempty"`
	PhotoshopID int     `json:"photoshop_id,omitempty"`
}

type CreateComponentResult struct {
	Data []database.PhotoshopElement `json:"data,omitempty"`
}

func CreateComponentUseCase(
	ctx context.Context,
	req CreateComponentRequest,
	queries *database.Queries,
	db *pgx.Conn,
	log *zap.Logger,
) (*CreateComponentResult, error) {
	tx, err := db.Begin(ctx)
	if err != nil {
		err = shared.WrapWithAppError(err, "cant start transaction", "")
		log.Error(err.Error())
		return nil, err
	}
	defer tx.Rollback(ctx)
	qtx := queries.WithTx(tx)
	comp, err := qtx.CreateComponent(ctx, database.CreateComponentParams{
		PhotoshopID: int32(req.PhotoshopID),
		Width:       pgtype.Int4{Int32: 0, Valid: true},
		Height:      pgtype.Int4{Int32: 0, Valid: true},
		Type: database.NullComponentType{
			ComponentType: database.ComponentType(req.Type),
			Valid:         true,
		},
	})
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha na criação do componente", "")
		log.Error(err.Error())
		return nil, err
	}
	elUpdated, err := queries.UpdateManyPhotoshopElement(
		ctx,
		database.UpdateManyPhotoshopElementParams{
			PhotoshopID:         int32(req.PhotoshopID),
			ComponentIDDoUpdate: true,
			ComponentID:         comp.ID,
			Ids:                 req.ElementsID,
		},
	)
	if err != nil {
		return nil, shared.WrapWithAppError(err, "Falha ao atualizar elementos", "")
	}
	tx.Commit(ctx)
	return &CreateComponentResult{
		Data: elUpdated,
	}, nil
}
