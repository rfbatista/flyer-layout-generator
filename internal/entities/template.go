package entities

import (
	"strconv"
	"time"
)

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
	Name           string                   `json:"name,omitempty"`
	Width          int32                    `json:"width,omitempty"`
	Height         int32                    `json:"height,omitempty"`
	Distortion     TemplateDistortion       `json:"distortion,omitempty"`
	SlotsPositions []TemplateSlotsPositions `json:"slots_positions,omitempty"`
	SlotsX         int32                    `json:"x,omitempty"`
	MaxSlotsX      int32                    `json:"max_slots_x,omitempty"`
	SlotsY         int32                    `json:"y,omitempty"`
	MaxSlotsY      int32                    `json:"max_slots_y,omitempty"`
	CreatedAt      time.Time                `json:"created_at,omitempty"`
}

const (
	minSlotSize = 50
	maxSlots    = 6
)

func (t *Template) Grids() []Grid {
	var g []Grid
	if true {
		grid, _ := NewGrid(
			WithDefault(t.Width, t.Height),
			WithCells(int32(3), int32(3)),
		)
		g = append(g, *grid)
		return g
	}
	maxXSlots := t.Width / minSlotSize
	if maxXSlots > maxSlots {
		maxXSlots = maxSlots
	}

	maxYSlots := t.Height / minSlotSize
	if maxYSlots > maxSlots {
		maxYSlots = maxSlots
	}

	for x := 1; x <= int(maxXSlots); x++ {
		slotWidth := int(t.Width) / x
		if slotWidth < minSlotSize {
			break
		}
		for y := 1; y <= int(maxYSlots); y++ {
			slotHeight := int(t.Height) / y
			if slotHeight < minSlotSize {
				break
			}
			grid, _ := NewGrid(
				WithDefault(t.Width, t.Height),
				WithCells(int32(x), int32(y)),
			)
			if grid != nil {
				g = append(
					g, *grid,
				)
			}
		}
	}
	return g
}

func (t *Template) MaxSlotsXText() string {
	if t.MaxSlotsX == 0 {
		return ""
	}
	return strconv.FormatInt(int64(t.MaxSlotsX), 10)
}

func (t *Template) MaxSlotsYText() string {
	if t.MaxSlotsY == 0 {
		return ""
	}
	return strconv.FormatInt(int64(t.MaxSlotsY), 10)
}

func (t *Template) WidthS() string {
	return strconv.FormatInt(int64(t.Width), 10)
}

func (t *Template) HeightS() string {
	return strconv.FormatInt(int64(t.Height), 10)
}

func (t *Template) CreatedAtText() string {
	return t.CreatedAt.Format(timeformat)
}

func NewTemplateType(t string) TemplateType {
	switch t {
	case "slots":
		return TemplateSlotsType
	default:
		return TemplateDistortionType
	}
}

func (d *Template) SID() string {
	return strconv.FormatInt(int64(d.ID), 10)
}

func (d *Template) SWidth() string {
	return strconv.FormatInt(int64(d.Width), 10)
}

func (d *Template) SHeigth() string {
	return strconv.FormatInt(int64(d.Height), 10)
}
