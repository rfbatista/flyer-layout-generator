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
	ID             int32               `json:"id,omitempty"`
	RequestID      int32               `json:"request_id,omitempty"`
	Background     *LayoutComponent    `json:"background,omitempty"`
	BackgroundList []LayoutComponent   `json:"background_list,omitempty"`
	ImageURL       string              `json:"image_url,omitempty"`
	DesignID       int32               `json:"design_id,omitempty"`
	Width          int32               `json:"width,omitempty"`
	Height         int32               `json:"height,omitempty"`
	Components     []LayoutComponent   `json:"components,omitempty"`
	Elements       []LayoutElement     `json:"elements,omitempty"`
	Template       Template            `json:"template,omitempty"`
	Grid           Grid                `json:"grid,omitempty"`
	Stages         []string            `json:"stages,omitempty"`
	Config         LayoutRequestConfig `json:"config,omitempty"`
	DesignAssets   []DesignAsset       `json:"design_assets,omitempty"`
}

func ListToPrioritiesMap(list []string) map[string]int {
	prioritiesMap := make(map[string]int)
	for i, v := range list {
		compType := StringToComponentType(v)
		prioritiesMap[compType.ToString()] = i
	}
	return prioritiesMap
}
