package entities

func NewContainer(ul Point, dr Point) Container {
	return Container{UpperLeft: ul, DownRight: dr}
}

type Container struct {
	UpperLeft Point
	DownRight Point
}

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

func (c *Container) Width() int32 {
	return c.UpperLeft.DisplacementFrom(c.DownRight).X
}

func (c *Container) Height() int32 {
	return c.UpperLeft.DisplacementFrom(c.DownRight).Y
}
