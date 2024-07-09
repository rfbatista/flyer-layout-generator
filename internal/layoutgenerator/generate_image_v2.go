package layoutgenerator

import (
	"algvisual/database"
	"algvisual/internal/designassets"
	"algvisual/internal/entities"
	"algvisual/internal/grammars"
	"algvisual/internal/renderer"
	"algvisual/internal/templates"
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type GenerateImageV2Input struct {
	LayoutID              int32 `param:"layout_id"   json:"layout_id,omitempty"`
	RequestID             int32
	TemplateID            int32          `param:"template_id" json:"template_id,omitempty"`
	LimitSizerPerElement  bool           `                    json:"limit_sizer_per_element,omitempty"`
	AnchorElements        bool           `                    json:"anchor_elements,omitempty"`
	ShowGrid              bool           `                    json:"show_grid,omitempty"               form:"show_grid"`
	MinimiumComponentSize int32          `                    json:"minimium_component_size,omitempty"`
	MinimiumTextSize      int32          `                    json:"minimium_text_size,omitempty"`
	SlotsX                int32          `                    json:"slots_x,omitempty"                 form:"grid_x"`
	SlotsY                int32          `                    json:"slots_y,omitempty"                 form:"grid_y"`
	Padding               int32          `                    json:"padding,omitempty"                 form:"padding"`
	KeepProportions       bool           `                    json:"keep_proportions,omitempty"`
	Priorities            []string       `                    json:"priorities,omitempty"              form:"priority[]"`
	LayoutPriorities      map[string]int `                    json:"priorities,omitempty"`
}

type GenerateImageV2Output struct {
	ImageURL string `json:"image_url,omitempty"`
	Layout   entities.Layout
}

func GenerateImageV2UseCase(
	ctx context.Context,
	req GenerateImageV2Input,
	db *database.Queries,
	tservice templates.TemplatesService,
	log *zap.Logger,
	render renderer.RendererService,
	pool *pgxpool.Pool,
) (*GenerateImageV2Output, error) {
	out, err := GetLayoutByIDUseCase(ctx, db, GetLayoutByIDRequest{
		LayoutID: req.LayoutID,
	})
	if err != nil {
		log.Error("failed to get layout by id", zap.Error(err))
		return nil, err
	}
	layout, temp, err := getLayoutToRun(ctx, GetLayoutInput{
		TemplateID: req.TemplateID,
		LayoutID:   req.LayoutID,
		DesignID:   out.Layout.DesignID,
		Padding:    req.Padding,
		Priorities: req.Priorities,
	}, db, log, out.Layout.Width, out.Layout.Height)
	if err != nil {
		log.Error("failed to get layout to run", zap.Error(err))
		return nil, err
	}
	layout.Config.Priorities = req.LayoutPriorities
	newLayout, err := grammars.RunV2(*layout, *temp, req.SlotsX, req.SlotsY, log)
	if err != nil {
		log.Error("failed to generate new layout", zap.Error(err))
		return nil, err
	}
	log.Debug("summary of layout output",
		zap.Int("total of components", len(newLayout.Components)),
	)
	assets, err := designassets.GetDesignAssetsByDesignIdUseCase(
		ctx,
		designassets.GetDesignAssetsByDesignIdInput{
			DesignID: out.Layout.DesignID,
		},
		db,
	)
	if err != nil {
		log.Error("failed to render new layout", zap.Error(err))
		return nil, err
	}
	newLayout.DesignAssets = assets.Data
	var elements []entities.LayoutElement
	for _, c := range newLayout.Components {
		elements = append(elements, c.Elements...)
	}
	newLayout.Elements = elements
	checkResult, err := CheckLayoutSimilaritiesUseCase(ctx, CheckLayoutSimilaritiesInput{
		RequestID: req.RequestID,
		Layout:    *newLayout,
	}, db)
	if err != nil {
		log.Error("failed to check new layout similarity", zap.Error(err))
		return nil, err
	}
	if checkResult.HaveSimilar {
		return nil, errors.New("similar layout was found")
	}
	imageResult, err := render.RenderPNGImage(ctx, renderer.RenderPngImageInput{Layout: *newLayout})
	if err != nil {
		log.Error("failed to render new layout", zap.Error(err))
		return nil, err
	}
	newLayout.ImageURL = imageResult.ImageURL
	newLayout.RequestID = req.RequestID
	layoutCreated, err := SaveLayout(ctx, *newLayout, db, pool)
	if err != nil {
		log.Error("failed to save new layout", zap.Error(err))
		return nil, err
	}
	return &GenerateImageV2Output{
		ImageURL: imageResult.ImageURL,
		Layout:   *layoutCreated,
	}, nil
}
