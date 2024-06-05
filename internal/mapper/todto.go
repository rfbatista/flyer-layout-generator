package mapper

import "algvisual/internal/entities"

func LayoutToDto(e entities.Layout) entities.LayoutDTO {
	var cc []entities.DesignComponentDTO
	for _, c := range e.Components {
		cc = append(cc, DesignComponentToDto(c))
	}
	var bg *entities.DesignComponentDTO
	if e.Background != nil {
		bge := DesignComponentToDto(*e.Background)
		bg = &bge
	}
	return entities.LayoutDTO{
		ID:         e.ID,
		Background: bg,
		DesignID:   e.DesignID,
		Width:      e.Width,
		Height:     e.Height,
		Components: cc,
		Template:   e.Template,
		Grid:       e.Grid,
	}
}

func DesignComponentToDto(e entities.DesignComponent) entities.DesignComponentDTO {
	var elements []entities.DesignElementDTO
	for _, c := range e.Elements {
		elements = append(elements, DesignElementToDto(c))
	}
	return entities.DesignComponentDTO{
		ID:       e.ID,
		Elements: elements,
		DesignID: e.DesignID,
		Width:    e.Width(),
		Height:   e.Height(),
		Color:    e.Color,
		Type:     e.Type,
		Xi:       e.UpLeft().X,
		Xii:      e.DownRight().X,
		Yi:       e.UpLeft().Y,
		Yii:      e.DownRight().Y,
	}
}

func DesignElementToDto(e entities.DesignElement) entities.DesignElementDTO {
	return entities.DesignElementDTO{
		ID:             e.ID,
		Xi:             e.UpLeft().X,
		Xii:            e.DownRight().X,
		Yi:             e.UpLeft().Y,
		Yii:            e.DownRight().Y,
		LayerID:        e.LayerID,
		Width:          e.Width(),
		Height:         e.Height(),
		Kind:           e.Kind,
		Name:           e.Name,
		IsGroup:        e.IsGroup,
		GroupId:        e.GroupId,
		Level:          e.Level,
		DesignID:       e.DesignID,
		ImageURL:       e.ImageURL,
		Text:           e.Text,
		ImageExtension: e.ImageExtension,
		ComponentID:    e.ComponentID,
	}
}
