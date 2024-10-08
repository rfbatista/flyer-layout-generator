package mapper

import "algvisual/internal/domain/entities"

func LayoutToDatabase(e entities.Layout) entities.LayoutDTO {
	var cc []entities.LayoutComponentDTO
	for _, c := range e.Components {
		cc = append(cc, DesignComponentToDto(c))
	}
	var bg *entities.LayoutComponentDTO
	if e.Background != nil {
		bge := DesignComponentToDto(*e.Background)
		bg = &bge
	}
	return entities.LayoutDTO{
		ID:         e.ID,
		Background: bg,
		ImageURL:   e.ImageURL,
		DesignID:   e.DesignID,
		Width:      e.Width,
		Height:     e.Height,
		Components: cc,
		Template:   e.Template,
		Grid:       GridToDto(e.Grid),
	}
}

func DesignComponentToDatabase(e entities.LayoutComponent) entities.LayoutComponentDTO {
	var elements []entities.LayoutElementDTO
	for _, c := range e.Elements {
		elements = append(elements, DesignElementToDto(c))
	}
	return entities.LayoutComponentDTO{
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
