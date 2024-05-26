package grammars

import (
	"algvisual/internal/entities"
	"testing"
)

func TestPositionComponent(t *testing.T) {
	// Setup initial data
	component1 := entities.DesignComponent{ID: 1, Xi: 0, Yi: 0}
	component2 := entities.DesignComponent{ID: 2, Xi: 0, Yi: 0}
	twistedComponent := entities.DesignComponent{ID: 1, Xi: 640, Yi: 364}

	world := World{
		Components: []entities.DesignComponent{component1, component2},
		TwistedDesign: entities.Layout{
			Components: []entities.DesignComponent{twistedComponent},
		},
	}
	prancheta := entities.Layout{
		Components: []entities.DesignComponent{component1},
	}

	// Call the function
	newWorld, newPrancheta := PositionComponent(world, prancheta, 1)

	// Verify the results
	if newWorld.Components[0].Xi != 0 || newWorld.Components[0].Yi != 0 {
		t.Errorf("Expected component position to be (0, 0), but got (%d, %d)",
			newWorld.Components[0].Xi, newWorld.Components[0].Yi)
	}

	// Ensure prancheta is returned unchanged or correctly modified based on your logic
	if newPrancheta.Components[0].Xi != 640 || newPrancheta.Components[0].Yi != 364 {
		t.Errorf("Expected to new prancheta have element with postion (640, 364), but got (%d, %d)",
			newPrancheta.Components[0].Xi, newPrancheta.Components[0].Yi)
	}
}
