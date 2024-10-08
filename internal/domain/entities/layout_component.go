package entities

import "fmt"

type LayoutComponent struct {
	ID             int32           `json:"id,omitempty"`
	DesignID       int32           `json:"design_id,omitempty"`
	Elements       []LayoutElement `json:"elements,omitempty"`
	FWidth         int32           `json:"width,omitempty"`
	FHeight        int32           `json:"height,omitempty"`
	Color          string          `json:"color,omitempty"`
	Type           string          `json:"type,omitempty"`
	Xsnaped        bool            `json:"xsnaped,omitempty"`
	Ysnaped        bool            `json:"ysnaped,omitempty"`
	LeftGap        Position        `json:"left_gap,omitempty"`
	RightGap       Position        `json:"right_gap,omitempty"`
	UpGap          Position        `json:"up_gap,omitempty"`
	DownGap        Position        `json:"down_gap,omitempty"`
	InnerContainer Container
	OuterContainer Container
	GridContainer  Container
	Priority       int32
	Positions      []Position
	Pivot          Position
}

func (d *LayoutComponent) Width() int32 {
	return d.InnerContainer.Width()
}

func (d *LayoutComponent) Height() int32 {
	return d.InnerContainer.Height()
}

func (d *LayoutComponent) MoveTo(p Point) {
	displacement := d.InnerContainer.DisplacementFrom(p)
	d.InnerContainer.Move(displacement)
	d.OuterContainer.Move(displacement)
	for idx := range d.Elements {
		d.Elements[idx].Move(displacement)
	}
}

func (d *LayoutComponent) Center() Point {
	return d.InnerContainer.Center()
}

func (d *LayoutComponent) UpLeft() Point {
	return d.InnerContainer.UpperLeft
}

func (d *LayoutComponent) DownRight() Point {
	return d.InnerContainer.DownRight
}

func (d *LayoutComponent) OrderPriority() int32 {
	if d.Priority != 0 {
		return d.Priority
	}
	switch d.Type {
	case "logotipo_marca":
		return 1
	case "logotipo_produto":
		return 2
	case "texto_cta":
		return 4
	case "oferta":
		return 5
	case "texto":
		return 6
	case "icone":
		return 7
	}
	return 6
}

func (d *LayoutComponent) IsBackground() bool {
	return d.Type == "background" || d.Type == "plano-de-fundo"
}

func (d *LayoutComponent) CenterInContainer(r Container) {
	xi := r.UpperLeft.X
	yi := r.UpperLeft.Y
	if r.Width() > d.Width() {
		xi = xi + ((r.Width() - d.Width()) / 2)
	}
	if r.Height() > d.FHeight {
		yi = yi + ((r.Height() - d.Height()) / 2)
	}
	d.MoveTo(NewPoint(xi, yi))
}

func (d *LayoutComponent) ApplyPadding(p int32) {
	d.InnerContainer.Padding(p)
	d.OuterContainer.Padding(p)
	for i := range d.Elements {
		d.Elements[i].InnerContainer.Padding(p)
		d.Elements[i].OuterContainer.Padding(p)
	}
}

func (d *LayoutComponent) ScaleToFitInSize(w, h int32) {
	scaleFactor := calculateScaleFactor(
		float64(d.InnerContainer.Width()),
		float64(d.InnerContainer.Height()),
		float64(w),
		float64(h),
	)
	origin := d.InnerContainer.UpperLeft
	d.InnerContainer.Scale(scaleFactor)
	c := d.OuterContainer.UpperLeft
	xdToOrigin := float64(c.X - origin.X)
	p := NewPoint(
		int32(xdToOrigin*float64(scaleFactor)),
		int32(float64(c.Y-origin.Y)*float64(scaleFactor)),
	)
	d.OuterContainer.MoveTo(p)
	d.OuterContainer.Scale(scaleFactor)
	for i := range d.Elements {
		innerTo := newPosition(scaleFactor, origin, d.Elements[i].InnerContainer.UpperLeft)
		d.Elements[i].InnerContainer.MoveTo(innerTo)
		outerTo := newPosition(
			scaleFactor,
			d.Elements[i].InnerContainer.UpperLeft,
			d.Elements[i].OuterContainer.UpperLeft,
		)
		d.Elements[i].OuterContainer.MoveTo(outerTo)
		d.Elements[i].Scale(scaleFactor)
	}
	d.FWidth = d.InnerContainer.Width()
	d.FHeight = d.InnerContainer.Height()
}

