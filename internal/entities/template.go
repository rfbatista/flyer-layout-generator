package entities

import (
	"errors"
	"strconv"
	"time"
)

type TemplateType string

const (
	TemplateSlotsType      TemplateType = "slots"
	TemplateDistortionType TemplateType = "distortion"
	TemplateTypeAdaptation TemplateType = "adaptation"
	TemplateUnknownType    TemplateType = "unknown"
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
	maxSlots    = 8
)

type slot struct {
	x int
	y int
}

func (t *Template) Grids() []Grid {
	var g []Grid
	maxXSlots := t.Width / minSlotSize
	if maxXSlots > maxSlots {
		maxXSlots = maxSlots
	}

	maxYSlots := t.Height / minSlotSize
	if maxYSlots > maxSlots {
		maxYSlots = maxSlots
	}

	// for x := 2; x <= 6; x++ {
	// 	for y := 2; y <= 6; y++ {
	// 		grid, _ := t.CreateGrid(x, y)
	// 		if grid != nil {
	// 			g = append(
	// 				g, *grid,
	// 			)
	// 		}
	// 	}
	// }
	grids := []slot{
		{x: 2, y: 6},
		{x: 4, y: 6},
		{x: 4, y: 1},
		{x: 4, y: 2},
		{x: 4, y: 3},
		{x: 4, y: 4},
		{x: 1, y: 4},
		{x: 2, y: 4},
		{x: 3, y: 4},
		{x: 2, y: 6},
		{x: 3, y: 6},
		{x: 4, y: 6},
		{x: 5, y: 6},
		{x: 6, y: 6},
		{x: 6, y: 2},
		{x: 6, y: 3},
		{x: 6, y: 4},
		{x: 6, y: 5},
		{x: 6, y: 1},
	}

	for _, gr := range grids {
		grid, _ := t.CreateGrid(gr.x, gr.y)
		if grid != nil {
			g = append(
				g, *grid,
			)
		}
	}

	return g
}

func (t *Template) CreateGrid(x, y int) (*Grid, error) {
	slotWidth := float64(t.Width) / float64(x)
	if slotWidth < float64(minSlotSize) {
		return nil, errors.New("slow width < minimum slot size")
	}
	slotHeight := float64(t.Height) / float64(y)
	if slotHeight < float64(minSlotSize) {
		return nil, errors.New("slow height < minimum slot size")
	}
	grid, _ := NewGrid(
		WithDefault(t.Width, t.Height),
		WithPivot(int32(slotWidth), int32(slotHeight)),
		WithCells(int32(x), int32(y)),
	)
	return grid, nil
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
