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

type DesignComponentDTO struct {
	ID       int32              `json:"id,omitempty"`
	DesignID int32              `json:"design_id,omitempty"`
	Elements []DesignElementDTO `json:"elements,omitempty"`
	Width    int32              `json:"width"`
	Height   int32              `json:"height"`
	Color    string             `json:"color,omitempty"`
	Type     string             `json:"type,omitempty"`
	Xi       int32              `json:"xi"`
	Xii      int32              `json:"xii"`
	Yi       int32              `json:"yi"`
	Yii      int32              `json:"yii"`
	BboxXi   int32              `json:"bbox_xi,omitempty"`
	BboxXii  int32              `json:"bbox_xii,omitempty"`
	BboxYi   int32              `json:"bbox_yi,omitempty"`
	BboxYii  int32              `json:"bbox_yii,omitempty"`
	Xsnaped  bool               `json:"xsnaped,omitempty"`
	Ysnaped  bool               `json:"ysnaped,omitempty"`
	LeftGap  Position           `json:"left_gap,omitempty"`
	RightGap Position           `json:"right_gap,omitempty"`
}

type DesignComponent struct {
	ID             int32           `json:"id,omitempty"`
	DesignID       int32           `json:"design_id,omitempty"`
	Elements       []DesignElement `json:"elements,omitempty"`
	FWidth         int32           `json:"width,omitempty"`
	FHeight        int32           `json:"height,omitempty"`
	Color          string          `json:"color,omitempty"`
	Type           string          `json:"type,omitempty"`
	Xi             int32           `json:"xi,omitempty"`
	Xii            int32           `json:"xii,omitempty"`
	Yi             int32           `json:"yi,omitempty"`
	Yii            int32           `json:"yii,omitempty"`
	BboxXi         int32           `json:"bbox_xi,omitempty"`
	BboxXii        int32           `json:"bbox_xii,omitempty"`
	BboxYi         int32           `json:"bbox_yi,omitempty"`
	BboxYii        int32           `json:"bbox_yii,omitempty"`
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

func (d *DesignComponent) Width() int32 {
	return d.InnerContainer.Width()
}

func (d *DesignComponent) Height() int32 {
	return d.InnerContainer.Height()
}

func (d *DesignComponent) MoveTo(p Point) {
	displacement := d.InnerContainer.DisplacementFrom(p)
	d.InnerContainer.Move(displacement)
	d.OuterContainer.Move(displacement)
	for idx := range d.Elements {
		d.Elements[idx].Move(displacement)
	}
}

func (d *DesignComponent) Center() Point {
	return d.InnerContainer.Center()
}

func (d *DesignComponent) UpLeft() Point {
	return d.InnerContainer.UpperLeft
}

func (d *DesignComponent) DownRight() Point {
	return d.InnerContainer.DownRight
}

func (d *DesignComponent) OrderPriority() int32 {
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

func (d *DesignComponent) BboxWidth() int32 {
	return d.BboxXii - d.BboxXi
}

func (d *DesignComponent) BboxHeigth() int32 {
	return d.BboxYii - d.BboxYi
}

func (d *DesignComponent) IsBackground() bool {
	return d.Type == "background"
}

func (d *DesignComponent) CenterInRegion(r GridCell) {
	xi := r.Xi
	yi := r.Yi
	if r.Width() > d.Width() {
		xi = r.Xi + ((r.Width() - d.FWidth) / 2)
	}
	if r.Height() > d.FHeight {
		yi = r.Yi + ((r.Height() - d.FHeight) / 2)
	}
	d.SetPosition(xi, yi)
}

func (d *DesignComponent) CenterInContainer(r Container) {
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

func (d *DesignComponent) ApplyPadding(p int32) {
	d.InnerContainer.Padding(p)
	d.OuterContainer.Padding(p)
	for i := range d.Elements {
		d.Elements[i].InnerContainer.Padding(p)
		d.Elements[i].OuterContainer.Padding(p)
	}
}

func (d *DesignComponent) ScaleToFitInSize(w, h int32) {
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

func (d *DesignComponent) ScaleToFillInSize(w, h int32) {
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

func (d *DesignComponent) ScaleTo(wscale, hscale float64) {
	d.FHeight = int32(float64(d.FHeight) * hscale)
	d.FWidth = int32(float64(d.FWidth) * wscale)
	d.Xi = int32(float64(d.Xi) * wscale)
	d.Yi = int32(float64(d.Yi) * hscale)
	d.Xii = d.Xi + d.FWidth
	d.Yii = d.Yi + d.FHeight
	d.ScaleElements(wscale, hscale)
	d.ScaleElementsPositions(wscale, hscale)
}

func (d *DesignComponent) ScaleWithoutMoving(wscale, hscale float64) {
	d.FHeight = int32(float64(d.FHeight) * hscale)
	d.FWidth = int32(float64(d.FWidth) * wscale)
	nxi := int32(float64(d.Xi) * wscale)
	nyi := int32(float64(d.Yi) * hscale)
	// movimentacao realizada
	mxi := nxi - d.Xi
	myi := nyi - d.Yi
	d.Xi = nxi
	d.Yi = nyi
	d.Xii = d.Xi + d.FWidth
	d.Yii = d.Yi + d.FHeight
	d.ScaleElements(wscale, hscale)
	for i := range d.Elements {
		el := &d.Elements[i]
		el.Xi = mxi + el.Xi
		el.Yi = myi + el.Yi
		el.Xii = el.Xi + el.FWidth
		el.Yii = el.Yi + el.FHeight
	}
}

func (d *DesignComponent) ScaleElements(wscale, hscale float64) {
	for i := range d.Elements {
		el := &d.Elements[i]
		el.FWidth = int32(float64(el.FWidth) * wscale)
		el.FHeight = int32(float64(el.FHeight) * hscale)
	}
}

func (d *DesignComponent) SetPosition(xi, yi int32) {
	xdif := xi - d.Xi
	ydif := yi - d.Yi
	d.Xi = xi
	d.Yi = yi
	d.Xii = xi + d.FWidth
	d.Yii = yi + d.FHeight
	for i := range d.Elements {
		el := &d.Elements[i]
		el.Xi += xdif
		el.Yi += ydif
		el.Xii += xdif
		el.Yii += ydif
	}
	return
}

func (d *DesignComponent) ScaleElementsPositions(wscale, hscale float64) {
	for i := range d.Elements {
		el := &d.Elements[i]
		el.Xi = int32(float64(el.Xi) * wscale)
		el.Yi = int32(float64(el.Yi) * hscale)
		el.Xii = int32(float64(el.Xii) * wscale)
		el.Yii = int32(float64(el.Yii) * hscale)
	}
}
