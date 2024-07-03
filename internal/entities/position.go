package entities

func NewPosition(x, y int32) Position {
	return Position{X: x, Y: y}
}

type Position struct {
	X, Y int32
}

func (p *Position) MoveUp() {
	p.Y -= 1
}

func (p *Position) MoveDown() {
	p.Y += 1
}

func (p *Position) MoveLeft() {
	p.X -= 1
}

func (p *Position) MoveRight() {
	p.X += 1
}
