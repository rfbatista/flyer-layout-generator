package renderer

import (
	"algvisual/internal/entities"
	"algvisual/internal/infra"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

type RenderPngImageInput struct {
	Layout entities.Layout
}

type RenderPngImageOutput struct {
	ImagePath string `json:"path_url,omitempty"`
	ImageURL  string `json:"image_url,omitempty"`
}

func RenderPngImageUseCase(
	ctx context.Context,
	req RenderPngImageInput,
	storage infra.FileStorage,
	cfg *infra.AppConfig,
	log *zap.Logger,
	textDrawer *TextDrawer,
) (*RenderPngImageOutput, error) {
	board := image.NewRGBA(image.Rect(0, 0, int(req.Layout.Width), int(req.Layout.Height)))
	log.Debug("image rendering summary",
		zap.Int("number of components", len(req.Layout.Components)),
		zap.Int("layout witdth", int(req.Layout.Width)),
		zap.Int("layout height", int(req.Layout.Height)),
	)
	c := req.Layout.Background
	for _, e := range c.Elements {
		img, err := storage.LoadImageFromURL(e.ImageURL)
		if err != nil {
			return nil, err
		}
		nimg := resize.Resize(
			uint(e.OuterContainer.Width()),
			uint(e.OuterContainer.Height()),
			img,
			resize.Lanczos2,
		)
		bounds := e.OuterContainer.Rect()
		pos := image.Point{}
		draw.Draw(
			board,
			bounds,
			nimg,
			pos,
			draw.Over,
		)
	}
	for _, c := range req.Layout.BackgroundList {
		for _, e := range c.Elements {
			img, err := storage.LoadImageFromURL(e.ImageURL)
			if err != nil {
				return nil, err
			}
			nimg := resize.Resize(
				uint(e.OuterContainer.Width()),
				uint(e.OuterContainer.Height()),
				img,
				resize.Lanczos2,
			)
			bounds := e.OuterContainer.Rect()
			pos := image.Point{}
			draw.Draw(
				board,
				bounds,
				nimg,
				pos,
				draw.Over,
			)
		}
	}
	for _, c := range req.Layout.Components {
		for _, e := range c.Elements {
			if e.Kind == "typess" {
				text := e.PickTextFromProperty()
				size := textDrawer.FindTextSizeToFillContainer(text, e.OuterContainer)
				textDrawer.addLabel(
					board,
					e.OuterContainer.Position().X,
					e.OuterContainer.Position().Y,
					text,
					int32(size),
				)
				DrawContainer(e.OuterContainer, board)
				continue
			}
			img, err := storage.LoadImageFromURL(e.ImageURL)
			if err != nil {
				return nil, err
			}
			nimg := resize.Resize(
				uint(e.OuterContainer.Width()),
				uint(e.OuterContainer.Height()),
				img,
				resize.Lanczos2,
			)
			bounds := e.OuterContainer.Rect()
			pos := image.Point{}
			draw.Draw(
				board,
				bounds,
				nimg,
				pos,
				draw.Over,
			)
		}
	}

	// draw grid
	// borderColor := color.RGBA{255, 0, 0, 255} // Red color
	// for _, g := range req.Layout.Grid.GetCells() {
	// 	// Draw the rectangle borders
	// 	for x := g.Xi; x < g.Xii; x++ {
	// 		board.Set(int(x), int(g.Yi), borderColor)  // Top border
	// 		board.Set(int(x), int(g.Yii), borderColor) // Bottom border
	// 	}
	// 	for y := g.Yi; y < g.Yii; y++ {
	// 		board.Set(int(g.Xi), int(y), borderColor)  // Left border
	// 		board.Set(int(g.Xii), int(y), borderColor) // Right border
	// 	}
	// }

	uniqid, _ := uuid.NewRandom()
	name := uniqid.String()
	imageURL, err := storage.SaveImage(name, board)
	if err != nil {
		return nil, err
	}
	return &RenderPngImageOutput{
		ImagePath: imageURL,
		ImageURL:  fmt.Sprintf("/api/v1/images/%s.png", name),
	}, nil
}

func DrawContainer(g entities.Container, board *image.RGBA) {
	borderColor := color.RGBA{255, 0, 0, 255} // Red color
	// Draw the rectangle borders
	for x := g.UpperLeft.X; x < g.DownRight.X; x++ {
		board.Set(int(x), int(g.UpperLeft.Y), borderColor) // Top border
		board.Set(int(x), int(g.DownRight.Y), borderColor) // Bottom border
	}
	for y := g.UpperLeft.Y; y < g.DownRight.Y; y++ {
		board.Set(int(g.UpperLeft.X), int(y), borderColor) // Left border
		board.Set(int(g.DownRight.X), int(y), borderColor) // Right border
	}
}
