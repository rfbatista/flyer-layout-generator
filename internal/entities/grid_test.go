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
		if len(t1.Cells()) != 4 {
			tt.Errorf("wrong number of regions: %d", len(t1.Cells()))
		}
	})

	t.Run("should crete grid using cells number", func(tt *testing.T) {
		t1, err := NewGrid(WithDefault(200, 200), WithCells(4, 4))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		if len(t1.Cells()) != 16 {
			tt.Errorf("wrong number of regions: %d", len(t1.Cells()))
		}
		if t1.Cells()[0].Width() != 50 {
			tt.Errorf("wrong width: %d", t1.Cells()[0].Width())
		}
		t2, err := NewGrid(WithDefault(200, 200), WithCells(1, 1))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		if len(t2.Cells()) != 1 {
			tt.Errorf("wrong number of regions: %d", len(t1.Cells()))
		}
		if t2.Cells()[0].Width() != 200 {
			tt.Errorf("wrong width: %d", t2.Cells()[0].Width())
		}
		t3, err := NewGrid(WithDefault(200, 200), WithCells(2, 2))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		if len(t3.Cells()) != 4 {
			tt.Errorf("wrong number of regions: %d", len(t1.Cells()))
		}
		if t3.Cells()[0].Width() != 100 {
			tt.Errorf("wrong width: %d", t3.Cells()[0].Width())
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

	t.Run("should create a container fit the provided positions", func(tt *testing.T) {
		t1, err := NewGrid(WithDefault(200, 200), WithCells(4, 4))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		cont := t1.PointsToContainer([]Point{NewPoint(0, 0), NewPoint(0, 1)})
		if cont.Height() != 100 {
			tt.Errorf("expected 100 but received %d", cont.Height())
		}
		if cont.Width() != 50 {
			tt.Errorf("expected 50 but received %d", cont.Width())
		}
	})

	t.Run("should check if it fits", func(tt *testing.T) {
		t1, err := NewGrid(WithDefault(300, 300), WithCells(3, 3))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		itFits1 := t1.Fits(0, 1, NewContainer(NewPoint(0, 0), NewPoint(50, 250)))
		if itFits1 {
			tt.Errorf("expected to not fit in")
		}
		itFits2 := t1.Fits(0, 1, NewContainer(NewPoint(0, 0), NewPoint(50, 50)))
		if !itFits2 {
			tt.Errorf("expected to fit in")
		}
		itFits3 := t1.Fits(2, 2, NewContainer(NewPoint(0, 0), NewPoint(150, 50)))
		if itFits3 {
			tt.Errorf("expected to not fit in")
		}
	})

	t.Run("should find the correct positions to fit the container", func(tt *testing.T) {
		t1, err := NewGrid(WithDefault(300, 300), WithCells(3, 3))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		x, y, found := t1.FindSpace(
			NewPosition(0, 2),
			NewContainer(NewPoint(0, 0), NewPoint(100, 150)),
		)
		if !found {
			tt.Errorf("error finding position in grid: %v", err)
		}
		if x != 0 {
			tt.Errorf("expected 0 but received %d", x)
		}
		if y != 1 {
			tt.Errorf("expected 1 but received %d", y)
		}
	})

	t.Run("should crete a list of points that fit the container", func(tt *testing.T) {
		t1, err := NewGrid(WithDefault(300, 300), WithCells(3, 3))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		c, err := t1.FindPositionsToFitBasedOnPivot(
			NewPosition(0, 2),
			NewContainer(NewPoint(0, 0), NewPoint(100, 150)),
		)
		if err != nil {
			tt.Errorf("error finding position: %v", err)
		}
		if len(c) != 2 {
			tt.Errorf("expected 2 but received %d", len(c))
		}
	})

	t.Run("should find a position to fit the grid container", func(tt *testing.T) {
		tt.Run("first case", func(tt *testing.T) {
			t1, err := NewGrid(WithDefault(300, 300), WithCells(3, 3))
			if err != nil {
				tt.Errorf("error creating grid: %v", err)
			}
			gridc := NewGridContainer(NewPosition(0, 0), NewPosition(1, 1))
			_, found, err := t1.FindPositionToFitGridContainer(
				NewPosition(0, 2),
				gridc,
				10,
			)
			if err != nil {
				tt.Errorf("error finding position: %v", err)
			}
			if !found {
				tt.Errorf("expected to fit container")
			}
		})
		tt.Run("second case", func(tt *testing.T) {
			t2, err := NewGrid(WithDefault(300, 300), WithCells(3, 3))
			if err != nil {
				tt.Errorf("error creating grid: %v", err)
			}
			t2.position[0][0].Ocupy(10)
			t2.position[0][1].Ocupy(10)
			t2.position[1][0].Ocupy(10)
			gridc2 := NewGridContainer(NewPosition(0, 0), NewPosition(1, 1))
			gridcresult2, found2, err2 := t2.FindPositionToFitGridContainer(
				NewPosition(0, 2),
				gridc2,
				11,
			)
			if err2 != nil {
				tt.Errorf("error finding position: %s", err2.Error())
			}
			if !found2 {
				tt.Errorf("expected to fit container")
			}
			if gridcresult2.UpLeft.Y != 1 && gridcresult2.UpLeft.X != 1 {
				tt.Errorf("expected to be 1 received %d", gridcresult2.UpLeft.Y)
			}
		})
	})

	t.Run("should check if have a colision", func(ttt *testing.T) {
		ttt.Run("case 1", func(t *testing.T) {
			t2, err := NewGrid(WithDefault(300, 300), WithCells(3, 3))
			if err != nil {
				ttt.Errorf("error creating grid: %v", err)
			}
			t2.position[0][0].Ocupy(10)
			t2.position[0][1].Ocupy(10)
			t2.position[2][1].Ocupy(10)
			gridc2 := NewGridContainer(NewPosition(1, 1), NewPosition(2, 2))
			colision := t2.CheckGridContainerColision(gridc2, 11)
			if !colision {
				t.Errorf("should have got a colision")
			}
		})
		ttt.Run("case 1", func(t *testing.T) {
			t2, err := NewGrid(WithDefault(300, 300), WithCells(3, 3))
			if err != nil {
				ttt.Errorf("error creating grid: %v", err)
			}
			t2.position[0][0].Ocupy(10)
			t2.position[0][1].Ocupy(10)
			t2.position[1][0].Ocupy(10)
			gridc2 := NewGridContainer(NewPosition(1, 1), NewPosition(2, 2))
			colision := t2.CheckGridContainerColision(gridc2, 11)
			if colision {
				t.Errorf("should not have got a colision")
			}
		})
	})

	t.Run("should not find a position to fit the grid container", func(tt *testing.T) {
		t3, err := NewGrid(WithDefault(300, 300), WithCells(3, 3))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		t3.position[0][0].Ocupy(10)
		t3.position[0][1].Ocupy(10)
		t3.position[1][0].Ocupy(10)
		t3.position[1][1].Ocupy(10)
		gridc3 := NewGridContainer(NewPosition(0, 0), NewPosition(1, 1))
		d, found, err3 := t3.FindPositionToFitGridContainer(
			NewPosition(0, 2),
			gridc3,
			11,
		)
		if found {
			tt.Errorf("expected to not found a position %+v\n", d)
		}
		if err3 == nil {
			tt.Errorf("expected to return an error %+v\n", d)
		}
	})

	t.Run("should create a grid content from a container", func(tt *testing.T) {
		cont := NewContainer(NewPoint(0, 0), NewPoint(150, 150))
		t3, err := NewGrid(WithDefault(300, 300), WithCells(3, 3))
		if err != nil {
			tt.Error("error should be nil")
		}
		gridc := t3.ContainerToGridContainer(cont)
		if gridc.Width() != 2 {
			tt.Errorf("expected to be 2 but received %d", gridc.Width())
		}
	})

	t.Run("should check if in the list of positions have a different element", func(tt *testing.T) {
		t3, err := NewGrid(WithDefault(300, 300), WithCells(3, 3))
		if err != nil {
			tt.Errorf("error creating grid: %v", err)
		}
		t3.position[0][0].Ocupy(10)
		t3.position[0][1].Ocupy(10)
		t3.position[1][0].Ocupy(10)
		t3.position[1][1].Ocupy(10)
		if !t3.IsPositionListOcupiedByOtherThanThisId([]Position{NewPosition(0, 0)}, 3) {
			tt.Error("should have found other item in this position")
		}
		if !t3.IsPositionListOcupiedByOtherThanThisId([]Position{NewPosition(0, 0), NewPosition(1, 0)}, 3) {
			tt.Error("should have found other item in this position")
		}
		if !t3.IsPositionListOcupiedByOtherThanThisId([]Position{NewPosition(2, 2), NewPosition(1, 1)}, 3) {
			tt.Error("should have found other item in this position")
		}
		if t3.IsPositionListOcupiedByOtherThanThisId([]Position{NewPosition(2, 2), NewPosition(2, 1)}, 3) {
			tt.Error("should have found other item in this position")
		}
	})
}
