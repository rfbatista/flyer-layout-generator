package entities

type Layout struct {
	ID         int32
	DesignID   int32
	Width      int32             `json:"width,omitempty"`
	Height     int32             `json:"height,omitempty"`
	Components []DesignComponent `json:"components,omitempty"`
	Template   Template          `json:"template,omitempty"`
	Grid       Grid              `json:"grid"`
}
