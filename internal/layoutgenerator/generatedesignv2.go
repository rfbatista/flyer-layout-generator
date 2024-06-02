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
	"math/rand"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type GenerateDesignRequestv2 struct {
	PhotoshopID int32                        `form:"photoshop_id" json:"photoshop_id,omitempty"`
	TemplateID  int32                        `form:"template_id" json:"template_id,omitempty"`
	Config      entities.LayoutRequestConfig `json:"config,omitempty"`
}

type GenerateDesignResultv2 struct {
	Data       *GenerateImageResult `json:"data,omitempty"`
	TwistedURL string
}

func GenerateDesignUseCasev2(
	ctx context.Context,
	req GenerateDesignRequestv2,
	queries *database.Queries,
	db *pgxpool.Pool,
	config infra.AppConfig,
	log *zap.Logger,
) (*GenerateDesignResultv2, error) {
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
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	world := grammars.World{
		OriginalDesign: mapper.TodesignEntitie(designFile),
		Components:     components,
		Elements:       eelements,
		PivotWidth:     components[r.Intn(len(components))].Width,
		PivotHeight:    components[r.Intn(len(components))].Height,
		Config:         req.Config,
	}
	prancheta := entities.Layout{
		DesignID:   req.PhotoshopID,
		Width:      etemplate.Width,
		Height:     etemplate.Height,
		Template:   etemplate,
		Background: bg,
	}
	world, nprancheta, _ := grammars.Run(world, prancheta, log)
	res, err := GenerateImageFromPrancheta(GenerateImageRequest{
		DesignFile: designFile.FileUrl.String,
		Prancheta:  nprancheta,
	}, log, config)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao tentar gerar imagem", "")
		return nil, err
	}
	err = SaveLayout(ctx, nprancheta, queries, db)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao salvar layout", "")
		return nil, err
	}
	return &GenerateDesignResultv2{
		Data: res,
	}, nil
}
