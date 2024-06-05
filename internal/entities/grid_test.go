package entities

import (
	"testing"
)

func TestNewGrid(t *testing.T) {
	t.Run("should crete grid using pivot", func(tt *testing.T) {
		t1, err := NewGrid(WithDefault(200, 200), WithPivot(100, 100))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		if len(t1.Cells) != 4 {
			tt.Errorf("wrong number of regions: %d", len(t1.Cells))
		}
	})

	t.Run("should crete grid using cells number", func(tt *testing.T) {
		t1, err := NewGrid(WithDefault(200, 200), WithCells(4, 4))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		if len(t1.Cells) != 16 {
			tt.Errorf("wrong number of regions: %d", len(t1.Cells))
		}
		if t1.Cells[0].Width() != 50 {
			tt.Errorf("wrong width: %d", t1.Cells[0].Width())
		}
		t2, err := NewGrid(WithDefault(200, 200), WithCells(1, 1))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		if len(t2.Cells) != 1 {
			tt.Errorf("wrong number of regions: %d", len(t1.Cells))
		}
		if t2.Cells[0].Width() != 200 {
			tt.Errorf("wrong width: %d", t2.Cells[0].Width())
		}
		t3, err := NewGrid(WithDefault(200, 200), WithCells(2, 2))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		if len(t3.Cells) != 4 {
			tt.Errorf("wrong number of regions: %d", len(t1.Cells))
		}
		if t3.Cells[0].Width() != 100 {
			tt.Errorf("wrong width: %d", t3.Cells[0].Width())
		}
	})

	t.Run("should return the where the point is", func(tt *testing.T) {
		t1, err := NewGrid(WithDefault(200, 200), WithCells(4, 4))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		if t1.WhereIsPoint(NewPoint(50, 50)) == nil {
			tt.Errorf("no cell was found for point")
		}
	})
}
