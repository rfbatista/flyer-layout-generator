package mapper

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
	"encoding/json"
	"fmt"
	"time"
)

func LayoutRequestToDomain(raw database.LayoutRequest) entities.LayoutRequest {
	var createdAt *time.Time
	var startedAt *time.Time
	var stoppedAt *time.Time
	var errorAt *time.Time
	if raw.CreatedAt.Valid {
		createdAt = &raw.CreatedAt.Time
	}
	return entities.LayoutRequest{
		ID:        int32(raw.ID),
		DesignID:  raw.DesignID.Int32,
		CreatedAt: createdAt,
		StartedAt: startedAt,
		StoppedAt: stoppedAt,
		ErrorAt:   errorAt,
	}
}

func LayoutRequestJobToDomain(raw database.LayoutRequestsJob) entities.LayoutRequestJob {
	l := entities.LayoutRequestJob{
		ID:         int32(raw.ID),
		RequestID:  raw.RequestID.Int32,
		TemplateID: raw.TemplateID.Int32,
		Status:     raw.Status.String,
		Log:        raw.Log.String,
		ImageURL:   raw.ImageUrl.String,
	}
	if raw.Config.String != "" {
		var c entities.LayoutRequestConfig
		err := json.Unmarshal([]byte(raw.Config.String), &c)
		if err != nil {
			fmt.Println("falha ao realizar parser da config")
		} else {
			l.Config = &c
		}
	}
	if raw.StartedAt.Valid {
		l.StartedAt = &raw.StartedAt.Time
	}
	if raw.CreatedAt.Valid {
		l.CreatedAt = &raw.CreatedAt.Time
	}
	if raw.StoppedAt.Valid {
		l.StoppedAt = &raw.StoppedAt.Time
	}
	if raw.FinishedAt.Valid {
		l.FinishedAt = &raw.FinishedAt.Time
	}
	if raw.ErrorAt.Valid {
		l.ErrorAt = &raw.ErrorAt.Time
	}
	return l
}

func DesignFileToDomain(raw database.Design) entities.DesignFile {
	return entities.DesignFile{
		ID:             raw.ID,
		Name:           raw.Name,
		Filepath:       raw.FileUrl.String,
		FileExtension:  raw.FileExtension.String,
		ImagePath:      raw.ImageUrl.String,
		ImageURL:       raw.ImageUrl.String,
		ImageExtension: raw.ImageExtension.String,
		Width:          raw.Width.Int32,
		Height:         raw.Height.Int32,
		CreatedAt:      raw.CreatedAt.Time,
		IsProcessed:    raw.IsProccessed.Bool,
	}
}

func LayoutToDomain(raw database.Layout) entities.Layout {
	return entities.Layout{
		ID:     int32(raw.ID),
		Width:  raw.Width.Int32,
		Height: raw.Height.Int32,
	}
}

func LayoutRegionToDomain(raw database.LayoutRegion) entities.GridCell {
	return entities.GridCell{
		Xi:  raw.Xi.Int32,
		Xii: raw.Xii.Int32,
		Yi:  raw.Yi.Int32,
		Yii: raw.Yii.Int32,
	}
}

func LayoutComponentToDomain(raw database.LayoutComponent) entities.DesignComponent {
	return entities.DesignComponent{
		ID:      int32(raw.ID),
		FWidth:  raw.Width.Int32,
		FHeight: raw.Height.Int32,
		Xi:      raw.Xi.Int32,
		Xii:     raw.Xii.Int32,
		Yi:      raw.Yi.Int32,
		Yii:     raw.Yii.Int32,
		Color:   raw.Color.String,
		Type:    string(raw.Type.String),
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

func TodesignEntitie(raw database.Design) entities.DesignFile {
	return entities.DesignFile{
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

func TodesignComponentEntitie(raw database.DesignComponent) entities.DesignComponent {
	return entities.DesignComponent{
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

func ToDesignElementEntitie(raw database.DesignElement) entities.DesignElement {
	return entities.DesignElement{
		ID:          raw.ID,
		Xi:          raw.Xi.Int32,
		Xii:         raw.Xii.Int32,
		Yi:          raw.Yi.Int32,
		Yii:         raw.Yii.Int32,
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

func ToDesignElementEntitieList(raw []database.DesignElement) []entities.DesignElement {
	var e []entities.DesignElement
	for _, r := range raw {
		e = append(e, ToDesignElementEntitie(r))
	}
	return e
}

func TemplateToDomain(raw database.Template) entities.Template {
	return entities.Template{
		ID:        raw.ID,
		Name:      raw.Name,
		Width:     raw.Width.Int32,
		Height:    raw.Height.Int32,
		Type:      entities.NewTemplateType(string(raw.Type.TemplateType)),
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
