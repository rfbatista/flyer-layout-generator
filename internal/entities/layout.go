package entities

type Layout struct {
	Width      int32             `json:"width,omitempty"`
	Height     int32             `json:"height,omitempty"`
	Components []DesignComponent `json:"components,omitempty"`
	Template   Template          `json:"templates,omitempty"`
	Grid       Grid              `json:"grid"`
}
