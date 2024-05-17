package entities

type DesignElement struct {
	ID             int32  `json:"id,omitempty"`
	Xi             int32  `json:"xi"`
	Xii            int32  `json:"xii"`
	Yi             int32  `json:"yi"`
	Yii            int32  `json:"yii"`
	LayerID        string `json:"layer_id"`
	Width          int32  `json:"width"`
	Height         int32  `json:"height"`
	Kind           string `json:"kind"`
	Name           string `json:"name"`
	IsGroup        bool   `json:"is_group"`
	GroupId        int32  `json:"group_id"`
	Level          int32  `json:"level"`
	DesignID       int32  `json:"photoshop_id"`
	ImageURL       string `json:"image,omitempty"`
	Text           string `json:"text,omitempty"`
	ImageExtension string `json:"image_extension,omitempty"`
	ComponentID    int32  `json:"component_id"`
}
