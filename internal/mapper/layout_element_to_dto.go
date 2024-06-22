package mapper

import "algvisual/internal/entities"

func DesignElementToDto(e entities.LayoutElement) entities.LayoutElementDTO {
	return entities.LayoutElementDTO{
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
