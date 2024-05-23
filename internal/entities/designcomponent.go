package entities

type DesignComponent struct {
	ID       int32           `json:"id,omitempty"`
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
