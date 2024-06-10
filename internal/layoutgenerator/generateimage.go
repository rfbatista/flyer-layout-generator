package layoutgenerator

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/grammars"
	"algvisual/internal/infra"
	"algvisual/internal/mapper"
	"algvisual/internal/shared"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type GenerateImage struct {
	PhotoshopID           int32 `form:"design_id"   json:"photoshop_id,omitempty"`
	TemplateID            int32 `form:"template_id" json:"template_id,omitempty"`
	LimitSizerPerElement  bool  `                   json:"limit_sizer_per_element,omitempty"`
	AnchorElements        bool  `                   json:"anchor_elements,omitempty"`
	ShowGrid              bool  `                   json:"show_grid,omitempty"`
	MinimiumComponentSize int32 `                   json:"minimium_component_size,omitempty"`
	MinimiumTextSize      int32 `                   json:"minimium_text_size,omitempty"`
	SlotsX                int32 `form:"grid_x"      json:"slots_x,omitempty"`
	SlotsY                int32 `form:"grid_y"      json:"slots_y,omitempty"`
	Padding               int32 `form:"padding"                   json:"padding,omitempty"`
	KeepProportions       bool  `                   json:"keep_proportions,omitempty"`
}

type GenerateImageOutput struct {
	Data       *GenerateImageResultV2 `json:"data,omitempty"`
	TwistedURL string
}

func GenerateImageUseCase(
	ctx context.Context,
	req GenerateImage,
	queries *database.Queries,
	db *pgxpool.Pool,
	config infra.AppConfig,
	log *zap.Logger,
) (*GenerateImageOutput, error) {
	designFile, err := queries.Getdesign(ctx, req.PhotoshopID)
	if err != nil {
		err = shared.WrapWithAppError(err, "Não foi possivel encontrar o photoshop", "")
		return nil, err
	}
	template, err := queries.GetTemplate(ctx, req.TemplateID)
	if err != nil {
		err = shared.WrapWithAppError(
			err,
			fmt.Sprintf("Não foi possivel encontrar o template %d", req.TemplateID),
			"",
		)
		return nil, err
	}
	etemplate := mapper.TemplateToDomain(template.Template)
	elements, err := queries.GetElements(ctx, designFile.ID)
	if err != nil {
		err = shared.WrapWithAppError(
			err,
			"Não foi possivel encontrar os elementos do arquivo Photoshop",
			err.Error(),
		)
		return nil, err
	}
	var eelements []entities.DesignElement
	for _, el := range elements {
		eelements = append(eelements, mapper.ToDesignElementEntitie(el))
	}
	compHash := make(map[int32][]entities.DesignElement)
	for _, c := range eelements {
		if c.ComponentID != 0 {
			compHash[c.ComponentID] = append(compHash[c.ComponentID], c)
		}
	}
	var components []entities.DesignComponent
	var bg *entities.DesignComponent
	for k := range compHash {
		data, compErr := queries.GetComponentByID(ctx, k)
		if compErr != nil {
			compErr = shared.WrapWithAppError(
				compErr,
				"Não foi possivel encontrar os componentes do arquivo Photoshop",
				"",
			)
			return nil, compErr
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
		return nil, shared.NewAppError(
			400,
			"nenhum componente definido para o design escolhido",
			"",
		)
	}
	prancheta := entities.Layout{
		DesignID:   req.PhotoshopID,
		Width:      designFile.Width.Int32,
		Height:     designFile.Height.Int32,
		Template:   etemplate,
		Background: bg,
		Components: components,
		Config: entities.LayoutRequestConfig{
			Padding: req.Padding,
		},
	}
	nprancheta, _ := grammars.RunV2(prancheta, etemplate, req.SlotsX, req.SlotsY, log)
	if !req.ShowGrid {
		nprancheta.Grid = entities.Grid{}
	}
	res, err := GenerateImageFromPranchetaV2(GenerateImageRequestV2{
		DesignFile: designFile.FileUrl.String,
		Prancheta:  mapper.LayoutToDto(*nprancheta),
	}, log, config)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao tentar gerar imagem", "")
		return nil, err
	}
	err = SaveLayout(ctx, *nprancheta, queries, db)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao salvar layout", "")
		return nil, err
	}
	return &GenerateImageOutput{
		Data: res,
	}, nil
}
