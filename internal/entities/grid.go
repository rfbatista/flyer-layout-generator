package entities

import (
	"errors"
	"fmt"
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
	var grid Grid
	grid.SlotsX = in.CellsX
	grid.SlotsY = in.CellsY
	sizeX := int32(float64(in.width) / float64(in.CellsX))
	sizeY := int32(float64(in.height) / float64(in.CellsY))
	xi := int32(0)
	xii := sizeX
	grid.position = make([][]GridCell, in.CellsX)
	for x := int32(0); x < in.CellsX; x++ {
		grid.position[x] = make([]GridCell, in.CellsY)
	}
	for x := int32(0); x < in.CellsX; x++ {
		yi := int32(0)
		yii := sizeY
		for y := int32(0); y < in.CellsY; y++ {
			cell := NewCell(xi, yi, xii, yii)
			cell.SetPosition(x, y)
			grid.position[x][y] = *cell
			yi = yi + sizeY
			yii = yi + sizeY
		}
		xi = xi + sizeX
		xii = xi + sizeX
	}
	grid.width = in.width
	grid.height = in.height
	grid.slotWidth = sizeX
	grid.slotHeight = sizeY
	return &grid, nil
}

type Grid struct {
	AllCells   []GridCell `json:"regions"`
	position   [][]GridCell
	width      int32
	height     int32
	slotWidth  int32
	slotHeight int32
	SlotsX     int32
	SlotsY     int32
}

func (g *Grid) PrintGrid() {
	// Determine the grid dimensions
	fmt.Println("\n")
	gridWidth := len(g.position)
	if gridWidth == 0 {
		fmt.Println("Grid is empty")
		return
	}
	gridHeight := len(g.position[0])
	fmt.Printf(
		"\n X: %d Y: %d height: %d width: %d \n",
		g.SlotsX,
		g.SlotsY,
		g.slotHeight,
		g.slotWidth,
	)
	// Iterate through each cell and print the grid
	for y := 0; y < gridHeight; y++ {
		// Print horizontal separator
		if y == 0 {
			fmt.Print("+")
			for x := 0; x < gridWidth; x++ {
				fmt.Print("---+")
			}
			fmt.Println()
		}

		// Print cell content
		fmt.Print("|")
		for x := 0; x < gridWidth; x++ {
			cell := g.position[x][y]
			if cell.IsOcupied() {
				fmt.Print(" X |")
			} else {
				fmt.Print("   |")
			}
		}
		fmt.Println()

		// Print horizontal separator
		fmt.Print("+")
		for x := 0; x < gridWidth; x++ {
			fmt.Print("---+")
		}
		fmt.Println("\n")
	}
}

func (g *Grid) Cells() []GridCell {
	var cells []GridCell
	for _, c := range g.position {
		for _, c1 := range c {
			cells = append(cells, c1)
		}
	}
	return cells
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
	return findOverlappingRegions(e, g.Cells())
}

