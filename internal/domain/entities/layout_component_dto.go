package entities

type LayoutComponentDTO struct {
	ID       int32              `json:"id,omitempty"`
	DesignID int32              `json:"design_id,omitempty"`
	Elements []LayoutElementDTO `json:"elements,omitempty"`
	Width    int32              `json:"width"`
	Height   int32              `json:"height"`
	Color    string             `json:"color,omitempty"`
	Type     string             `json:"type,omitempty"`
	Xi       int32              `json:"xi"`
	Xii      int32              `json:"xii"`
	Yi       int32              `json:"yi"`
	Yii      int32              `json:"yii"`
	BboxXi   int32              `json:"bbox_xi,omitempty"`
	BboxXii  int32              `json:"bbox_xii,omitempty"`
	BboxYi   int32              `json:"bbox_yi,omitempty"`
	BboxYii  int32              `json:"bbox_yii,omitempty"`
	Xsnaped  bool               `json:"xsnaped,omitempty"`
	Ysnaped  bool               `json:"ysnaped,omitempty"`
	LeftGap  Position           `json:"left_gap,omitempty"`
	RightGap Position           `json:"right_gap,omitempty"`
}
