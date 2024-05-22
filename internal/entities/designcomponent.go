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

func (d *DesignComponent) ScaleElements(wscale, hscale float64) {
	for i := range d.Elements {
		el := &d.Elements[i]
		el.Width = int32(float64(el.Width) * wscale)
		el.Height = int32(float64(el.Height) * hscale)
	}
}

func (d *DesignComponent) SetPosition(xi, yi int32) {
	d.Xi = xi
	d.Yi = xi
	d.Xii = xi + d.Width
	d.Yii = yi + d.Height
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
