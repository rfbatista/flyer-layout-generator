package grammars

import (
	"algvisual/internal/entities"
)

type World struct {
	OriginalDesign entities.DesignFile
	Components     []entities.LayoutComponent
	Elements       []entities.LayoutElement
	PivotWidth     int32
	PivotHeight    int32
	TwistedDesign  entities.Layout
	Config         entities.LayoutRequestConfig
}

type Grammar func(world World, prancheta entities.Layout) (*World, *entities.Layout, error)
