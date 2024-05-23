package grammars

import (
	"algvisual/internal/entities"
	"go.uber.org/zap"
)

type World struct {
	OriginalDesign entities.DesignFile
	Components     []entities.DesignComponent
	Elements       []entities.DesignElement
	TwistedDesign  entities.Prancheta
}

type Grammar func(world World, prancheta entities.Prancheta) (*World, *entities.Prancheta, error)

func Run(
	world World,
	prancheta entities.Prancheta,
	log *zap.Logger,
) (World, entities.Prancheta, error) {
	world.TwistedDesign = TwistDesign(world, prancheta, log)
	for i := 0; i <= len(world.Components); i++ {
		c := Pick(world, prancheta, log)
		if c == nil {
			return world, prancheta, nil
		}
		prancheta.Components = append(prancheta.Components, *c)
		_, prancheta = PositionComponent(world, prancheta, c.ID)
		_, prancheta = ScaleComponent(world, prancheta, c.ID)
		_, prancheta = CenterComponent(world, prancheta, c.ID)
	}
	return world, prancheta, nil
}
