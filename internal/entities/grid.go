package entities

import (
	"math"
)

type gridOption struct {
	width, height, padding, pivotX, pivotY, CellsX, CellsY int32
}

type GridOption func(options *gridOption) error

func WithDefault(width, height int32) GridOption {
	return func(options *gridOption) error {
		options.width = width
		options.height = height
		return nil
	}
}

func WithCells(cellsX, cellsY int32) GridOption {
	return func(options *gridOption) error {
		options.CellsX = cellsX
		options.CellsY = cellsY
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
	if in.CellsY > 0 && in.CellsX > 0 {
		var grid Grid
		grid.SlotsX = in.CellsX
		grid.SlotsY = in.CellsY
		sizeX := int32(float64(in.width) / float64(in.CellsX))
		sizeY := int32(float64(in.height) / float64(in.CellsY))
		var regions []GridCell
		xi := int32(0)
		xii := sizeX
		for x := int32(0); x < in.CellsX; x++ {
			yi := int32(0)
			yii := sizeY
			for y := int32(0); y < in.CellsY; y++ {
				regions = append(regions, *NewCell(xi, yi, xii, yii))
				yi = yi + sizeY
				yii = yi + sizeY
			}
			xi = xi + sizeX
			xii = xi + sizeX
		}
		grid.Cells = regions
		return &grid, nil
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
	var regions []GridCell
	for x := int32(0); x < slotsX; x++ {
		yi := int32(restY / 2)
		yii := yi + in.pivotY
		for y := int32(0); y < slotsY; y++ {
			regions = append(regions, *NewCell(xi, yi, xii-1, yii-1))
			yi = yi + in.pivotY
			yii = yii + in.pivotY
		}
		xi = xi + in.pivotX
		xii = xi + in.pivotX
	}
	grid.Cells = regions
	return &grid, nil
}

type Grid struct {
	Cells  []GridCell `json:"regions"`
	SlotsX int32
	SlotsY int32
}

type OverlapResult struct {
	Region  GridCell `json:"region"`
	Overlap int32    `json:"overlap"`
}

func findOverlappingRegions(rect DesignComponent, regions []GridCell) []OverlapResult {
	var overlappingRegions []OverlapResult
	for _, region := range regions {
		if overlap, area := isOverlap(rect, region); overlap {
			overlappingRegions = append(
				overlappingRegions,
				OverlapResult{Region: region, Overlap: area},
			)
		}
	}
	return overlappingRegions
}

func isOverlap(rect DesignComponent, region GridCell) (bool, int32) {
	// Calculate the overlapping area if the rectangle overlaps with the region
	xOverlap := min(rect.Xii, region.Xii) - max(rect.Xi, region.Xi)
	yOverlap := min(rect.Yii, region.Yii) - max(rect.Yi, region.Yi)
	if xOverlap > 0 && yOverlap > 0 {
		return true, xOverlap * yOverlap
	}
	return false, 0
}

func (g *Grid) FindOverlappingRegions(e DesignComponent) []OverlapResult {
	return findOverlappingRegions(e, g.Cells)
}

func (g *Grid) WhereToSnap(e DesignComponent) (GridCell, bool) {
	snapToLeft := true
	upleft := NewPointp(e.Xi, e.Yi)
	upright := NewPointp(e.Xii, e.Yi)
	downleft := NewPointp(e.Xi, e.Yii)
	downright := NewPointp(e.Xii, e.Yii)
	smallerDistance := int32(999999)
	nearestRegion := g.Cells[0]
	for _, region := range g.Cells {
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

// Find in which cell the component is
func (g *Grid) WhereIsPoint(p Point) *GridCell {
	for idx := range g.Cells {
		if g.Cells[idx].IsIn(p) {
			return &g.Cells[idx]
		}
	}
	return nil
}

// Find in which cell the id is present
func (g *Grid) WhereIsId(id int32) *GridCell {
	for idx := range g.Cells {
		if g.Cells[idx].IsIdIn(id) {
			return &g.Cells[idx]
		}
	}
	return nil
}

// Find in which cell the component is
func (g *Grid) OcupyCell(c GridCell, id int32) {
	for idx := range g.Cells {
		if g.Cells[idx].ID.String() == c.ID.String() {
			g.Cells[idx].Ocupy(id)
		}
	}
}

func (g *Grid) RemoveAllRegionsInThisPosition(xi, yi, xii, yii int32) {
	for _, region := range g.Cells {
		if region.Xi > xii || region.Xii < xi {
			continue
		}
		if region.Yi > yii || region.Yii < yi {
			continue
		}
		g.RemoveRegion(region)
	}
}

func (g *Grid) RemoveRegion(re GridCell) {
	g.Cells = filterRegion(g.Cells, func(r GridCell) bool {
		if r.ID == re.ID {
			return false
		}
		return true
	})
}

func filterRegion(ss []GridCell, test func(r GridCell) bool) (ret []GridCell) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
