package entities

func NewPosition(x, y int32) Position {
	return Position{X: x, Y: y}
}

type Position struct {
	X, Y int32
}

type DesignComponent struct {
	ID       int32           `json:"id"`
	DesignID int32           `json:"design_id,omitempty"`
	Elements []DesignElement `json:"elements,omitempty"`
	Width    int32           `json:"width"`
	Height   int32           `json:"height"`
	Color    string          `json:"color,omitempty"`
	Type     string          `json:"type,omitempty"`
	Xi       int32           `json:"xi"`
	Xii      int32           `json:"xii"`
	Yi       int32           `json:"yi"`
	Yii      int32           `json:"yii"`
	BboxXi   int32           `json:"bbox_xi"`
	BboxXii  int32           `json:"bbox_xii"`
	BboxYi   int32           `json:"bbox_yi"`
	BboxYii  int32           `json:"bbox_yii"`
	Xsnaped  bool
	Ysnaped  bool
	LeftGap  Position
	RightGap Position
	UpGap    Position
	DownGap  Position
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

func (d *DesignComponent) CenterInRegion(r Region) {
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

func calculateScaleFactor(elementWidth, elementHeight, containerWidth, containerHeight float64) float64 {
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
