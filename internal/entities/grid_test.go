package entities

import (
	"testing"
)

func TestNewGrid(t *testing.T) {
	t1, err := NewGrid(WithDefault(200, 200), WithPivot(100, 100))
	if err != nil {
		t.Errorf("error creating grid: %v", err)
	}
	if len(t1.Regions) != 4 {
		t.Errorf("wrong number of regions: %d", len(t1.Regions))
	}
}
