package entities

import "time"

type Photoshop struct {
	ID             int32     `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	Filepath       string    `json:"filepath,omitempty"`
	FileExtension  string    `json:"file_extension,omitempty"`
	ImagePath      string    `json:"image_path,omitempty"`
	ImageExtension string    `json:"image_extension,omitempty"`
	Width          int32     `json:"width,omitempty"`
	Height         int32     `json:"height,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

type PhotoshopElement struct {
	ID             int32  `json:"id,omitempty"`
	Xi             int32  `json:"xi,omitempty"`
	Xii            int32  `json:"xii,omitempty"`
	Yi             int32  `json:"yi,omitempty"`
	Yii            int32  `json:"yii,omitempty"`
	LayerID        string `json:"layer_id,omitempty"`
	Width          int32  `json:"width,omitempty"`
	Height         int32  `json:"height,omitempty"`
	Kind           string `json:"kind,omitempty"`
	Name           string `json:"name,omitempty"`
	IsGroup        bool   `json:"is_group,omitempty"`
	GroupId        int32  `json:"group_id,omitempty"`
	Level          int32  `json:"level,omitempty"`
	PhotoshopId    int32  `json:"photoshop_id,omitempty"`
	Image          string `json:"image,omitempty"`
	Text           string `json:"text,omitempty"`
	ImageExtension string `json:"image_extension,omitempty"`
	ComponentID    int32  `json:"component_id,omitempty"`
}

type PhotoshopComponent struct {
	ID       int32
	Elements []PhotoshopElement
}
