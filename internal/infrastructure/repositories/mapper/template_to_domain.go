package mapper

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
)

func TemplateToDomain(raw database.Template) entities.Template {
	return entities.Template{
		ID:        raw.ID,
		Name:      raw.Name,
		Width:     raw.Width.Int32,
		Height:    raw.Height.Int32,
		MaxSlotsX: raw.MaxSlotsX.Int32,
		MaxSlotsY: raw.MaxSlotsY.Int32,
		CreatedAt: raw.CreatedAt.Time,
	}
}

func ToTemplateSlotEntitie(raw database.TemplatesSlot) entities.TemplateSlotsPositions {
	return entities.TemplateSlotsPositions{
		Xi:     raw.Xi.Int32,
		Yi:     raw.Yi.Int32,
		Width:  raw.Width.Int32,
		Height: raw.Height.Int32,
	}
}

func ToTemplateDistortionEntitie(
	raw database.TemplatesDistortion,
) entities.TemplateDistortion {
	return entities.TemplateDistortion{
		X: raw.X.Int32,
		Y: raw.Y.Int32,
	}
}

func TemplateTypeToDatabase(t entities.TemplateType) database.TemplateType {
	switch t {
	case entities.TemplateTypeAdaptation:
		return database.TemplateTypeAdaptation
	case entities.TemplateTypePublic:
		return database.TemplateTypePublic
	}

	return database.TemplateTypeUnknown
}
