package layoutgenerator

import (
	"algvisual/database"
	"algvisual/internal/entities"
	"algvisual/internal/renderer"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type UpdateLayoutElementPositionInput struct {
	ID       int32          `param:"element_id" json:"id,omitempty"`
	Position entities.Point `                   json:"position,omitempty"`
}

type UpdateLayoutElementPositionOutput struct {
	Data entities.LayoutElement
}

func UpdateLayoutElementPositionUseCase(
	ctx context.Context,
	req UpdateLayoutElementPositionInput,
	db *database.Queries,
	render renderer.RendererService,
) (*UpdateLayoutElementPositionOutput, error) {
	element, err := GetLayoutElementByIdUseCase(ctx, GetLayoutElementByIdInput{ID: req.ID}, db)
	if err != nil {
		return nil, err
	}
	element.Data.MoveOnOuter(req.Position)
	_, err = db.UpdateLayoutElementPosition(ctx, database.UpdateLayoutElementPositionParams{
		ID:       element.Data.ID,
		Xi:       pgtype.Int4{Int32: int32(element.Data.OuterContainer.UpperLeft.X), Valid: true},
		Xii:      pgtype.Int4{Int32: int32(element.Data.OuterContainer.DownRight.X), Valid: true},
		Yi:       pgtype.Int4{Int32: int32(element.Data.OuterContainer.UpperLeft.Y), Valid: true},
		Yii:      pgtype.Int4{Int32: int32(element.Data.OuterContainer.DownRight.Y), Valid: true},
		InnerXi:  pgtype.Int4{Int32: int32(element.Data.InnerXi), Valid: true},
		InnerXii: pgtype.Int4{Int32: int32(element.Data.InnerXii), Valid: true},
		InnerYi:  pgtype.Int4{Int32: int32(element.Data.InnerYi), Valid: true},
		InnerYii: pgtype.Int4{Int32: int32(element.Data.InnerYii), Valid: true},
	})
	if err != nil {
		return nil, err
	}
	layout, err := GetLayoutByIDUseCase(ctx, db, GetLayoutByIDRequest{
		LayoutID: element.Data.LayoutID,
	})
	if err != nil {
		return nil, err
	}
	out, err := render.RenderPNGImage(ctx, renderer.RenderPngImageInput{
		Layout: layout.Layout,
	})
	if err != nil {
		return nil, err
	}
	err = db.UpdateLayoutImagByID(ctx, database.UpdateLayoutImagByIDParams{
		ID:       int64(layout.Layout.ID),
		ImageUrl: pgtype.Text{String: out.ImageURL, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &UpdateLayoutElementPositionOutput{Data: element.Data}, nil
}
