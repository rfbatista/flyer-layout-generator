package layoutgenerator

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/mapper"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SaveLayout(
	ctx context.Context,
	l entities.Layout,
	queries *database.Queries,
	db *pgxpool.Pool,
) (*entities.Layout, error) {
	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	qtx := queries.WithTx(tx)
	layoutCreated, err := qtx.CreateLayout(ctx, database.CreateLayoutParams{
		Width:      pgtype.Int4{Int32: l.Width, Valid: true},
		Height:     pgtype.Int4{Int32: l.Height, Valid: true},
		DesignID:   pgtype.Int4{Int32: l.DesignID, Valid: true},
		IsOriginal: pgtype.Bool{},
		ImageUrl:   pgtype.Text{String: l.ImageURL, Valid: true},
		Stages:     l.Stages,
	})
	if err != nil {
		return nil, err
	}
	ld := mapper.LayoutToDomain(layoutCreated)
	for _, c := range l.Components {
		comp := mapper.LayoutComponentFromDomain(c)
		componentCreated, err := qtx.CreateLayoutComponent(
			ctx,
			database.CreateLayoutComponentParams{
				LayoutID: int32(layoutCreated.ID),
				DesignID: l.DesignID,
				Xi:       comp.Xi,
				Xii:      comp.Xii,
				Yi:       comp.Yi,
				Yii:      comp.Yii,
				Type:     comp.Type,
				Color:    comp.Color,
				BboxXi:   comp.BboxXi,
				BboxYi:   comp.BboxYi,
				BboxYii:  comp.BboxYii,
				BboxXii:  comp.BboxXii,
			},
		)
		if err != nil {
			return nil, err
		}
		for _, i := range c.Elements {
			ele, err := qtx.CreateElement(
				ctx,
				mapper.LayoutElementToCreateElement(
					i,
					int32(layoutCreated.ID),
					l.DesignID,
					componentCreated.ID,
				),
			)
			if err != nil {
				return nil, err
			}
			dele := mapper.ToDesignElementEntitie(ele)
			c.Elements = append(c.Elements, dele)
		}
		ld.Components = append(ld.Components, c)
	}
	tx.Commit(ctx)
	return &ld, nil
}
