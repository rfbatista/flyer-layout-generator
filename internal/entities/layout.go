package entities

type LayoutDTO struct {
	ID         int32                `json:"id,omitempty"`
	Background *LayoutComponentDTO  `json:"background,omitempty"`
	ImageURL   string               `json:"image_url,omitempty"`
	DesignID   int32                `json:"design_id,omitempty"`
	Width      int32                `json:"width,omitempty"`
	Height     int32                `json:"height,omitempty"`
	Components []LayoutComponentDTO `json:"components,omitempty"`
	Template   Template             `json:"template,omitempty"`
	Grid       GridDTO              `json:"grid,omitempty"`
}

// Prancheta
type Layout struct {
	ID         int32               `json:"id,omitempty"`
	Background *LayoutComponent    `json:"background,omitempty"`
	ImageURL   string              `json:"image_url,omitempty"`
	DesignID   int32               `json:"design_id,omitempty"`
	Width      int32               `json:"width,omitempty"`
	Height     int32               `json:"height,omitempty"`
	Components []LayoutComponent   `json:"components,omitempty"`
	Elements   []LayoutElement     `json:"elements,omitempty"`
	Template   Template            `json:"template,omitempty"`
	Grid       Grid                `json:"grid,omitempty"`
	Stages     []Layout            `json:"stages,omitempty"`
	Config     LayoutRequestConfig `json:"config,omitempty"`
}

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
	NPriorities           []string       `json:"priorities,omitempty"              form:"priority[]"`
}

func ListToPrioritiesMap(list []string) map[string]int {
	prioritiesMap := make(map[string]int)
	for i, v := range list {
		prioritiesMap[v] = i
	}
	return prioritiesMap
}
