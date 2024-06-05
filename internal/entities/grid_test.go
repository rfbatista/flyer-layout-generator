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

	t.Run("should return container to fit", func(tt *testing.T) {
		t1, err := NewGrid(WithDefault(200, 200), WithCells(4, 4))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		cont := NewContainer(NewPoint(0, 0), NewPoint(50, 60))
		if cont.Height() != 60 {
			tt.Errorf("expected 60 but received %d", cont.Height())
		}
		c := t1.ContainerToFit(cont)
		if c.Width() != 50 {
			tt.Errorf("expected 50 but received %d", c.Width())
		}
		if c.Height() != 100 {
			tt.Errorf("expected 100 but received %d", c.Height())
		}
	})

	t.Run("should create a container fot the provided positions", func(tt *testing.T) {
		t1, err := NewGrid(WithDefault(200, 200), WithCells(4, 4))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		cont := t1.PositionsToContainer([]Point{NewPoint(0, 0), NewPoint(0, 1)})
		if cont.Height() != 100 {
			tt.Errorf("expected 100 but received %d", cont.Height())
		}
		if cont.Width() != 50 {
			tt.Errorf("expected 50 but received %d", cont.Width())
		}
	})

	t.Run("should find the correct positions to fit the container", func(tt *testing.T) {
		t1, err := NewGrid(WithDefault(300, 300), WithCells(3, 3))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		points, err := t1.FindPositionsToFitBasedOnPivot(NewPoint(2, 2), NewContainer(NewPoint(0, 0), NewPoint(150, 100)))
		if err != nil {
			tt.Errorf("error finding position in grid: %v", err)
		}
		if len(points) != 2 {
			tt.Errorf("expected 2 but received %d, %+v\n", len(points), points)
		}
		points2, err := t1.FindPositionsToFitBasedOnPivot(NewPoint(2, 2), NewContainer(NewPoint(0, 0), NewPoint(210, 130)))
		if err != nil {
			tt.Errorf("error finding position in grid: %v", err)
		}
		if len(points2) != 6 {
			tt.Errorf("expected 6 but received %d, %+v\n", len(points2), points2)
		}
	})
}
