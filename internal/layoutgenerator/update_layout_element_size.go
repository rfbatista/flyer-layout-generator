package layoutgenerator

import (
	"algvisual/database"
	"algvisual/internal/designassets"
	"algvisual/internal/entities"
	"algvisual/internal/renderer"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type UpdateLayoutElementSizeInput struct {
	ID     int32 `param:"element_id" json:"id,omitempty"`
	Width  int32 `                   json:"width,omitempty"`
	Height int32 `                   json:"height,omitempty"`
}

type UpdateLayoutElementSizeOutput struct {
	Data entities.LayoutElement `json:"data,omitempty"`
}

func UpdateLayoutElementSizeUseCase(
	ctx context.Context,
	req UpdateLayoutElementSizeInput,
	db *database.Queries,
	render renderer.RendererService,
	das *designassets.DesignAssetService,
) (*UpdateLayoutElementSizeOutput, error) {
	element, err := GetLayoutElementByIdUseCase(ctx, GetLayoutElementByIdInput{
		ID: req.ID,
	}, db)
	if err != nil {
		return nil, err
	}
	w := element.Data.Width()
	scale := float64(req.Width) / float64(w)
	element.Data.ScaleFix(scale)
	_, err = db.UpdateLayoutElementSize(ctx, database.UpdateLayoutElementSizeParams{
		ID:       element.Data.ID,
		Xi:       pgtype.Int4{Int32: int32(element.Data.OuterContainer.UpperLeft.X), Valid: true},
		Xii:      pgtype.Int4{Int32: int32(element.Data.OuterContainer.DownRight.X), Valid: true},
		Yi:       pgtype.Int4{Int32: int32(element.Data.OuterContainer.UpperLeft.Y), Valid: true},
		Yii:      pgtype.Int4{Int32: int32(element.Data.OuterContainer.DownRight.Y), Valid: true},
		InnerXi:  pgtype.Int4{Int32: int32(element.Data.InnerContainer.UpperLeft.X), Valid: true},
		InnerXii: pgtype.Int4{Int32: int32(element.Data.InnerContainer.DownRight.X), Valid: true},
		InnerYi:  pgtype.Int4{Int32: int32(element.Data.InnerContainer.UpperLeft.Y), Valid: true},
		InnerYii: pgtype.Int4{Int32: int32(element.Data.InnerContainer.DownRight.Y), Valid: true},
		Width:    pgtype.Int4{Int32: int32(req.Width), Valid: true},
		Height:   pgtype.Int4{Int32: int32(req.Height), Valid: true},
	})
	if err != nil {
		return nil, err
	}
	layout, err := GetLayoutByIDUseCase(ctx, db, GetLayoutByIDRequest{
		LayoutID: element.Data.LayoutID,
	}, das)
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
	return &UpdateLayoutElementSizeOutput{
		Data: element.Data,
	}, nil
}
