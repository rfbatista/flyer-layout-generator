package entities

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type LayoutElementDTO struct {
	ID             int32  `json:"id"`
	Xi             int32  `json:"xi"`
	Xii            int32  `json:"xii"`
	Yi             int32  `json:"yi"`
	Yii            int32  `json:"yii"`
	InnerXi        int32  `json:"inner_xi"`
	InnerXii       int32  `json:"inner_xii"`
	InnerYi        int32  `json:"inner_yi"`
	InnerYii       int32  `json:"inner_yii"`
	LayerID        string `json:"layer_id"`
	Width          int32  `json:"width"`
	Height         int32  `json:"height"`
	Kind           string `json:"kind"`
	Name           string `json:"name"`
	IsGroup        bool   `json:"is_group"`
	GroupId        int32  `json:"group_id"`
	Level          int32  `json:"level"`
	DesignID       int32  `json:"photoshop_id"`
	ImageURL       string `json:"image"`
	Text           string `json:"text"`
	ImageExtension string `json:"image_extension"`
	ComponentID    int32  `json:"component_id"`
}

func NewLayoutElement(xi, xii, yi, yii int32) LayoutElement {
	return LayoutElement{
		OuterContainer: NewContainer(NewPoint(xi, yi), NewPoint(xii, yii)),
	}
}

type LayoutElement struct {
	ID             int32                     `json:"id,omitempty"`
	Xi             int32                     `json:"xi,omitempty"`
	Xii            int32                     `json:"xii,omitempty"`
	Yi             int32                     `json:"yi,omitempty"`
	Yii            int32                     `json:"yii,omitempty"`
	AssetID        int32                     `json:"asset_id,omitempty"`
	InnerXi        int32                     `json:"inner_xi,omitempty"`
	InnerXii       int32                     `json:"inner_xii,omitempty"`
	InnerYi        int32                     `json:"inner_yi,omitempty"`
	InnerYii       int32                     `json:"inner_yii,omitempty"`
	LayerID        string                    `json:"layer_id,omitempty"`
	FWidth         int32                     `json:"width,omitempty"`
	FHeight        int32                     `json:"height,omitempty"`
	Kind           string                    `json:"kind,omitempty"`
	Name           string                    `json:"name,omitempty"`
	IsGroup        bool                      `json:"is_group,omitempty"`
	GroupId        int32                     `json:"group_id,omitempty"`
	Level          int32                     `json:"level,omitempty"`
	DesignID       int32                     `json:"photoshop_id,omitempty"`
	ImageURL       string                    `json:"image,omitempty"`
	Text           string                    `json:"text,omitempty"`
	ImageExtension string                    `json:"image_extension,omitempty"`
	ComponentID    int32                     `json:"component_id,omitempty"`
	InnerContainer Container                 `json:"inner_container,omitempty"`
	OuterContainer Container                 `json:"outer_container,omitempty"`
	Properties     []DesignAssetPropertyData `json:"properties,omitempty"`
}

func (d *LayoutElement) TextFromProperty() string {
	for _, p := range d.Properties {
		if p.Key == DesignAssetPropertyText.ToString() {
			return p.Value
		}
	}
	return ""
}

