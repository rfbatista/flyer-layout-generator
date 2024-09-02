package layoutgenerator

import (
	"algvisual/internal/application/usecases/grammars"
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

type GenerateDesignRequestv3 struct {
	PhotoshopID int32                        `form:"photoshop_id" json:"photoshop_id,omitempty"`
	TemplateID  int32                        `form:"template_id"  json:"template_id,omitempty"`
	Config      entities.LayoutRequestConfig `                    json:"config,omitempty"`
}

type GenerateDesignResultv3 struct {
	Data       *GenerateImageResultV2 `json:"data,omitempty"`
	TwistedURL string
	Layout     *entities.Layout
}

func GenerateDesignUseCasev3(
	ctx context.Context,
	req GenerateDesignRequestv3,
	queries *database.Queries,
	db *pgxpool.Pool,
	config config.AppConfig,
	log *zap.Logger,
) (*GenerateDesignResultv3, error) {
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
	var eelements []entities.LayoutElement
	for _, el := range elements {
		eelements = append(eelements, mapper.ToDesignElementEntitie(el))
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
	}
	nprancheta, err := grammars.RunV1(
		prancheta,
		mapper.TemplateToDomain(template.Template),
		req.Config.SlotsX,
		req.Config.SlotsY,
		log,
	)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao tentar gerar imagem", "")
		return nil, err
	}
	res, err := GenerateImageFromPranchetaV2(GenerateImageRequestV2{
		DesignFile: designFile.FileUrl.String,
		Prancheta:  mapper.LayoutToDto(*nprancheta),
	}, log, config)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao tentar gerar imagem", "")
		return nil, err
	}
	layoutCreated, err := SaveLayout(ctx, *nprancheta, queries, db)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao salvar layout", "")
		return nil, err
	}
	return &GenerateDesignResultv3{
		Data:   res,
		Layout: layoutCreated,
	}, nil
}
