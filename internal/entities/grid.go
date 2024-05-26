package entities

import (
	"github.com/google/uuid"
	"math"
)

func NewRegion(x, y, xx, yy int32) *Region {
	return &Region{
		ID:  uuid.New(),
		Xi:  x,
		Yi:  y,
		Xii: xx,
		Yii: yy,
	}
}

type Region struct {
	ID  uuid.UUID `json:"id"`
	Xi  int32     `json:"xi"`
	Xii int32     `json:"xii"`
	Yi  int32     `json:"yi"`
	Yii int32     `json:"yii"`
}

func (r *Region) DistanceToEdge(p Point) int32 {
	smallerDistance := int32(math.Abs(float64(r.Xi - p.X)))
	if smallerDistance < r.Xii-p.X {
		smallerDistance = r.Xii - p.X
	}
	if smallerDistance < r.Yi-p.Y {
		smallerDistance = r.Yi - p.Y
	}
	if smallerDistance < r.Yii-p.Y {
		smallerDistance = r.Yii - p.Y
	}
	return smallerDistance
}

type gridOption struct {
	width, height, padding, pivotX, pivotY int32
}

type GridOption func(options *gridOption) error

func WithDefault(width, height int32) GridOption {
	return func(options *gridOption) error {
		options.width = width
		options.height = height
		return nil
	}
}

func WithPadding(padding int32) GridOption {
	return func(options *gridOption) error {
		options.padding = padding
		return nil
	}
}

func WithPivot(pivotX, pivotY int32) GridOption {
	return func(options *gridOption) error {
		options.pivotX = pivotX
		options.pivotY = pivotY
		return nil
	}
}

func NewGrid(opts ...GridOption) (*Grid, error) {
	var in gridOption
	for _, opt := range opts {
		err := opt(&in)
		if err != nil {
			return nil, err
		}
	}
	restX := math.Mod(float64(in.width), float64(in.pivotX))
	restY := math.Mod(float64(in.height), float64(in.pivotY))
	slotsX := (in.width - int32(restX)) / in.pivotX
	slotsY := (in.height - int32(restY)) / in.pivotY
	if slotsX == 0 {
		slotsX = 1
	}
	if slotsY == 0 {
		slotsY = 1
	}
	var grid Grid
	xi := int32(restX / 2)
	xii := int32(xi + in.pivotX)
	var regions []Region
	for x := int32(0); x < slotsX; x++ {
		yi := int32(restY / 2)
		yii := yi + in.pivotY
		for y := int32(0); y < slotsY; y++ {
			regions = append(regions, *NewRegion(xi, yi, xii, yii))
			yi = yi + in.pivotY
			yii = yii + in.pivotY
		}
		xi = xi + in.pivotX
		xii = xi + in.pivotX
	}
	grid.Regions = regions
	return &grid, nil
}

func NewPoint(x, y int32) *Point {
	return &Point{X: x, Y: y}
}

type Point struct {
	X, Y int32
}

type Grid struct {
	Regions []Region `json:"regions"`
}

func (g *Grid) WhereToSnap(e DesignComponent) (Region, bool) {
	snapToLeft := true
	upleft := NewPoint(e.Xi, e.Yi)
	upright := NewPoint(e.Xii, e.Yi)
	downleft := NewPoint(e.Xi, e.Yii)
	downright := NewPoint(e.Xii, e.Yii)
	smallerDistance := int32(999999)
	nearestRegion := g.Regions[0]
	for _, region := range g.Regions {
		if smallerDistance > region.DistanceToEdge(*upleft) {
			smallerDistance = region.DistanceToEdge(*upleft)
			snapToLeft = true
			nearestRegion = region
		}
		if smallerDistance > region.DistanceToEdge(*upright) {
			smallerDistance = region.DistanceToEdge(*upright)
			snapToLeft = false
			nearestRegion = region
		}
		if smallerDistance > region.DistanceToEdge(*downleft) {
			smallerDistance = region.DistanceToEdge(*downleft)
			snapToLeft = true
			nearestRegion = region
		}
		if smallerDistance > region.DistanceToEdge(*downright) {
			smallerDistance = region.DistanceToEdge(*downright)
			snapToLeft = false
			nearestRegion = region
		}
	}
	return nearestRegion, snapToLeft
}

func (g *Grid) RemoveAllRegionsInThisPosition(xi, yi, xii, yii int32) {
	for _, region := range g.Regions {
		if region.Xi > xii || region.Xii < xi {
			continue
		}
		if region.Yi > yii || region.Yii < yi {
			continue
		}
		g.RemoveRegion(region)
	}
}

func (g *Grid) RemoveRegion(re Region) {
	g.Regions = filterRegion(g.Regions, func(r Region) bool {
		if r.ID == re.ID {
			return false
		}
		return true
	})

}

func filterRegion(ss []Region, test func(r Region) bool) (ret []Region) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
