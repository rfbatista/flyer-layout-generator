package grammars

import (
	"algvisual/internal/entities"
	"encoding/json"

	"go.uber.org/zap"
)

func TwistDesign(world World, prancheta entities.Layout, log *zap.Logger) entities.Layout {
	log.Info("twisting design")
	wscale, hscale := calculateScaleToFit(
		world.OriginalDesign.Width,
		world.OriginalDesign.Height,
		prancheta.Width,
		prancheta.Height,
	)
	twistedPrancheta, _ := CloneMyStruct(&prancheta)
	twistedPrancheta.Components = make([]entities.DesignComponent, len(world.Components))
	for idx, c := range world.Components {
		cc, _ := CloneComponent(&c)
		cc.ScaleTo(wscale, hscale)
		twistedPrancheta.Components[idx] = *cc
	}
	return *twistedPrancheta
}

func calculateScaleToFit(
	containerWidth, containerHeight, elementWidth, elementHeight int32,
) (float64, float64) {
	widthScale := float64(elementWidth) / float64(containerWidth)
	heightScale := float64(elementHeight) / float64(containerHeight)
	return widthScale, heightScale
}

func scaleComponent(c entities.DesignComponent, wprorp, hprorp float64) entities.DesignComponent {
	c.FWidth = int32(float64(c.FWidth) * wprorp)
	c.FHeight = int32(float64(c.FHeight) * hprorp)
	c.Xi = int32(float64(c.Xi) * wprorp)
	c.Yi = int32(float64(c.Yi) * hprorp)
	c.Xii = int32(float64(c.Xii) * wprorp)
	c.Yii = int32(float64(c.Yii) * hprorp)
	return c
}

func CloneMyStruct(orig *entities.Layout) (*entities.Layout, error) {
	origJSON, err := json.Marshal(orig)
	if err != nil {
		return nil, err
	}

	clone := entities.Layout{}
	if err = json.Unmarshal(origJSON, &clone); err != nil {
		return nil, err
	}

	return &clone, nil
}

func CloneComponent(orig *entities.DesignComponent) (*entities.DesignComponent, error) {
	origJSON, err := json.Marshal(orig)
	if err != nil {
		return nil, err
	}

	clone := entities.DesignComponent{}
	if err = json.Unmarshal(origJSON, &clone); err != nil {
		return nil, err
	}

	return &clone, nil
}
