package entities

import (
	"fmt"
	"math"
)

func NewContainer(ul Point, dr Point) Container {
	return Container{UpperLeft: ul, DownRight: dr, width: dr.X - ul.X, heigth: dr.Y - ul.Y}
}

type Container struct {
	width     int32
	heigth    int32
	UpperLeft Point
	DownRight Point
}

func (c *Container) Print() {
	fmt.Printf("Container Width: %d Height: %d", c.Width(), c.Height())
}

// Move the container to a new point position, using up left as origin
func (c *Container) MoveTo(p Point) {
	newDownRightPosition := NewPoint(p.X+c.Width(), p.Y+c.Height())
	c.UpperLeft.Move(c.UpperLeft.DisplacementFrom(p))
	c.DownRight.Move(c.DownRight.DisplacementFrom(newDownRightPosition))
}

func (c *Container) DisplacementFrom(p Point) Point {
	return c.UpperLeft.DisplacementFrom(p)
}

func (c *Container) Move(p Point) {
	c.UpperLeft.Move(p)
	c.DownRight.Move(p)
}

func (c *Container) Padding(p int32) {
	var px, py int32
	if c.Width() > c.Height() {
		py = int32(math.Round(float64(p) * (float64(c.Height()) / float64(c.Width()))))
		px = p
	} else {
		px = int32(math.Round(float64(p) * (float64(c.Height()) / float64(c.Width()))))
		py = p
	}
	c.UpperLeft.X += px
	c.UpperLeft.Y += py
	c.DownRight.X -= px
	c.DownRight.Y -= py
}

func (c *Container) Width() int32 {
	return c.UpperLeft.DisplacementFrom(c.DownRight).X
}

func (c *Container) Height() int32 {
	return c.UpperLeft.DisplacementFrom(c.DownRight).Y
}

func (c *Container) Scale(s float64) {
	c.DownRight = NewPoint(
		int32(float64(c.Width())*s)+c.UpperLeft.X,
		int32(float64(c.Height())*s)+c.UpperLeft.Y,
	)
	c.width = c.Width()
	c.heigth = c.Height()
}

func (c *Container) Center() Point {
	x := (float64(c.Width()) / 2) + float64(c.UpperLeft.X)
	y := (float64(c.Height()) / 2) + float64(c.UpperLeft.Y)
	return NewPoint(int32(x), int32(y))
}
