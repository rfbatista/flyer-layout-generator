package entities

type LayoutDTO struct {
	ID         int32                `json:"id,omitempty"`
	Background *DesignComponentDTO  `json:"background,omitempty"`
	DesignID   int32                `json:"design_id,omitempty"`
	Width      int32                `json:"width,omitempty"`
	Height     int32                `json:"height,omitempty"`
	Components []DesignComponentDTO `json:"components,omitempty"`
	Template   Template             `json:"template,omitempty"`
	Grid       GridDTO              `json:"grid,omitempty"`
}

// Prancheta
type Layout struct {
	ID         int32             `json:"id,omitempty"`
	Background *DesignComponent  `json:"background,omitempty"`
	DesignID   int32             `json:"design_id,omitempty"`
	Width      int32             `json:"width,omitempty"`
	Height     int32             `json:"height,omitempty"`
	Components []DesignComponent `json:"components,omitempty"`
	Template   Template          `json:"template,omitempty"`
	Grid       Grid              `json:"grid,omitempty"`
	Stages     []Layout
	Config     LayoutRequestConfig
}
