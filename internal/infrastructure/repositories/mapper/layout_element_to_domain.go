package mapper

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/database"
)

func ToDesignElementEntitie(raw database.LayoutElement) entities.LayoutElement {
	return entities.LayoutElement{
		ID:          raw.ID,
		Xi:          raw.Xi.Int32,
		Xii:         raw.Xii.Int32,
		Yi:          raw.Yi.Int32,
		Yii:         raw.Yii.Int32,
		LayoutID:    raw.LayoutID,
		AssetID:     raw.AssetID,
		InnerXi:     raw.InnerXi.Int32,
		InnerXii:    raw.InnerXii.Int32,
		InnerYi:     raw.InnerYi.Int32,
		InnerYii:    raw.InnerYii.Int32,
		LayerID:     raw.LayerID.String,
		FWidth:      raw.Width.Int32,
		FHeight:     raw.Height.Int32,
		Kind:        raw.Kind.String,
		Name:        raw.Name.String,
		IsGroup:     raw.IsGroup.Bool,
		GroupId:     raw.GroupID.Int32,
		Level:       raw.Level.Int32,
		DesignID:    raw.DesignID,
		ImageURL:    raw.ImageUrl.String,
		Text:        raw.Text.String,
		ComponentID: raw.ComponentID.Int32,
		OuterContainer: entities.NewContainer(
			entities.NewPoint(raw.Xi.Int32, raw.Yi.Int32),
			entities.NewPoint(raw.Xii.Int32, raw.Yii.Int32),
		),
		InnerContainer: entities.NewContainer(
			entities.NewPoint(raw.InnerXi.Int32, raw.InnerYi.Int32),
			entities.NewPoint(raw.InnerXii.Int32, raw.InnerYii.Int32),
		),
	}
}

func ToDesignElementEntitieList(raw []database.LayoutElement) []entities.LayoutElement {
	var e []entities.LayoutElement
	for _, r := range raw {
		e = append(e, ToDesignElementEntitie(r))
	}
	return e
}
