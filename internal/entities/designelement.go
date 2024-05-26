package entities

import (
	"math"
)

type DesignElement struct {
	ID             int32  `json:"id,omitempty"`
	Xi             int32  `json:"xi"`
	Xii            int32  `json:"xii"`
	Yi             int32  `json:"yi"`
	Yii            int32  `json:"yii"`
	LayerID        string `json:"layer_id"`
	Width          int32  `json:"width"`
	Height         int32  `json:"height"`
	Kind           string `json:"kind"`
	Name           string `json:"name"`
	IsGroup        bool   `json:"is_group"`
	GroupId        int32  `json:"group_id"`
	Level          int32  `json:"level"`
	DesignID       int32  `json:"photoshop_id"`
	ImageURL       string `json:"image,omitempty"`
	Text           string `json:"text,omitempty"`
	ImageExtension string `json:"image_extension,omitempty"`
	ComponentID    int32  `json:"component_id"`
}

func (d *DesignComponent) Widthf() float64 {
	return float64(d.Width)
}

func (d *DesignComponent) Heightf() float64 {
	return float64(d.Width)
}

func deepcopyDesignElement(element DesignElement) DesignElement {
	return DesignElement{
		Width:  element.Width,
		Height: element.Height,
		Xi:     element.Xi,
		Yi:     element.Yi,
		Xii:    element.Xii,
		Yii:    element.Yii,
	}
}

func resizeDesignElement(element DesignElement, width int32, height int32) DesignElement {
	nelement := deepcopyDesignElement(element)
	nelement.Width = width
	nelement.Height = height
	nelement.Xi = int32(math.Round(float64(element.Xi) * float64(width) / float64(element.Width)))
	nelement.Yi = int32(math.Round(float64(element.Yi) * float64(height) / float64(element.Height)))
	nelement.Xii = nelement.Xi + nelement.Width
	nelement.Yii = nelement.Yi + nelement.Height
	return nelement
}

func resizeComponent(
	component DesignComponent,
	widthProportion float64,
	heightProportion float64,
) DesignComponent {
	widthProp := (component.Widthf() * widthProportion) / component.Widthf()
	heightProp := (component.Heightf() * heightProportion) / component.Heightf()
	ncomponent := component
	ncomponent.Xi = int32(float64(component.Xi) * widthProp)
	ncomponent.Yi = int32(float64(component.Yi) * heightProp)
	ncomponent.Xii = int32(float64(component.Xii) * widthProp)
	ncomponent.Yii = int32(float64(component.Yii) * heightProp)
	nelements := make([]DesignElement, len(component.Elements))
	for i, elem := range component.Elements {
		nelement := resizeDesignElement(
			elem,
			int32(math.Round(float64(elem.Width)*widthProp)),
			int32(math.Round(float64(elem.Height)*heightProp)),
		)
		nelements[i] = nelement
	}
	ncomponent.Elements = nelements
	return ncomponent
}

func resizeComponents(
	components []DesignComponent,
	widthProportion float64,
	heightProportion float64,
) []DesignComponent {
	ncomponents := make([]DesignComponent, len(components))
	for i, comp := range components {
		ncomponent := resizeComponent(comp, widthProportion, heightProportion)
		ncomponents[i] = ncomponent
	}
	return ncomponents
}

func distortImageTo(from Dimension, to Dimension, components []DesignComponent) []DesignComponent {
	widthProp := float64(to.Width) / float64(from.Width)
	heightProp := float64(to.Height) / float64(from.Height)
	return resizeComponents(components, widthProp, heightProp)
}
