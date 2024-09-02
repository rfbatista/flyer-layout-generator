package entities

func NewGridContainer(ul, dr Position) GridContainer {
	return GridContainer{
		UpLeft:    ul,
		DownRight: dr,
		DownLeft:  NewPosition(ul.X, dr.Y),
		UpRight:   NewPosition(dr.X, ul.Y),
	}
}

func NewGridContainerFromPoints(positions []Position) GridContainer {
	upleft := positions[0]
	downright := positions[0]
	for _, p := range positions {
		if p.X <= upleft.X && p.Y <= upleft.Y {
			upleft = p
		}
		if p.X >= downright.X && p.Y >= downright.Y {
			downright = p
		}
	}
	return NewGridContainer(upleft, downright)
}

type GridContainer struct {
	UpLeft    Position
	DownRight Position
	DownLeft  Position
	UpRight   Position
}

func (g *GridContainer) Width() int32 {
	return (g.UpRight.X - g.UpLeft.X) + 1
}

func (g *GridContainer) Height() int32 {
	return (g.DownLeft.Y - g.UpLeft.Y) + 1
}

func (g *GridContainer) ToOrigin() {
	w := g.Width()
	h := g.Height()
	g.UpLeft = NewPosition(0, 0)
	g.UpRight = NewPosition(w-1, 0)
	g.DownLeft = NewPosition(0, h-1)
	g.DownRight = NewPosition(w-1, h-1)
}

func (g *GridContainer) MoveUp() {
	g.UpLeft.MoveUp()
	g.UpRight.MoveUp()
	g.DownLeft.MoveUp()
	g.DownRight.MoveUp()
}

func (g *GridContainer) MoveDown() {
	g.UpLeft.MoveDown()
	g.UpRight.MoveDown()
	g.DownLeft.MoveDown()
	g.DownRight.MoveDown()
}

func (g *GridContainer) MoveLeft() {
	g.UpLeft.MoveLeft()
	g.UpRight.MoveLeft()
	g.DownLeft.MoveLeft()
	g.DownRight.MoveLeft()
}

func (g *GridContainer) MoveRight() {
	g.UpLeft.MoveRight()
	g.UpRight.MoveRight()
	g.DownLeft.MoveRight()
	g.DownRight.MoveRight()
}

func (g *GridContainer) HavePosition(p Position) bool {
	for x := g.UpLeft.X; x <= g.UpRight.X; x++ {
		for y := g.UpLeft.Y; y <= g.DownLeft.Y; y++ {
			if p.X == x && p.Y == y {
				return true
			}
		}
	}
	return false
}

func (g *GridContainer) HavePoint(p Point) bool {
	for x := g.UpLeft.X; x <= g.UpRight.X; x++ {
		for y := g.UpLeft.Y; y <= g.DownLeft.Y; y++ {
			if p.X == x && p.Y == y {
				return true
			}
		}
	}
	return false
}

func (g *GridContainer) ToContainer(cellWidth, cellHeight int32) Container {
	wp := (g.UpRight.X - g.UpLeft.X + 1) * cellWidth
	hp := (g.DownLeft.Y - g.UpLeft.Y + 1) * cellHeight
	return NewContainer(
		NewPoint(g.UpLeft.X*cellWidth, g.UpLeft.Y*cellHeight),
		NewPoint((g.UpLeft.X*cellWidth)+wp, (g.UpLeft.Y*cellHeight)+hp),
	)
}
