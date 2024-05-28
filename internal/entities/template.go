package entities

type TemplateType string

const (
	TemplateSlotsType      TemplateType = "slots"
	TemplateDistortionType TemplateType = "distortion"
)

func (s TemplateType) String() string {
	return string(s)
}

type TemplateSlotsPositions struct {
	Xi     int32
	Yi     int32
	Width  int32
	Height int32
}

type TemplateDistortion struct {
	X int32 `json:"x,omitempty"`
	Y int32 `json:"y,omitempty"`
}

type Template struct {
	ID             int32                    `json:"id,omitempty"`
	Type           TemplateType             `json:"type,omitempty"`
	Width          int32                    `json:"width,omitempty"`
	Height         int32                    `json:"height,omitempty"`
	Distortion     TemplateDistortion       `json:"distortion,omitempty"`
	SlotsPositions []TemplateSlotsPositions `json:"slots_positions,omitempty"`
	SlotsX         int32                    `json:"x,omitempty"`
	SlotsY         int32                    `json:"y,omitempty"`
}

func NewTemplateType(t string) TemplateType {
	switch t {
	case "slots":
		return TemplateSlotsType
	default:
		return TemplateDistortionType
	}
}
