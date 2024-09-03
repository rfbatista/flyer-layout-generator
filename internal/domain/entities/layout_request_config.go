package entities

type LayoutRequestConfig struct {
	LimitSizerPerElement  bool           `json:"limit_sizer_per_element,omitempty"`
	AnchorElements        bool           `json:"anchor_elements,omitempty"`
	ShowGrid              bool           `json:"show_grid,omitempty"`
	MinimiumComponentSize int32          `json:"minimium_component_size,omitempty"`
	MinimiumTextSize      int32          `json:"minimium_text_size,omitempty"`
	SlotsX                int32          `json:"slots_x,omitempty"`
	SlotsY                int32          `json:"slots_y,omitempty"`
	Grid                  Grid           `json:"grid,omitempty"`
	Padding               int32          `json:"padding,omitempty"`
	KeepProportions       bool           `json:"keep_proportions,omitempty"`
	Priorities            map[string]int `json:"priorities,omitempty"`
	TemplatesID           []int32        `json:"template_id,omitempty"`
}
