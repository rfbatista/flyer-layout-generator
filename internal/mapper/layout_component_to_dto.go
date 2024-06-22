package mapper

import "algvisual/internal/entities"

func DesignComponentToDto(e entities.LayoutComponent) entities.LayoutComponentDTO {
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
