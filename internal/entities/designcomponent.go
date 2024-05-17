package entities

type DesignComponent struct {
	ID       int32           `json:"id,omitempty"`
	Elements []DesignElement `json:"elements,omitempty"`
	Width    int32           `json:"width"`
	Height   int32           `json:"height"`
	Color    string          `json:"color,omitempty"`
	Type     string          `json:"type,omitempty"`
	Xi       int32           `json:"xi"`
	Xii      int32           `json:"xii"`
	Yi       int32           `json:"yi"`
	Yii      int32           `json:"yii"`
	BboxXi   int32           `json:"bbox_xi"`
	BboxXii  int32           `json:"bbox_xii"`
	BboxYi   int32           `json:"bbox_yi"`
	BboxYii  int32           `json:"bbox_yii"`
}
