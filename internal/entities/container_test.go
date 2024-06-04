package entities

import "testing"

func TestContainer(t *testing.T) {
	t.Run("should have expected width and height", func(tt *testing.T) {
		c1 := NewContainer(NewPoint(0, 0), NewPoint(10, 10))
		if c1.Width() != 10 {
			tt.Errorf("expected width 10 but received %d", c1.Width())
		}

		if c1.Height() != 10 {
			tt.Errorf("expected height 10 but received %d", c1.Height())
		}
	})

	t.Run("should move container as expected", func(tt *testing.T) {
		c1 := NewContainer(NewPoint(0, 0), NewPoint(10, 10))
		c1.MoveTo(NewPoint(5, 5))
		if c1.UpperLeft.Y != 5 {
			tt.Errorf("expected to be in y = 5, but is in y = %d", c1.UpperLeft.Y)
		}
		if c1.UpperLeft.X != 5 {
			tt.Errorf("expected to be in x = 5, but is in x = %d", c1.UpperLeft.X)
		}
		if c1.DownRight.X != 15 {
			tt.Errorf("expected to be in x = 15, but is in x = %d", c1.DownRight.X)
		}
		if c1.DownRight.Y != 15 {
			tt.Errorf("expected to be in y = 15, but is in y = %d", c1.DownRight.Y)
		}

		c2 := NewContainer(NewPoint(500, 450), NewPoint(550, 500))
		c2.MoveTo(NewPoint(235, 125))
		if c2.UpperLeft.Y != 125 {
			tt.Errorf("expected to be in y = 125, but is in y = %d", c2.UpperLeft.Y)
		}
		if c2.UpperLeft.X != 235 {
			tt.Errorf("expected to be in x = 235, but is in x = %d", c2.UpperLeft.X)
		}
		if c2.DownRight.X != 285 {
			tt.Errorf("expected to be in x = 285, but is in x = %d", c2.DownRight.X)
		}
		if c2.DownRight.Y != 175 {
			tt.Errorf("expected to be in y = 175, but is in y = %d", c2.DownRight.Y)
		}
	})
}
