package entities

import "fmt"

func NewPoint(x, y int32) Point {
	return Point{X: x, Y: y}
}

func NewPointp(x, y int32) *Point {
	return &Point{X: x, Y: y}
}

type Point struct {
	X, Y int32
}

func (p *Point) DisplacementFrom(d Point) Point {
	return Point{X: d.X - p.X, Y: d.Y - p.Y}
}

func (p *Point) Move(d Point) {
	p.X += d.X
	p.Y += d.Y
}

func (p *Point) Print() {
	fmt.Printf("\nPoint x: %d, y: %d", p.X, p.Y)
}
