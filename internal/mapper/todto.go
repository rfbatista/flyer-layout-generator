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
		Grid:       GridToDto(e.Grid),
	}
}

func GridToDto(g entities.Grid) entities.GridDTO {
	return entities.GridDTO{
		AllCells: g.Cells(),
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
		Width:    e.OuterContainer.Width(),
		Height:   e.OuterContainer.Height(),
		Color:    e.Color,
		Type:     e.Type,
		Xi:       e.OuterContainer.UpperLeft.X,
		Xii:      e.OuterContainer.DownRight.X,
		Yi:       e.OuterContainer.UpperLeft.Y,
		Yii:      e.OuterContainer.DownRight.Y,
	}
}

func DesignElementToDto(e entities.DesignElement) entities.DesignElementDTO {
	return entities.DesignElementDTO{
		ID:             e.ID,
		Xi:             e.OuterContainer.UpperLeft.X,
		Xii:            e.OuterContainer.DownRight.X,
		Yi:             e.OuterContainer.UpperLeft.Y,
		Yii:            e.OuterContainer.DownRight.Y,
		InnerXi:        e.InnerContainer.UpperLeft.X,
		InnerXii:       e.InnerContainer.DownRight.X,
		InnerYi:        e.InnerContainer.UpperLeft.Y,
		InnerYii:       e.InnerContainer.DownRight.Y,
		LayerID:        e.LayerID,
		Width:          e.OuterContainer.Width(),
		Height:         e.OuterContainer.Height(),
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
