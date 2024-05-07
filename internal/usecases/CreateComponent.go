package usecases

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"algvisual/internal/database"
	"algvisual/internal/shared"
)

type CreateComponentRequest struct {
	Type        string  `json:"type,omitempty"`
	ElementsID  []int32 `json:"elements_id,omitempty"`
	Color       string  `json:"color,omitempty"`
	ComponentID string  `json:"component_id,omitempty"`
	PhotoshopID int     `json:"photoshop_id,omitempty" param:"photoshop_id"`
}

type CreateComponentResult struct {
	Data []database.DesignElement `json:"data,omitempty"`
}

func CreateComponentUseCase(
	ctx context.Context,
	req CreateComponentRequest,
	queries *database.Queries,
	db *pgxpool.Pool,
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
	elements, err := qtx.GetdesignElementsByIDlist(ctx, req.ElementsID)
	if err != nil && err != sql.ErrNoRows {
		err = shared.WrapWithAppError(err, "Falha ao buscar elementos do photoshop", err.Error())
		log.Error(err.Error())
		return nil, err
	}
	if len(elements) == 0 {
		return nil, shared.NewAppError(400, "Nenhum elemento encontrado", "")
	}
	xi, yi, xii, yii, width, heigh := calculateBoundaringBoxForComponent(elements)
	comp, err := qtx.CreateComponent(ctx, database.CreateComponentParams{
		DesignID: int32(req.PhotoshopID),
		Width:    pgtype.Int4{Int32: width, Valid: true},
		Height:   pgtype.Int4{Int32: heigh, Valid: true},
		Xi:       pgtype.Int4{Int32: xi, Valid: true},
		Xii:      pgtype.Int4{Int32: xii, Valid: true},
		Yi:       pgtype.Int4{Int32: yi, Valid: true},
		Yii:      pgtype.Int4{Int32: yii, Valid: true},
		Color:    pgtype.Text{String: req.Color, Valid: req.Color != ""},
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
	elUpdated, err := qtx.UpdateManydesignElement(
		ctx,
		database.UpdateManydesignElementParams{
			DesignID:            int32(req.PhotoshopID),
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

func calculateBoundaringBoxForComponent(
	elements []database.DesignElement,
) (int32, int32, int32, int32, int32, int32) {
	xi := elements[0].Xi.Int32
	xii := elements[0].Xii.Int32
	yi := elements[0].Yi.Int32
	yii := elements[0].Yii.Int32
	for _, elem := range elements {
		if elem.Xi.Int32 < xi {
			xi = elem.Xi.Int32
		}
		if elem.Xii.Int32 > xii {
			xii = elem.Xii.Int32
		}
		if elem.Yi.Int32 < yi {
			yi = elem.Yi.Int32
		}
		if elem.Yii.Int32 > yii {
			yii = elem.Yii.Int32
		}
	}
	return xi, yi, xii, yii, xii - xi, yii - yi
}
