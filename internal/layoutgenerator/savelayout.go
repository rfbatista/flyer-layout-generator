package layoutgenerator

import (
	"algvisual/internal/database"
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
) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	qtx := queries.WithTx(tx)
	layoutCreated, err := qtx.CreateLayout(ctx, database.CreateLayoutParams{
		Width:    pgtype.Int4{Int32: l.Width, Valid: true},
		Height:   pgtype.Int4{Int32: l.Height, Valid: true},
		DesignID: pgtype.Int4{Int32: l.DesignID, Valid: true},
	})
	if err != nil {
		return err
	}
	for _, c := range l.Components {
		comp := mapper.LayoutComponentFromDomain(c)
		_, err = qtx.CreateLayoutComponent(ctx, database.CreateLayoutComponentParams{
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
		})
		if err != nil {
			return err
		}
	}
	for _, region := range l.Grid.Cells {
		e := mapper.LayoutRegionFromDomain(region)
		_, err = qtx.CreateLayoutRegion(ctx, database.CreateLayoutRegionParams{
			LayoutID: int32(layoutCreated.ID),
			Xi:       e.Xi,
			Xii:      e.Xii,
			Yi:       e.Yi,
			Yii:      e.Yii,
		})
		if err != nil {
			return err
		}
	}
	temp := mapper.LayoutTemplateFromDomain(l.Template)
	_, err = qtx.CreateLayoutTemplate(ctx, database.CreateLayoutTemplateParams{
		LayoutID: int32(layoutCreated.ID),
		Type:     temp.Type,
		Width:    temp.Width,
		Height:   temp.Height,
	})
	if err != nil {
		return err
	}
	return nil
}
