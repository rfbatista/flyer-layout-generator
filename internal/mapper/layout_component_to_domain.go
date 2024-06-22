package mapper

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
)

func TodesignComponentEntitie(raw database.LayoutComponent) entities.LayoutComponent {
	return entities.LayoutComponent{
		ID:      raw.ID,
		FWidth:  raw.Width.Int32,
		FHeight: raw.Height.Int32,
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
		OuterContainer: entities.NewContainer(
			entities.NewPoint(raw.Xi.Int32, raw.Yi.Int32),
			entities.NewPoint(raw.Xii.Int32, raw.Yii.Int32),
		),
		InnerContainer: entities.NewContainer(
			entities.NewPoint(raw.InnerXi.Int32, raw.InnerYi.Int32),
			entities.NewPoint(raw.InnerXii.Int32, raw.InnerYii.Int32),
		),
		Priority: raw.Priority.Int32,
	}
}

func LayoutComponentToDomain(raw database.LayoutComponent) entities.LayoutComponent {
	return entities.LayoutComponent{
		ID:      int32(raw.ID),
		FWidth:  raw.Width.Int32,
		FHeight: raw.Height.Int32,
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
		OuterContainer: entities.NewContainer(
			entities.NewPoint(raw.Xi.Int32, raw.Yi.Int32),
			entities.NewPoint(raw.Xii.Int32, raw.Yii.Int32),
		),
		InnerContainer: entities.NewContainer(
			entities.NewPoint(raw.Xi.Int32, raw.Yi.Int32),
			entities.NewPoint(raw.Xii.Int32, raw.Yii.Int32),
		),
	}
}
