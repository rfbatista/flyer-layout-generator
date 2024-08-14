package usecase

import (
	"algvisual/database"
	"algvisual/internal/shared"
	"database/sql"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type SetBackgroundUseCaseRequest struct {
	DesignID int32   `param:"design_id" json:"photoshop_id,omitempty"`
	Elements []int32 `                  json:"elements,omitempty"     forms:"elements" body:"elements"`
}

type SetBackgroundUseCaseResult struct {
	Data []database.LayoutElement
}

func SetBackgroundUseCase(
	c echo.Context,
	queries *database.Queries,
	db *pgxpool.Pool,
	req SetBackgroundUseCaseRequest,
	log *zap.Logger,
) (*SetBackgroundUseCaseResult, error) {
	ctx := c.Request().Context()
	tx, err := db.Begin(ctx)
	if err != nil {
		err = shared.WrapWithAppError(err, "cant start transaction", "")
		log.Error(err.Error())
		return nil, err
	}
	defer tx.Rollback(ctx)
	qtx := queries.WithTx(tx)
	comp, err := qtx.GetdesignBackgroundComponent(ctx, int32(req.DesignID))
	if err != nil && err == sql.ErrNoRows {
		err = shared.WrapWithAppError(err, "Falha ao procurar plano de fundo existente", "")
		log.Error(err.Error())
		return nil, err
	}
	if comp.ID == 0 {
		comp, err = qtx.CreateComponent(ctx, database.CreateComponentParams{
			DesignID: req.DesignID,
			Width:    pgtype.Int4{Int32: 0, Valid: true},
			Height:   pgtype.Int4{Int32: 0, Valid: true},
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
	elUpdated, err := qtx.UpdateManydesignElement(
		ctx,
		database.UpdateManydesignElementParams{
			DesignID:            req.DesignID,
			ComponentIDDoUpdate: true,
			ComponentID:         comp.ID,
			Ids:                 req.Elements,
		},
	)
	if err != nil {
		return nil, shared.WrapWithAppError(
			err,
			"Falha ao criar atualizar elementos do plano de fundo",
			"",
		)
	}
	return &SetBackgroundUseCaseResult{
		Data: elUpdated,
	}, nil
}
