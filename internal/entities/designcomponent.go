package entities

func NewPosition(x, y int32) Position {
	return Position{X: x, Y: y}
}

type Position struct {
	X, Y int32
}

type DesignComponentDTO struct {
	ID       int32           `json:"id,omitempty"`
	DesignID int32           `json:"design_id,omitempty"`
	Elements []DesignElement `json:"elements,omitempty"`
	Width    int32           `json:"width,omitempty"`
	Height   int32           `json:"height,omitempty"`
	Color    string          `json:"color,omitempty"`
	Type     string          `json:"type,omitempty"`
	Xi       int32           `json:"xi,omitempty"`
	Xii      int32           `json:"xii,omitempty"`
	Yi       int32           `json:"yi,omitempty"`
	Yii      int32           `json:"yii,omitempty"`
	BboxXi   int32           `json:"bbox_xi,omitempty"`
	BboxXii  int32           `json:"bbox_xii,omitempty"`
	BboxYi   int32           `json:"bbox_yi,omitempty"`
	BboxYii  int32           `json:"bbox_yii,omitempty"`
	Xsnaped  bool            `json:"xsnaped,omitempty"`
	Ysnaped  bool            `json:"ysnaped,omitempty"`
	LeftGap  Position        `json:"left_gap,omitempty"`
	RightGap Position        `json:"right_gap,omitempty"`
}

type DesignComponent struct {
	ID             int32           `json:"id,omitempty"`
	DesignID       int32           `json:"design_id,omitempty"`
	Elements       []DesignElement `json:"elements,omitempty"`
	Width          int32           `json:"width,omitempty"`
	Height         int32           `json:"height,omitempty"`
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
	innerContainer Container
	outerContainer Container
}

func (d *DesignComponent) MoveTo(p Point) {
	displacement := d.innerContainer.DisplacementFrom(p)
	d.innerContainer.Move(displacement)
	d.outerContainer.Move(displacement)
}

func (d *DesignComponent) Center() Point {
	return Point{}
}

func (d *DesignComponent) UpLeft() Point {
	return d.innerContainer.UpperLeft
}

func (d *DesignComponent) DownRight() Point {
	return d.innerContainer.DownRight
}

func (d *DesignComponent) OrderPriority() int32 {
	switch d.Type {
	case "produto":
		return 1
	case "logo":
		return 2
	case "oferta":
		return 4
	case "modelo":
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

func (d *DesignComponent) CenterInRegion(r Cell) {
	xi := r.Xi
	yi := r.Yi
	if r.Width() > d.Width {
		xi = r.Xi + ((r.Width() - d.Width) / 2)
	}
	if r.Height() > d.Height {
		yi = r.Yi + ((r.Height() - d.Height) / 2)
	}
	d.SetPosition(xi, yi)
}

func (d *DesignComponent) ScaleToFitInSize(w, h int32) {
	scaleFactor := calculateScaleFactor(float64(d.Width), float64(d.Height), float64(w), float64(h))
	d.ScaleTo(scaleFactor, scaleFactor)
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

func (d *DesignComponent) ScaleTo(wscale, hscale float64) {
	d.Height = int32(float64(d.Height) * hscale)
	d.Width = int32(float64(d.Width) * wscale)
	d.Xi = int32(float64(d.Xi) * wscale)
	d.Yi = int32(float64(d.Yi) * hscale)
	d.Xii = d.Xi + d.Width
	d.Yii = d.Yi + d.Height
	d.ScaleElements(wscale, hscale)
	d.ScaleElementsPositions(wscale, hscale)
}

func (d *DesignComponent) ScaleWithoutMoving(wscale, hscale float64) {
	d.Height = int32(float64(d.Height) * hscale)
	d.Width = int32(float64(d.Width) * wscale)
	nxi := int32(float64(d.Xi) * wscale)
	nyi := int32(float64(d.Yi) * hscale)
	// movimentacao realizada
	mxi := nxi - d.Xi
	myi := nyi - d.Yi
	d.Xi = nxi
	d.Yi = nyi
	d.Xii = d.Xi + d.Width
	d.Yii = d.Yi + d.Height
	d.ScaleElements(wscale, hscale)
	for i := range d.Elements {
		el := &d.Elements[i]
		el.Xi = mxi + el.Xi
		el.Yi = myi + el.Yi
		el.Xii = el.Xi + el.Width
		el.Yii = el.Yi + el.Height
	}
}

func (d *DesignComponent) ScaleElements(wscale, hscale float64) {
	for i := range d.Elements {
		el := &d.Elements[i]
		el.Width = int32(float64(el.Width) * wscale)
		el.Height = int32(float64(el.Height) * hscale)
	}
}

func (d *DesignComponent) SetPosition(xi, yi int32) {
	xdif := xi - d.Xi
	ydif := yi - d.Yi
	d.Xi = xi
	d.Yi = yi
	d.Xii = xi + d.Width
	d.Yii = yi + d.Height
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