func (d *LayoutElement) PickTextFromProperty() string {
	var texts []string
	for _, p := range d.Properties {
		if p.Key == DesignAssetPropertyText.ToString() {
			texts = append(texts, p.Value)
		}
	}
	if len(texts) == 0 {
		return ""
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	// Pick a random index from the texts slice
	randomIndex := rand.Intn(len(texts))

	return texts[randomIndex]
}

func (d *LayoutElement) PositionText() string {
	return fmt.Sprintf(
		"{xi:%d,yi:%d,xii:%d,yii:%d}",
		d.InnerContainer.UpperLeft.X,
		d.InnerContainer.UpperLeft.Y,
		d.InnerContainer.DownRight.X,
		d.InnerContainer.DownRight.Y,
	)
}

func (d *LayoutElement) Width() int32 {
	return d.InnerContainer.Width()
}

func (d *LayoutElement) OWidth() int32 {
	return d.OuterContainer.Width()
}

func (d *LayoutElement) Height() int32 {
	return d.InnerContainer.Height()
}

func (d *LayoutElement) OHeight() int32 {
	return d.OuterContainer.Height()
}

func (d *LayoutElement) UpLeft() Point {
	return d.InnerContainer.UpperLeft
}

func (d *LayoutElement) DownRight() Point {
	return d.InnerContainer.DownRight
}

func (d *LayoutElement) Center() Point {
	return d.InnerContainer.Center()
}

func (d *LayoutElement) ScaleFix(s float64) {
	nw := float64(d.OuterContainer.Width()) * s
	nh := float64(d.OuterContainer.Height()) * s
	outer := NewContainer(
		d.OuterContainer.UpperLeft,
		NewPoint(
			d.OuterContainer.UpperLeft.X+int32(nw),
			d.OuterContainer.UpperLeft.Y+int32(nh),
		),
	)
	niw := float64(d.InnerContainer.Width()) * s
	nih := float64(d.InnerContainer.Height()) * s
	nix := int32(float64(d.InnerContainer.UpperLeft.X)-float64(d.OuterContainer.UpperLeft.X)*s) + d.OuterContainer.UpperLeft.X
	niy := int32(float64(d.InnerContainer.UpperLeft.Y)-float64(d.OuterContainer.UpperLeft.Y)*s) + d.OuterContainer.UpperLeft.Y
	inner := NewContainer(
		NewPoint(nix, niy),
		NewPoint(nix+int32(niw), niy+int32(nih)),
	)
	d.OuterContainer = outer
	d.InnerContainer = inner
}

func (d *LayoutElement) Scale(s float64) {
	d.InnerContainer.Scale(s)
	d.OuterContainer.Scale(s)
}

func (d *LayoutElement) MoveOnOuter(p Point) {
	distance := d.OuterContainer.DisplacementFrom(p)
	d.InnerContainer.Move(distance)
	d.OuterContainer.Move(distance)
}

func (d *LayoutElement) MoveTo(p Point) {
	distance := d.InnerContainer.DisplacementFrom(p)
	d.InnerContainer.Move(distance)
	d.OuterContainer.Move(distance)
}

func (d *LayoutElement) Move(p Point) {
	d.InnerContainer.Move(p)
	d.OuterContainer.Move(p)
}

func (d *LayoutComponent) Widthf() float64 {
	return float64(d.FWidth)
}

func (d *LayoutComponent) Heightf() float64 {
	return float64(d.FWidth)
}

func deepcopyDesignElement(element LayoutElement) LayoutElement {
	return LayoutElement{
		FWidth:  element.FWidth,
		FHeight: element.FHeight,
		Xi:      element.Xi,
		Yi:      element.Yi,
		Xii:     element.Xii,
		Yii:     element.Yii,
	}
}

func resizeDesignElement(element LayoutElement, width int32, height int32) LayoutElement {
	nelement := deepcopyDesignElement(element)
	nelement.FWidth = width
	nelement.FHeight = height
	nelement.Xi = int32(math.Round(float64(element.Xi) * float64(width) / float64(element.FWidth)))
	nelement.Yi = int32(
		math.Round(float64(element.Yi) * float64(height) / float64(element.FHeight)),
	)
	nelement.Xii = nelement.Xi + nelement.FWidth
	nelement.Yii = nelement.Yi + nelement.FHeight
	return nelement
}

func resizeComponent(
	component LayoutComponent,
	widthProportion float64,
	heightProportion float64,
) LayoutComponent {
	widthProp := (component.Widthf() * widthProportion) / component.Widthf()
	heightProp := (component.Heightf() * heightProportion) / component.Heightf()
	ncomponent := component
	ncomponent.Xi = int32(float64(component.Xi) * widthProp)
	ncomponent.Yi = int32(float64(component.Yi) * heightProp)
	ncomponent.Xii = int32(float64(component.Xii) * widthProp)
	ncomponent.Yii = int32(float64(component.Yii) * heightProp)
	nelements := make([]LayoutElement, len(component.Elements))
	for i, elem := range component.Elements {
		nelement := resizeDesignElement(
			elem,
			int32(math.Round(float64(elem.FWidth)*widthProp)),
			int32(math.Round(float64(elem.FHeight)*heightProp)),
		)
		nelements[i] = nelement
	}
	ncomponent.Elements = nelements
	return ncomponent
}

func resizeComponents(
	components []LayoutComponent,
	widthProportion float64,
	heightProportion float64,
) []LayoutComponent {
	ncomponents := make([]LayoutComponent, len(components))
	for i, comp := range components {
		ncomponent := resizeComponent(comp, widthProportion, heightProportion)
		ncomponents[i] = ncomponent
	}
	return ncomponents
}

func distortImageTo(from Dimension, to Dimension, components []LayoutComponent) []LayoutComponent {
	widthProp := float64(to.Width) / float64(from.Width)
	heightProp := float64(to.Height) / float64(from.Height)
	return resizeComponents(components, widthProp, heightProp)
}
