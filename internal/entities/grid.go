package entities

import (
	"errors"
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
		grid.position = make([][]*GridCell, in.CellsX)
		for x := int32(0); x < in.CellsX; x++ {
			grid.position[x] = make([]*GridCell, in.CellsY)
		}
		for x := int32(0); x < in.CellsX; x++ {
			yi := int32(0)
			yii := sizeY
			for y := int32(0); y < in.CellsY; y++ {
				cell := NewCell(xi, yi, xii, yii)
				cell.SetPosition(x, y)
				regions = append(regions, *cell)
				grid.position[x][y] = cell
				yi = yi + sizeY
				yii = yi + sizeY
			}
			xi = xi + sizeX
			xii = xi + sizeX
		}
		grid.Cells = regions
		grid.width = in.width
		grid.height = in.height
		grid.slotWidth = sizeX
		grid.slotHeight = sizeY
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
	Cells      []GridCell `json:"regions"`
	position   [][]*GridCell
	width      int32
	height     int32
	slotWidth  int32
	slotHeight int32
	SlotsX     int32
	SlotsY     int32
}

func (g *Grid) Width() int32 {
	return g.width
}

func (g *Grid) Height() int32 {
	return g.height
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

// Calculate the space in cells number that a container need to fit in the grid
func (g *Grid) ContainerToFit(c Container) Container {
	w := g.Cells[0].Width()
	h := g.Cells[0].Height()
	x := math.Ceil(float64(c.Width()) / float64(w))
	y := math.Ceil(float64(c.Height()) / float64(h))
	return NewContainer(NewPoint(0, 0), NewPoint(int32(x)*w, int32(y)*h))
}

// Transform a list of positions in a container
func (g *Grid) PositionsToContainer(points []Point) Container {
	var c Container
	var cells []*GridCell
	for _, p := range points {
		cells = append(cells, g.position[p.X][p.Y])
	}
	xi := cells[0].UpLeft().X
	xii := cells[0].DownRigth().X
	yi := cells[0].UpLeft().Y
	yii := cells[0].DownRigth().Y

	for _, cell := range cells {
		if xi > cell.UpLeft().X {
			xi = cell.UpLeft().X
		}
		if xii < cell.DownRigth().X {
			xii = cell.DownRigth().X
		}
		if yi > cell.UpLeft().Y {
			yi = cell.UpLeft().Y
		}
		if yii < cell.DownRigth().Y {
			yii = cell.DownRigth().Y
		}
	}
	c.UpperLeft = NewPoint(int32(xi), int32(yi))
	c.DownRight = NewPoint(int32(xii), int32(yii))
	return c
}

// Find positions that fit the provided container from the provided cell pivot uuid
func (g *Grid) FindPositionsToFitBasedOnPivot(p Point, c Container) ([]Point, error) {
	var cc *GridCell
	var points []Point
	if len(g.position) <= int(p.X) && p.X < 0 {
		return points, errors.New("x point provided is out of boundaries")
	}
	if len(g.position[0]) <= int(p.Y) && p.Y < 0 {
		return points, errors.New("y point provided is out of boundaries")
	}
	cc = g.position[p.X][p.Y]
	if cc == nil {
		return points, errors.New("nenhuma celula definida com a posicicao especificada")
	}
	c.MoveTo(cc.UpLeft())
	// TODO: Calcular a realocação com base nas posicoes das celulas
	if cc.UpLeft().X+c.Width() > g.Width() {
		diff := cc.UpLeft().X + c.Width() - g.Width()
		cellsToMove := math.Ceil(float64(diff) / float64(g.slotWidth))
		xPos := cc.Position().X
		for x := int32(0); x <= int32(cellsToMove); x++ {
			xPos -= x
			if xPos < 0 {
				return points, errors.New("a position needed is out of the grid boundaries")
			}
			points = append(points, NewPoint(xPos, cc.positionY))
		}
	}

	if cc.UpLeft().Y+c.Height() > g.Height() {
		diff := cc.UpLeft().Y + c.Height() - g.Height()
		cellsToMove := math.Ceil(float64(diff) / float64(g.slotHeight))
		yPos := cc.Position().Y
		for y := int32(0); y <= int32(cellsToMove); y++ {
			yPos -= y
			if yPos < 0 {
				return points, errors.New("a position needed is out of the grid boundaries")
			}
			points = append(points, NewPoint(yPos, cc.positionY))
		}
	}
	return points, nil
}

// FindSpace searches for a space in the grid to fit the container
func (g *Grid) FindSpace(container Container) (int32, int32, bool) {
	gridWidth := g.Width()
	gridHeight := g.Height()
	for y := int32(0); y <= gridHeight-container.Height(); y++ {
		for x := int32(0); x <= gridWidth-container.Width(); x++ {
			if g.Fits(x, y, container) {
				return x, y, true
			}
		}
	}
	return -1, -1, false
}

// fits checks if the container fits at position (x, y) in the grid
func (g *Grid) Fits(x, y int32, container Container) bool {
	// Check if the container would extend beyond the grid boundaries
	if x+container.Width() > g.Width() || y+container.Height() > g.Height() {
		return false
	}

	// Check if the container overlaps with any occupied cells
	for i := int32(0); i < container.Height(); i++ {
		for j := int32(0); j < container.Width(); j++ {
			if g.position[y+i][x+j] != nil {
				return false
			}
		}
	}
	return true
}

// Occupy specified cell
func (g *Grid) OcupyCell(c GridCell, id int32) {
	for idx := range g.Cells {
		if g.Cells[idx].ID.String() == c.ID.String() {
			g.Cells[idx].Ocupy(id)
		}
	}
}

// Occupy cells by container
func (g *Grid) OcupyByPosition(p Point, id int32) *GridCell {
	cell := g.position[p.X][p.Y]
	if cell != nil {
		cell.Ocupy(id)
	}
	return cell
}

// Occupy cells by container
func (g *Grid) OcupyWithContainer(c Container, id int32) bool {
	for idx := range g.Cells {
		if g.Cells[idx].InstersectWithContainer(c) {
			if !g.OcupyCellAndCheck(g.Cells[idx], id) {
				return false
			}
		}
	}
	return true
}

// Occupy specified cell
func (g *Grid) OcupyCellAndCheck(c GridCell, id int32) bool {
	for idx := range g.Cells {
		if g.Cells[idx].ID.String() == c.ID.String() {
			g.Cells[idx].Ocupy(id)
			return true
		}
	}
	return false
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
