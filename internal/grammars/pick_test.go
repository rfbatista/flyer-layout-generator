package grammars

import (
	"algvisual/internal/entities"
	"go.uber.org/zap/zaptest"
	"testing"
)

func TestPick(t *testing.T) {
	logger := zaptest.NewLogger(t)

	component1 := entities.DesignComponent{ID: 1, Xi: 10, Yi: 10, Xii: 20, Yii: 20}
	component2 := entities.DesignComponent{ID: 2, Xi: 10, Yi: 10, Xii: 20, Yii: 20}
	component3 := entities.DesignComponent{ID: 3, Xi: 15, Yi: 15, Xii: 40, Yii: 40}

	world := World{
		Components: []entities.DesignComponent{component1, component2, component3},
		TwistedDesign: entities.Layout{
			Components: []entities.DesignComponent{component1, component2, component3},
		},
	}

	prancheta := entities.Layout{
		Width:  50,
		Height: 50,
		Components: []entities.DesignComponent{
			component1, component2,
		},
	}

	pickedComponent := Pick(world, prancheta, logger)

	if pickedComponent == nil {
		t.Fatalf("Expected to pick a component, but got nil")
	}

	if pickedComponent.ID != 3 {
		t.Errorf("Expected component with ID 3 to be picked, but got ID %d", pickedComponent.ID)
	}
}
