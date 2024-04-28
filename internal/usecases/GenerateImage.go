package usecases

import (
	"context"

	"algvisual/internal/database"
	"algvisual/internal/entities"
	"algvisual/internal/infra"
	"algvisual/internal/shared"
)

type GenerateImageRequest struct {
	PhotoshopID int `json:"photoshop_id,omitempty"`
	TemplateID  int `json:"template_id,omitempty"`
}

type GenerateImageResult struct {
	Data *infra.GeneratorResult `json:"data,omitempty"`
}

func GenerateImageUseCase(
	ctx context.Context,
	req GenerateImageRequest,
	client *infra.ImageGeneratorClient,
	queries *database.Queries,
) (*GenerateImageResult, error) {
	photoshop, err := queries.GetPhotoshop(ctx, int32(req.PhotoshopID))
	if err != nil {
		err = shared.WrapWithAppError(err, "Não foi possivel encontrar o arquivo Photoshop", "")
		return nil, err
	}
	template, err := queries.GetTemplate(ctx, int32(req.PhotoshopID))
	if err != nil {
		err = shared.WrapWithAppError(err, "Não foi possivel encontrar o arquivo Photoshop", "")
		return nil, err
	}
	slots, err := queries.GetTemplateSlots(ctx, template.Template.ID)
	if err != nil {
		err = shared.WrapWithAppError(err, "Não foi possivel encontrar o arquivo Photoshop", "")
		return nil, err
	}
	distortionConfig, err := queries.GetTemplateDistortion(ctx, template.Template.ID)
	if err != nil {
		err = shared.WrapWithAppError(err, "Não foi possivel encontrar o arquivo Photoshop", "")
		return nil, err
	}
	etemplate := database.ToTemplateEntitie(template.Template)
	etemplate.Distortion = database.ToTemplateDistortionEntitie(
		distortionConfig.TemplatesDistortion,
	)
	var eslots []entities.TemplateSlotsPositions
	for _, s := range slots {
		eslots = append(eslots, database.ToTemplateSlotEntitie(s.TemplatesSlot))
	}
	elements, err := queries.GetPhotoshopElements(ctx, int32(req.PhotoshopID))
	if err != nil {
		err = shared.WrapWithAppError(
			err,
			"Não foi possivel encontrar os elementos do arquivo Photoshop",
			"",
		)
		return nil, err
	}
	var eelements []entities.PhotoshopElement
	for _, el := range elements {
		eelements = append(eelements, database.ToPhotoshopElementEntitie(el))
	}
	compHash := make(map[int32][]entities.PhotoshopElement)
	for _, c := range eelements {
		if c.ComponentID != 0 {
			compHash[c.ComponentID] = append(compHash[c.ComponentID], c)
		}
	}
	var components []entities.PhotoshopComponent
	for k := range compHash {
		components = append(components, entities.PhotoshopComponent{
			ID:       k,
			Elements: compHash[k],
		})
	}
	etemplate.SlotsPositions = eslots
	result, err := client.GenerateImageWithDistortionStrategy(
		infra.GeneratorRequest{
			Photoshop:  database.ToPhotoshopEntitie(photoshop),
			Template:   etemplate,
			Elements:   eelements,
			Components: components,
		},
	)
	if err != nil {
		err = shared.WrapWithAppError(err, "Falha ao tentar gerar imagem", "")
		return nil, err
	}
	return &GenerateImageResult{
		Data: result,
	}, nil
}