func (g *Grid) WhereToSnap(e DesignComponent) (GridCell, bool) {
	snapToLeft := true
	upleft := NewPointp(e.Xi, e.Yi)
	upright := NewPointp(e.Xii, e.Yi)
	downleft := NewPointp(e.Xi, e.Yii)
	downright := NewPointp(e.Xii, e.Yii)
	smallerDistance := int32(999999)
	nearestRegion := g.Cells()[0]
	for _, region := range g.Cells() {
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
	for idx := range g.Cells() {
		if g.Cells()[idx].IsIn(p) {
			return &g.Cells()[idx]
		}
	}
	return nil
}

// Find in which cell the id is present
func (g *Grid) WhereIsId(id int32) *GridCell {
	for idx := range g.Cells() {
		if g.Cells()[idx].IsIdIn(id) {
			return &g.Cells()[idx]
		}
	}
	return nil
}

// Calculate the space in cells number that a container need to fit in the grid
func (g *Grid) ContainerToFit(c Container) Container {
	w := g.Cells()[0].Width()
	h := g.Cells()[0].Height()
	x := math.Ceil(float64(c.Width()) / float64(w))
	y := math.Ceil(float64(c.Height()) / float64(h))
	return NewContainer(NewPoint(0, 0), NewPoint(int32(x)*w, int32(y)*h))
}

// Transform a list of positions in a container
func (g *Grid) PositionsToContainer(points []Point) Container {
	var c Container
	var cells []*GridCell
	for _, p := range points {
		if p.X < 0 || p.X >= g.SlotsX || p.Y < 0 || p.Y >= g.SlotsY {
			continue
		}
		cells = append(cells, &g.position[p.X][p.Y])
	}
	if len(cells) == 0 {
		return c
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

// Find positions that fit the provided container from the provided cell pivot
func (g *Grid) FindPositionsToFitBasedOnPivot(p Point, c Container) ([]Point, error) {
	var cc *GridCell
	var points []Point
	if len(g.position) <= int(p.X) && p.X < 0 {
		return points, errors.New("x point provided is out of boundaries")
	}
	if len(g.position[0]) <= int(p.Y) && p.Y < 0 {
		return points, errors.New("y point provided is out of boundaries")
	}
	x, y, found := g.FindSpace(p, c)
	if !found {
		return points, errors.New("no position was found")
	}
	cc = &g.position[x][y]
	if cc == nil {
		return points, errors.New("nenhuma celula definida com a posicicao especificada")
	}
	points = append(points, NewPoint(x, y))
	c.MoveTo(cc.UpLeft())
	xcellsToMove := math.Ceil(float64(c.Width())/float64(g.slotWidth)) - 1
	xPos := cc.Position().X
	for x := int32(1); x <= int32(xcellsToMove); x++ {
		xPos += x
		points = append(points, NewPoint(xPos, cc.positionY))
	}

	cellsToMove := math.Ceil(float64(c.Height())/float64(g.slotHeight)) - 1
	yPos := cc.Position().Y
	for y := int32(1); y <= int32(cellsToMove); y++ {
		yPos += y
		points = append(points, NewPoint(cc.positionX, yPos))
	}
	return points, nil
}

// Find free positions that fit the provided container from the provided cell pivot
func (g *Grid) FindFreePositionsToFitBasedOnPivot(p Point, c Container) ([]Point, error) {
	var cc *GridCell
	var points []Point
	if len(g.position) <= int(p.X) && p.X < 0 {
		return points, errors.New("x point provided is out of boundaries")
	}
	if len(g.position[0]) <= int(p.Y) && p.Y < 0 {
		return points, errors.New("y point provided is out of boundaries")
	}
	x, y, found := g.FindSpace(p, c)
	if x >= 0 && y >= 0 {
		cc = &g.position[x][y]
	}
	if !found || cc.IsOcupied() || !g.Fits(cc.Position().X, cc.Position().Y, c) {
		ccc := g.FindFreeCellByReadingOrder()
		if ccc == nil {
			return points, errors.New("celula ja ocupada e nenhuma vazio foi identificada")
		}
		x, y, found = g.FindSpace(ccc.UpLeft(), c)
		if !found {
			return points, errors.New("no position was found")
		}
		cc = ccc
	}
	if cc == nil {
		return points, errors.New("nenhuma celula definida com a posicicao especificada")
	}
	points = append(points, NewPoint(x, y))
	c.MoveTo(cc.UpLeft())
	xcellsToMove := math.Ceil(float64(c.Width())/float64(g.slotWidth)) - 1
	xPos := cc.Position().X
	for x := int32(1); x <= int32(xcellsToMove); x++ {
		xPos += x
		points = append(points, NewPoint(xPos, cc.positionY))
	}
	cellsToMove := math.Ceil(float64(c.Height())/float64(g.slotHeight)) - 1
	yPos := cc.Position().Y
	for y := int32(1); y <= int32(cellsToMove); y++ {
		yPos += y
		points = append(points, NewPoint(cc.positionX, yPos))
	}
	return points, nil
}

// Find free cell using reading order
func (g *Grid) FindFreeCellByReadingOrder() *GridCell {
	for y := int32(0); y < g.SlotsY; y++ {
		for x := int32(0); x < g.SlotsX; x++ {
			pos := g.position[x][y]
			if pos.IsOcupied() {
				continue
			}
			return &pos
		}
	}
	return nil
}

// FindSpace searches for a space in the grid to fit the container
func (g *Grid) FindSpace(point Point, container Container) (int32, int32, bool) {
	pivotX := point.X
	pivotY := point.Y

	if g.Fits(pivotX, pivotY, container) {
		return pivotX, pivotY, true
	}
	// Check from the pivot position to the right corner
	for x := pivotX; x < g.SlotsX; x++ {
		y := pivotY
		if g.Fits(x, y, container) && !g.position[x][y].IsOcupied() {
			return x, y, true
		}
	}
	// Check from the pivot position to the left corner
	for x := pivotX; x >= 0; x-- {
		y := pivotY
		if g.Fits(x, y, container) && !g.position[x][y].IsOcupied() {
			return x, y, true
		}
	}
	// Check from the pivot position to up
	for y := pivotY; y >= 0; y-- {
		x := pivotX
		if g.Fits(x, y, container) && !g.position[x][y].IsOcupied() {
			return x, y, true
		}
	}
	// Check from the pivot position to down
	for y := pivotY; y < g.SlotsY; y++ {
		x := pivotX
		if g.Fits(x, y, container) && !g.position[x][y].IsOcupied() {
			return x, y, true
		}
	}

	return -1, -1, false
}

// fits checks if the container fits at position (x, y) in the grid
func (g *Grid) Fits(x, y int32, container Container) bool {
	// Check if the container would extend beyond the grid boundaries
	sizex := (x * g.slotWidth) + container.Width()
	sizey := (y * g.slotHeight) + container.Height()
	if sizex > g.Width() ||
		sizey > g.Height() {
		return false
	}
	return true
}

func (g *Grid) IsSpaceOcupied(x, y int32, container Container) bool {
	// Check if the container overlaps with any occupied cells
	for i := int32(0); i < container.Height(); i++ {
		for j := int32(0); j < container.Width(); j++ {
			if g.position[y+i][x+j].IsOcupied() {
				return false
			}
		}
	}
	return true
}

// Occupy specified cell
func (g *Grid) OcupyCell(c GridCell, id int32) {
	if len(g.position) > int(c.positionX) && len(g.position[c.positionX]) > int(c.positionY) {
		g.position[c.positionX][c.positionY].Ocupy(id)
	}
}

// Occupy cells by container
func (g *Grid) OcupyByPositionList(points []Point, id int32) []*GridCell {
	var cells []*GridCell
	for _, p := range points {
		cell := g.OcupyByPosition(p, id)
		cells = append(cells, cell)
	}
	return cells
}

// Occupy cells by container
func (g *Grid) OcupyByPosition(p Point, id int32) *GridCell {
	if len(g.position) > int(p.X) && len(g.position[p.X]) > int(p.Y) {
		g.position[p.X][p.Y].Ocupy(id)
		return nil
	}
	return nil
}

// Occupy cells by container
func (g *Grid) OcupyWithContainer(c Container, id int32) bool {
	for idx := range g.Cells() {
		if g.Cells()[idx].InstersectWithContainer(c) {
			if !g.OcupyCellAndCheck(g.Cells()[idx], id) {
				return false
			}
		}
	}
	return true
}

// Occupy specified cell
func (g *Grid) OcupyCellAndCheck(c GridCell, id int32) bool {
	for idx := range g.Cells() {
		if g.Cells()[idx].ID.String() == c.ID.String() {
			g.OcupyCell(c, id)
			return true
		}
	}
	return false
}

func (g *Grid) RemoveAllRegionsInThisPosition(xi, yi, xii, yii int32) {
	for _, region := range g.Cells() {
		if region.Xi > xii || region.Xii < xi {
			continue
		}
		if region.Yi > yii || region.Yii < yi {
			continue
		}
		g.RemoveRegion(region)
	}
}

func (g *Grid) GetSurroundFreeCells(p Point) []Point {
	var points []Point
	if p.X-1 > 0 && !g.IsPositionOcupied(NewPosition(p.X-1, p.Y)) {
		points = append(points, NewPoint(p.X-1, p.Y))
	}
	if p.Y-1 > 0 && !g.IsPositionOcupied(NewPosition(p.X, p.Y-1)) {
		points = append(points, NewPoint(p.X, p.Y-1))
	}
	if p.X+1 < g.SlotsX && !g.IsPositionOcupied(NewPosition(p.X+1, p.Y)) {
		points = append(points, NewPoint(p.X+1, p.Y))
	}
	if p.Y+1 > g.SlotsY && !g.IsPositionOcupied(NewPosition(p.X, p.Y+1)) {
		points = append(points, NewPoint(p.X, p.Y+1))
	}
	return points
}

// Check if the position is is ocupied
func (g *Grid) IsPositionOcupied(p Position) bool {
	return g.position[p.X][p.X].IsOcupied()
}

// Check if the component have space to grow
func (g *Grid) FindSpaceToGrow(p Point, c Container, id int32) (*Container, error) {
	var nCont *GridContainer
	var pos []Point
	scale := float64(1.0)
	for {
		co := NewContainer(c.UpperLeft, c.DownRight)
		co.Scale(scale)
		nnCont := g.ContainerToGridContainer(co)
		cont, found, err := g.FindPositionToFitGridContainer(
			p,
			nnCont,
			id,
		)
		if err != nil || !found {
			if nCont != nil {
				g.OcupyByPositionList(pos, id)
				c := nCont.ToContainer(g.slotWidth, g.slotHeight)
				return &c, nil
			}
			return nil, err
		}
		nCont = &cont
		scale += float64(0.1)
	}
}

func (g *Grid) FindPositionToFitGridContainer(
	p Point,
	c GridContainer,
	id int32,
) (GridContainer, bool, error) {
	// Boundary checks
	if p.X < 0 || p.X >= int32(len(g.position)) {
		return c, false, errors.New("x point provided is out of boundaries")
	}
	if p.Y < 0 || p.Y >= int32(len(g.position[0])) {
		return c, false, errors.New("y point provided is out of boundaries")
	}

	walkInX := g.SlotsX - c.Width()
	walkInY := g.SlotsY - c.Height()
	if walkInX < 0 || walkInY < 0 {
		return c, false, errors.New("grid container do not fit in this grid")
	}
	goDown := true
	c.ToOrigin()
	for x := 0; x <= int(walkInX); x++ {
		if !g.CheckGridContainerColision(c, id) && c.HavePoint(p) {
			return c, true, nil
		}
		for y := 0; y <= int(walkInY); y++ {
			if !g.CheckGridContainerColision(c, id) && c.HavePoint(p) {
				return c, true, nil
			}
			if goDown && c.DownRight.Y < g.SlotsY-1 {
				c.MoveDown()
			}
			if !goDown && c.UpRight.Y > 0 {
				c.MoveUp()
			}
			if !g.CheckGridContainerColision(c, id) && c.HavePoint(p) {
				return c, true, nil
			}
		}
		if c.UpRight.X == g.SlotsX-1 {
			continue
		}
		c.MoveRight()
		goDown = !goDown
	}
	return c, false, errors.New("position not found to fit container")
}

func (g *Grid) HaveColisionInList(points []Position) bool {
	for _, p := range points {
		if g.IsPositionOcupied(p) {
			return true
		}
	}
	return false
}

func (g *Grid) ContainerToPositions(c Container) []Position {
	var points []Position
	xcellsToMove := int32(math.Ceil(float64(c.Width())/float64(g.slotWidth))) - 1
	ycellsToMove := int32(math.Ceil(float64(c.Height())/float64(g.slotHeight))) - 1
	for dx := int32(0); dx <= xcellsToMove; dx++ {
		for dy := int32(0); dy <= ycellsToMove; dy++ {
			points = append(points, NewPosition(dx, dy))
		}
	}
	return points
}

func (g *Grid) ContainerToGridContainer(c Container) GridContainer {
	p := g.ContainerToPositions(c)
	return NewGridContainerFromPoints(p)
}

func (g *Grid) CheckGridContainerColision(c GridContainer, id int32) bool {
	for x := c.UpLeft.X; x <= c.UpRight.X; x++ {
		for y := c.UpLeft.Y; y <= c.DownLeft.Y; y++ {
			if x < 0 || x > int32(g.SlotsX-1) {
				return true
			}
			if y < 0 || y > int32(g.SlotsY-1) {
				return true
			}
			cellc := g.position[x][y]
			if !cellc.IsOnlyOcupiedBy(id) {
				return true
			}
		}
	}
	return false
}

func (g *Grid) RemoveRegion(re GridCell) {
}

func filterRegion(ss []GridCell, test func(r GridCell) bool) (ret []GridCell) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
