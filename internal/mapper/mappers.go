package mapper

import (
	"algvisual/internal/database"
	designprocessor "algvisual/internal/entities"
)

func TodesignEntitie(raw database.Design) designprocessor.DesignFile {
	return designprocessor.DesignFile{
		ID:             raw.ID,
		Width:          raw.Width.Int32,
		Height:         raw.Height.Int32,
		Name:           raw.Name,
		Filepath:       raw.FileUrl.String,
		ImageExtension: raw.ImageExtension.String,
		ImagePath:      raw.ImageUrl.String,
		CreatedAt:      raw.CreatedAt.Time,
	}
}

func TodesignComponentEntitie(raw database.DesignComponent) designprocessor.DesignComponent {
	return designprocessor.DesignComponent{
		ID:      raw.ID,
		Width:   raw.Width.Int32,
		Height:  raw.Height.Int32,
		Xi:      raw.Xi.Int32,
		Xii:     raw.Xii.Int32,
		Yi:      raw.Yi.Int32,
		Yii:     raw.Yii.Int32,
		Color:   raw.Color.String,
		Type:    string(raw.Type.ComponentType),
		BboxXi:  raw.BboxXi.Int32,
		BboxXii: raw.BboxXii.Int32,
		BboxYi:  raw.BboxYi.Int32,
		BboxYii: raw.BboxYii.Int32,
	}
}

func ToDesignElementEntitie(raw database.DesignElement) designprocessor.DesignElement {
	return designprocessor.DesignElement{
		ID:          raw.ID,
		Xi:          raw.Xi.Int32,
		Xii:         raw.Xii.Int32,
		Yi:          raw.Yi.Int32,
		Yii:         raw.Yii.Int32,
		LayerID:     raw.LayerID.String,
		Width:       raw.Width.Int32,
		Height:      raw.Height.Int32,
		Kind:        raw.Kind.String,
		Name:        raw.Name.String,
		IsGroup:     raw.IsGroup.Bool,
		GroupId:     raw.GroupID.Int32,
		Level:       raw.Level.Int32,
		DesignID:    raw.DesignID,
		ImageURL:    raw.ImageUrl.String,
		Text:        raw.Text.String,
		ComponentID: raw.ComponentID.Int32,
	}
}

func ToDesignElementEntitieList(raw []database.DesignElement) []designprocessor.DesignElement {
	var e []designprocessor.DesignElement
	for _, r := range raw {
		e = append(e, ToDesignElementEntitie(r))
	}
	return e
}

func ToTemplateEntitie(raw database.Template) designprocessor.Template {
	return designprocessor.Template{
		ID:     raw.ID,
		Width:  raw.Width.Int32,
		Height: raw.Height.Int32,
		Type:   designprocessor.NewTemplateType(string(raw.Type.TemplateType)),
	}
}

func ToTemplateSlotEntitie(raw database.TemplatesSlot) designprocessor.TemplateSlotsPositions {
	return designprocessor.TemplateSlotsPositions{
		Xi:     raw.Xi.Int32,
		Yi:     raw.Yi.Int32,
		Width:  raw.Width.Int32,
		Height: raw.Height.Int32,
	}
}

func ToTemplateDistortionEntitie(raw database.TemplatesDistortion) designprocessor.TemplateDistortion {
	return designprocessor.TemplateDistortion{
		X: raw.X.Int32,
		Y: raw.Y.Int32,
	}
}
