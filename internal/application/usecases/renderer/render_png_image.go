package renderer

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/storage"
	"algvisual/internal/shared"
	"context"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

type RenderPngImageInput struct {
	Layout entities.Layout `json:"layout,omitempty"`
}

type RenderPngImageOutput struct {
	ImagePath string `json:"path_url,omitempty"`
	ImageURL  string `json:"image_url,omitempty"`
}

func RenderPngImageUseCase(
	ctx context.Context,
	req RenderPngImageInput,
	storage storage.FileStorage,
	cfg *config.AppConfig,
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
	log.Debug("renderig Background")
	if c != nil {
		for _, e := range c.Elements {
			log.Debug("loading background image")
			img, err := storage.LoadImageFromURL(e.ImageURL)
			if err != nil {
				return nil, shared.NewInternalErrorWithDetails(
					"failed to load image to render it",
					err.Error(),
				)
			}
			log.Debug("loaded background image")
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
	for _, c := range req.Layout.BackgroundList {
		for _, e := range c.Elements {
			log.Debug("loading background list image")
			img, err := storage.LoadImageFromURL(e.ImageURL)
			if err != nil {
				return nil, err
			}
			log.Debug("finished background list image")
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
	if len(req.Layout.Components) == 0 {
		return nil, errors.New("no component to render, skipping image generation request")
	}
	log.Debug("renderig components")
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
			if e.OuterContainer.Width() < 50 {
				return nil, errors.New("width less than 50 pixels")
			}
			if e.OuterContainer.Height() < 50 {
				return nil, errors.New("height less than 50 pixels")
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
	log.Debug("drawing grid")
	// draw grid
	borderColor := color.RGBA{255, 0, 0, 255} // Red color
	for _, g := range req.Layout.Grid.GetCells() {
		// Draw the rectangle borders
		for x := g.Xi; x < g.Xii; x++ {
			board.Set(int(x), int(g.Yi), borderColor)  // Top border
			board.Set(int(x), int(g.Yii), borderColor) // Bottom border
		}
		for y := g.Yi; y < g.Yii; y++ {
			board.Set(int(g.Xi), int(y), borderColor)  // Left border
			board.Set(int(g.Xii), int(y), borderColor) // Right border
		}
	}

	uniqid, _ := uuid.NewRandom()
	name := uniqid.String()

	log.Debug("saving rendered image")
	imageURL, err := storage.SaveImage(name, board)
	if err != nil {
		return nil, shared.NewInternalErrorWithDetails("failed to save image", err.Error())
	}
	log.Debug("saved rendered image")
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
