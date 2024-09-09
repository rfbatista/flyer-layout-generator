package layoutgenerator

import (
	"algvisual/internal/application/usecases/designassets"
	"algvisual/internal/application/usecases/grammars"
	"algvisual/internal/application/usecases/renderer"
	"algvisual/internal/application/usecases/templates"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/database"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type GenerateLayoutUseCase struct {
	db       *database.Queries
	tservice templates.TemplatesService
	log      *zap.Logger
	render   renderer.RendererService
	pool     *pgxpool.Pool
	das      *designassets.DesignAssetService
	cfg      config.AppConfig
}

func NewGenerateLayoutUseCase(
	db *database.Queries,
	tservice templates.TemplatesService,
	log *zap.Logger,
	render renderer.RendererService,
	pool *pgxpool.Pool,
	das *designassets.DesignAssetService,
	cfg config.AppConfig,
) (*GenerateLayoutUseCase, error) {
	return &GenerateLayoutUseCase{
		db:       db,
		tservice: tservice,
		log:      log,
		render:   render,
		pool:     pool,
		das:      das,
		cfg:      cfg,
	}, nil
}

type GenerateImageV2Input struct {
	LayoutID              int32          `param:"layout_id"`
	RequestID             int32          `                    json:"request_id,omitempty"`
	TemplateID            int32          `param:"template_id"`
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
	ImageURL string          `json:"image_url,omitempty"`
	Layout   entities.Layout `json:"layout,omitempty"`
}

func (g GenerateLayoutUseCase) Execute(
	ctx context.Context,
	req GenerateImageV2Input,
) (*GenerateImageV2Output, error) {
	out, err := GetLayoutByIDUseCase(ctx, g.db, GetLayoutByIDRequest{
		LayoutID: req.LayoutID,
	}, g.das)
	if err != nil {
		g.log.Error("failed to get layout by id", zap.Error(err))
		return nil, err
	}
	layout, temp, err := getLayoutToRun(ctx, GetLayoutInput{
		TemplateID: req.TemplateID,
		LayoutID:   req.LayoutID,
		DesignID:   out.Layout.DesignID,
		Padding:    60,
		Priorities: req.Priorities,
	}, g.db, g.log, out.Layout.Width, out.Layout.Height, g.das)
	if err != nil {
		g.log.Error("failed to get layout to run", zap.Error(err))
		return nil, err
	}
	layout.Config.Priorities = req.LayoutPriorities
	newLayout, err := grammars.RunV2(*layout, *temp, req.SlotsX, req.SlotsY, g.log)
	if err != nil {
		g.log.Error("failed to generate new layout", zap.Error(err))
		return nil, err
	}
	g.log.Debug("summary of layout output",
		zap.Int("total of components", len(newLayout.Components)),
	)
	assets, err := g.das.GetDesignAssetByDesignID(
		ctx,
		designassets.GetDesignAssetsByDesignIdInput{
			DesignID: out.Layout.DesignID,
		},
	)
	if err != nil {
		g.log.Error("failed to render new layout", zap.Error(err))
		return nil, err
	}
	newLayout.DesignAssets = assets.Data
	var elements []entities.LayoutElement
	for _, c := range newLayout.Components {
		elements = append(elements, c.Elements...)
	}
	newLayout.Elements = elements
	// checkResult, err := CheckLayoutSimilaritiesUseCase(ctx, CheckLayoutSimilaritiesInput{
	// 	RequestID: req.RequestID,
	// 	Layout:    *newLayout,
	// }, g.db, g.das)
	// if err != nil {
	// 	g.log.Error("failed to check new layout similarity", zap.Error(err))
	// 	return nil, err
	// }
	// if checkResult.HaveSimilar {
	// 	return &GenerateImageV2Output{}, nil
	// }
	// shared.WriteDataToFileAsJSON(req, fmt.Sprintf("%s/replications.json", g.cfg.PhotoshopFilesPath))
	// shared.WriteDataToFileAsJSON(*newLayout, fmt.Sprintf("%s/replications.json", g.cfg.PhotoshopFilesPath))
	g.log.Debug("starting execution of png image render")
	imageResult, err := g.render.RenderPNGImage(
		ctx,
		renderer.RenderPngImageInput{Layout: *newLayout},
	)
	g.log.Debug("finished execution of png image render")
	if err != nil {
		g.log.Error("failed to render new layout", zap.Error(err))
		return nil, err
	}
	newLayout.ImageURL = imageResult.ImageURL
	newLayout.RequestID = req.RequestID
	newLayout.TemplateID = req.TemplateID
	layoutCreated, err := SaveLayout(ctx, *newLayout, g.db, g.pool)
	if err != nil {
		g.log.Error("failed to save new layout", zap.Error(err))
		return nil, err
	}
	return &GenerateImageV2Output{
		ImageURL: imageResult.ImageURL,
		Layout:   *layoutCreated,
	}, nil
}
