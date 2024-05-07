package layoutengine

import (
	"math"
)

type Dimension struct {
	Width  int
	Height int
}

type DesignElement struct {
	Width  int
	Height int
	Xi     int
	Yi     int
	Xii    int
	Yii    int
}

type Component struct {
	ID       int
	Width    float64
	Height   float64
	Xi       int
	Yi       int
	Xii      int
	Yii      int
	Elements []DesignElement
}

type DesignBoard struct {
	Width      int
	Height     int
	Components []Component
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

func resizeDesignElement(element DesignElement, width int, height int) DesignElement {
	nelement := deepcopyDesignElement(element)
	nelement.Width = width
	nelement.Height = height
	nelement.Xi = int(math.Round(float64(element.Xi) * float64(width) / float64(element.Width)))
	nelement.Yi = int(math.Round(float64(element.Yi) * float64(height) / float64(element.Height)))
	nelement.Xii = nelement.Xi + nelement.Width
	nelement.Yii = nelement.Yi + nelement.Height
	return nelement
}

func resizeComponent(
	component Component,
	widthProportion float64,
	heightProportion float64,
) Component {
	widthProp := (component.Width * widthProportion) / component.Width
	heightProp := (component.Height * heightProportion) / component.Height
	ncomponent := component
	ncomponent.Xi = int(float64(component.Xi) * widthProp)
	ncomponent.Yi = int(float64(component.Yi) * heightProp)
	ncomponent.Xii = int(float64(component.Xii) * widthProp)
	ncomponent.Yii = int(float64(component.Yii) * heightProp)
	nelements := make([]DesignElement, len(component.Elements))
	for i, elem := range component.Elements {
		nelement := resizeDesignElement(
			elem,
			int(math.Round(float64(elem.Width)*widthProp)),
			int(math.Round(float64(elem.Height)*heightProp)),
		)
		nelements[i] = nelement
	}
	ncomponent.Elements = nelements
	return ncomponent
}

func resizeComponents(
	components []Component,
	widthProportion float64,
	heightProportion float64,
) []Component {
	ncomponents := make([]Component, len(components))
	for i, comp := range components {
		ncomponent := resizeComponent(comp, widthProportion, heightProportion)
		ncomponents[i] = ncomponent
	}
	return ncomponents
}

func distortImageTo(prancheta DesignBoard, to Dimension) DesignBoard {
	widthProp := float64(to.Width) / float64(prancheta.Width)
	heightProp := float64(to.Height) / float64(prancheta.Height)
	prancheta.Components = resizeComponents(prancheta.Components, widthProp, heightProp)
	return prancheta
}

func main() {
	// Your code goes here
}
