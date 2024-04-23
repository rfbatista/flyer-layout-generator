package entities

type Photoshop struct {
	Filepath string `json:"filepath,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

type PhotoshopElement struct {
	Xi          int    `json:"xi,omitempty"`
	Xii         int    `json:"xii,omitempty"`
	Yi          int    `json:"yi,omitempty"`
	Yii         int    `json:"yii,omitempty"`
	LayerID     string `json:"layer_id,omitempty"`
	Width       int    `json:"width,omitempty"`
	Height      int    `json:"height,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Name        string `json:"name,omitempty"`
	IsGroup     bool   `json:"is_group,omitempty"`
	GroupId     int    `json:"group_id,omitempty"`
	Level       int    `json:"level,omitempty"`
	PhotoshopId int    `json:"photoshop_id,omitempty"`
	Image       string `json:"image,omitempty"`
	Text        string `json:"text,omitempty"`
}
