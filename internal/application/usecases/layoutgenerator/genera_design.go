package layoutgenerator

import (
	"context"

	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/repositories/mapper"
	"algvisual/internal/shared"
)

type GenerateDesignRequest struct {
	PhotoshopID int32 `form:"photoshop_id" json:"photoshop_id,omitempty"`
	TemplateID  int32 `form:"template_id"  json:"template_id,omitempty"`
}

type GenerateDesignResult struct {
	Data *infrastructure.GeneratorResult `json:"data,omitempty"`
}

func GenerateDesignUseCase(
	ctx context.Context,
	req GenerateDesignRequest,
	client *infrastructure.ImageGeneratorClient,
	queries *database.Queries,
) (*GenerateDesignResult, error) {
	photoshop, err := queries.Getdesign(ctx, req.PhotoshopID)
	if err != nil {
		err = shared.WrapWithAppError(err, "Não foi possivel encontrar o photoshop", "")
		return nil, err
	}
	template, err := queries.GetTemplate(ctx, req.TemplateID)
	if err != nil {
		err = shared.WrapWithAppError(err, "Não foi possivel encontrar o templates", "")
		return nil, err
	}
	etemplate := mapper.TemplateToDomain(template.Template)
	elements, err := queries.GetElements(ctx, photoshop.ID)
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
		components = append(components, comp)
	}
	result, err := client.GenerateImageWithDistortionStrategy(
		infrastructure.GeneratorRequest{
			Photoshop:  mapper.TodesignEntitie(photoshop),
			Template:   etemplate,
			Elements:   eelements,
			Components: components,
		},
	)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao tentar gerar imagem", "")
		return nil, err
	}
	return &GenerateDesignResult{
		Data: result,
	}, nil
}
