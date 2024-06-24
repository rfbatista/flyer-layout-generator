package designs

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/shared"
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type CreateComponentRequest struct {
	Type       string  `form:"type"       json:"type,omitempty"`
	ElementsID []int32 `form:"elements[]" json:"elements_id,omitempty"`
	Color      string  `                  json:"color,omitempty"`
	DesignID   int     `                  json:"photoshop_id,omitempty" param:"design_id"`
	LayoutID   int32   `                  json:"layout_id,omitempty"    param:"layout_id"`
}

type CreateComponentResult struct {
	Data []database.LayoutElement `json:"data,omitempty"`
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
	des, err := qtx.Getdesign(ctx, int32(req.DesignID))
	if err != nil && err != sql.ErrNoRows {
		err = shared.WrapWithAppError(err, "Falha ao buscar arquivo do design", err.Error())
		log.Error(err.Error())
		return nil, err
	}
	outer, inner := calculateContainersForComponent(elements)
	ctype := entities.StringToComponentType(req.Type)
	dtype, err := entities.ComponentTypeToDatabaseComponentType(ctype)
	if err != nil {
		log.Error("tipo de component invalido", zap.Error(err))
		return nil, err
	}
	comp, err := qtx.CreateComponent(ctx, database.CreateComponentParams{
		DesignID: int32(req.DesignID),
		LayoutID: req.LayoutID,
		Width:    pgtype.Int4{Int32: inner.Width(), Valid: true},
		Height:   pgtype.Int4{Int32: inner.Height(), Valid: true},
		Xi:       pgtype.Int4{Int32: outer.UpperLeft.X, Valid: true},
		Xii:      pgtype.Int4{Int32: outer.DownRight.X, Valid: true},
		Yi:       pgtype.Int4{Int32: outer.UpperLeft.Y, Valid: true},
		Yii:      pgtype.Int4{Int32: outer.DownRight.Y, Valid: true},
		InnerXi:  pgtype.Int4{Int32: inner.UpperLeft.X, Valid: true},
		InnerXii: pgtype.Int4{Int32: inner.DownRight.X, Valid: true},
		InnerYi:  pgtype.Int4{Int32: inner.UpperLeft.Y, Valid: true},
		InnerYii: pgtype.Int4{Int32: inner.DownRight.Y, Valid: true},
		BboxXi:   pgtype.Int4{Int32: 0, Valid: true},
		BboxYi:   pgtype.Int4{Int32: 0, Valid: true},
		BboxXii:  pgtype.Int4{Int32: des.Width.Int32, Valid: true},
		BboxYii:  pgtype.Int4{Int32: des.Height.Int32, Valid: true},
		Color:    pgtype.Text{String: req.Color, Valid: req.Color != ""},
		Type: database.NullComponentType{
			ComponentType: database.ComponentType(dtype),
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
			DesignID:            int32(req.DesignID),
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
	elements []database.LayoutElement,
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

func calculateContainersForComponent(
	elements []database.LayoutElement,
) (entities.Container, entities.Container) {
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
	outer := entities.NewContainer(entities.NewPoint(xi, yi), entities.NewPoint(xii, yii))

	inxi := elements[0].InnerXi.Int32
	inxii := elements[0].InnerXii.Int32
	inyi := elements[0].InnerYi.Int32
	inyii := elements[0].InnerYii.Int32
	for _, elem := range elements {
		if elem.Xi.Int32 < xi {
			inxi = elem.Xi.Int32
		}
		if elem.Xii.Int32 > xii {
			inxii = elem.Xii.Int32
		}
		if elem.Yi.Int32 < yi {
			inyi = elem.Yi.Int32
		}
		if elem.Yii.Int32 > yii {
			inyii = elem.Yii.Int32
		}
	}
	return outer, entities.NewContainer(
		entities.NewPoint(inxi, inyi),
		entities.NewPoint(inxii, inyii),
	)
}