func (d *LayoutComponent) ScaleToFillInSize(w, h int32) {
	scaleFactor := calculateGreaterScaleFactor(
		float64(d.InnerContainer.Width()),
		float64(d.InnerContainer.Height()),
		float64(w),
		float64(h),
	)
	origin := d.InnerContainer.UpperLeft
	d.InnerContainer.Scale(scaleFactor)
	c := d.OuterContainer.UpperLeft
	xdToOrigin := float64(c.X - origin.X)
	p := NewPoint(
		int32(xdToOrigin*float64(scaleFactor)),
		int32(float64(c.Y-origin.Y)*float64(scaleFactor)),
	)
	d.OuterContainer.MoveTo(p)
	d.OuterContainer.Scale(scaleFactor)
	for i := range d.Elements {
		innerTo := newPosition(scaleFactor, origin, d.Elements[i].InnerContainer.UpperLeft)
		d.Elements[i].InnerContainer.MoveTo(innerTo)
		outerTo := newPosition(
			scaleFactor,
			d.Elements[i].InnerContainer.UpperLeft,
			d.Elements[i].OuterContainer.UpperLeft,
		)
		d.Elements[i].OuterContainer.MoveTo(outerTo)
		d.Elements[i].Scale(scaleFactor)
	}
	d.FWidth = d.InnerContainer.Width()
	d.FHeight = d.InnerContainer.Height()
}

func newPosition(scaleFactor float64, origin Point, destino Point) Point {
	newDistance := NewPoint(
		int32(float64(destino.X-origin.X)*float64(scaleFactor)),
		int32(float64(destino.Y-origin.Y)*float64(scaleFactor)),
	)
	return NewPoint(
		newDistance.X+origin.X,
		newDistance.Y+origin.Y,
	)
}

func calculateScaleFactor(
	elementWidth, elementHeight, containerWidth, containerHeight float64,
) float64 {
	widthScaleFactor := containerWidth / elementWidth
	heightScaleFactor := containerHeight / elementHeight

	if widthScaleFactor < heightScaleFactor {
		return widthScaleFactor
	}
	return heightScaleFactor
}

func calculateGreaterScaleFactor(
	elementWidth, elementHeight, containerWidth, containerHeight float64,
) float64 {
	widthScaleFactor := containerWidth / elementWidth
	heightScaleFactor := containerHeight / elementHeight

	if widthScaleFactor > heightScaleFactor {
		return widthScaleFactor
	}
	return heightScaleFactor
}

func (d *LayoutComponent) ScaleElements(wscale, hscale float64) {
	for i := range d.Elements {
		el := &d.Elements[i]
		el.FWidth = int32(float64(el.FWidth) * wscale)
		el.FHeight = int32(float64(el.FHeight) * hscale)
	}
}

func (d *LayoutComponent) ScaleElementsPositions(wscale, hscale float64) {
	for i := range d.Elements {
		el := &d.Elements[i]
		el.Xi = int32(float64(el.Xi) * wscale)
		el.Yi = int32(float64(el.Yi) * hscale)
		el.Xii = int32(float64(el.Xii) * wscale)
		el.Yii = int32(float64(el.Yii) * hscale)
	}
}

func (d *LayoutComponent) PositionText() string {
	return fmt.Sprintf(
		"{xi:%d,yi:%d,xii:%d,yii:%d}",
		d.InnerContainer.UpperLeft.X,
		d.InnerContainer.UpperLeft.Y,
		d.InnerContainer.DownRight.X,
		d.InnerContainer.DownRight.Y,
	)
}
