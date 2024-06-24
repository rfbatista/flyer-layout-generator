package generate

type GenerateImage struct {
	PhotoshopID           int32    `form:"design_id"   json:"photoshop_id,omitempty"`
	LayoutID              int32    `form:"layout_id"   json:"layout_id,omitempty"`
	TemplateID            []int32  `form:"templates[]" json:"template_id,omitempty"`
	LimitSizerPerElement  bool     `                   json:"limit_sizer_per_element,omitempty"`
	AnchorElements        bool     `                   json:"anchor_elements,omitempty"`
	ShowGrid              bool     `form:"show_grid"   json:"show_grid,omitempty"`
	MinimiumComponentSize int32    `                   json:"minimium_component_size,omitempty"`
	MinimiumTextSize      int32    `                   json:"minimium_text_size,omitempty"`
	SlotsX                int32    `form:"grid_x"      json:"slots_x,omitempty"`
	SlotsY                int32    `form:"grid_y"      json:"slots_y,omitempty"`
	Padding               int32    `form:"padding"     json:"padding,omitempty"`
	KeepProportions       bool     `                   json:"keep_proportions,omitempty"`
	Priorities            []string `form:"priority[]"  json:"priorities,omitempty"`
}
