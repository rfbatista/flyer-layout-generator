package layoutgenerator

import (
	"algvisual/internal/application/usecases/designassets"
	"algvisual/internal/application/usecases/grammars"
	"algvisual/internal/application/usecases/renderer"
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories/mapper"
	"algvisual/internal/shared"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type GenerateImage struct {
	DesignID              int32    `form:"design_id"   json:"photoshop_id,omitempty"            param:"design_id"`
	LayoutID              int32    `form:"layout_id"   json:"layout_id,omitempty"               param:"layout_id"`
	TemplateID            int32    `form:"template_id" json:"template_id,omitempty"             param:"template_id"`
	LimitSizerPerElement  bool     `                   json:"limit_sizer_per_element,omitempty" param:"limit_sizer_per_element"`
	AnchorElements        bool     `                   json:"anchor_elements,omitempty"         param:"anchor_elements"`
	ShowGrid              bool     `form:"show_grid"   json:"show_grid,omitempty"               param:"show_grid"`
	MinimiumComponentSize int32    `                   json:"minimium_component_size,omitempty" param:"minimium_component_size"`
	MinimiumTextSize      int32    `                   json:"minimium_text_size,omitempty"      param:"minimium_text_size"`
	SlotsX                int32    `form:"grid_x"      json:"slots_x,omitempty"                 param:"slots_x"`
	SlotsY                int32    `form:"grid_y"      json:"slots_y,omitempty"                 param:"slots_y"`
	Padding               int32    `form:"padding"     json:"padding,omitempty"                 param:"padding"`
	KeepProportions       bool     `                   json:"keep_proportions,omitempty"        param:"keep_proportions"`
	Priorities            []string `form:"priority[]"  json:"priorities,omitempty"              param:"priorities"`
}

type GenerateImageOutput struct {
	Data       *GenerateImageResultV2 `json:"data,omitempty"`
	TwistedURL string
	Layout     *entities.Layout
}

type GetLayoutInput struct {
	TemplateID int32
	LayoutID   int32
	DesignID   int32
	Padding    int32    `form:"padding"    json:"padding,omitempty"    param:"padding"`
	Priorities []string `form:"priority[]" json:"priorities,omitempty" param:"priorities"`
}

func getLayoutToRun(
	ctx context.Context,
	req GetLayoutInput,
	queries *database.Queries,
	log *zap.Logger,
	width int32,
	height int32,
	das *designassets.DesignAssetService,
) (*entities.Layout, *entities.Template, error) {
	template, err := queries.GetTemplate(ctx, req.TemplateID)
	if err != nil {
		err = shared.WrapWithAppError(
			err,
			fmt.Sprintf("Não foi possivel encontrar o template %d", req.TemplateID),
			"",
		)
		return nil, nil, err
	}
	etemplate := mapper.TemplateToDomain(template.Template)
	elements, err := queries.GetElementsByLayoutID(ctx, req.LayoutID)
	if err != nil {
		err = shared.WrapWithAppError(
			err,
			"Não foi possivel encontrar os elementos do arquivo Photoshop",
			err.Error(),
		)
		return nil, nil, err
	}
	var eelements []entities.LayoutElement
	for _, el := range elements {
		assets, err := das.GetDesignAssetByID(
			ctx,
			designassets.GetDesignAssetByIdInput{ID: el.AssetID},
		)
		if err != nil {
			return nil, nil, err
		}
		element := mapper.ToDesignElementEntitie(el)
		element.Properties = append(element.Properties, assets.Data.Properties...)
		eelements = append(eelements, element)
	}
	compHash := make(map[int32][]entities.LayoutElement)
	for _, c := range eelements {
		if c.ComponentID != 0 {
			compHash[c.ComponentID] = append(compHash[c.ComponentID], c)
		}
	}
	var components []entities.LayoutComponent
	var bg *entities.LayoutComponent
	for k := range compHash {
		data, compErr := queries.GetComponentByID(ctx, k)
		if compErr != nil {
			compErr = shared.WrapWithAppError(
				compErr,
				"Não foi possivel encontrar os componentes do arquivo Photoshop",
				"",
			)
			return nil, nil, compErr
		}
		comp := mapper.TodesignComponentEntitie(data)
		comp.Elements = compHash[k]
		if comp.IsBackground() {
			bg = &comp
		} else {
			components = append(components, comp)
		}
	}
	if len(components) == 0 {
		return nil, nil, shared.NewAppError(
			400,
			"nenhum componente definido para o design escolhido",
			"",
		)
	}
	prancheta := entities.Layout{
		DesignID:   req.DesignID,
		Width:      width,
		Height:     height,
		Template:   etemplate,
		Background: bg,
		Components: components,
		Config: entities.LayoutRequestConfig{
			Padding:    req.Padding,
			ShowGrid:   true,
			Priorities: entities.ListToPrioritiesMap(req.Priorities),
		},
	}
	return &prancheta, &etemplate, nil
}

func GenerateImageUseCase(
	ctx context.Context,
	req GenerateImage,
	queries *database.Queries,
	db *pgxpool.Pool,
	config config.AppConfig,
	log *zap.Logger,
	render renderer.RendererService,
	das *designassets.DesignAssetService,
) (*GenerateImageOutput, error) {
	designFile, err := queries.Getdesign(ctx, req.DesignID)
	if err != nil {
		err = shared.WrapWithAppError(err, "Não foi possivel encontrar o photoshop", "")
		return nil, err
	}
	prancheta, etemplate, err := getLayoutToRun(
		ctx,
		GetLayoutInput{
			DesignID:   req.DesignID,
			TemplateID: req.TemplateID,
			LayoutID:   req.LayoutID,
			Padding:    req.Padding,
			Priorities: req.Priorities,
		},
		queries,
		log,
		designFile.Width.Int32,
		designFile.Height.Int32,
		das,
	)
	if err != nil {
		log.Error("failed to get layout to run", zap.Error(err))
		return nil, err
	}
	nprancheta, _ := grammars.RunV2(*prancheta, *etemplate, req.SlotsX, req.SlotsY, log)
	if !req.ShowGrid {
		nprancheta.Grid = entities.Grid{}
	}
	// res, err := GenerateImageFromPranchetaV2(GenerateImageRequestV2{
	// 	DesignFile: designFile.FileUrl.String,
	// 	Prancheta:  mapper.LayoutToDto(*nprancheta),
	// }, log, config)

	imageResult, err := render.RenderPNGImage(
		ctx,
		renderer.RenderPngImageInput{Layout: *nprancheta},
	)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao tentar gerar imagem", "")
		return nil, err
	}
	// for cidx := range nprancheta.Components {
	// 	for eidx := range nprancheta.Components[cidx].Elements {
	// 		id := nprancheta.Components[cidx].Elements[eidx].ID
	// 		for _, resultElements := range res.Elements {
	// 			if resultElements.ElementID == id {
	// 				nprancheta.Components[cidx].Elements[eidx].ImageURL = resultElements.ImageURL
	// 			}
	// 		}
	// 	}
	// }
	gerated := *nprancheta
	// gerated.ImageURL = res.ImageURL
	gerated.ImageURL = imageResult.ImageURL
	layoutCreated, err := SaveLayout(ctx, gerated, queries, db)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao salvar layout", "")
		return nil, err
	}
	return &GenerateImageOutput{
		Data:   &GenerateImageResultV2{ImageURL: imageResult.ImageURL},
		Layout: layoutCreated,
	}, nil
}
