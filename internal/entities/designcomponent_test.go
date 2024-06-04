package entities

import "testing"

func TestDesignComponent(t *testing.T) {
	t.Run("should move design component", func(tt *testing.T) {
		d := DesignComponent{
			innerContainer: NewContainer(NewPoint(5, 5), NewPoint(15, 15)),
			outerContainer: NewContainer(NewPoint(0, 0), NewPoint(20, 20)),
		}
		if d.outerContainer.UpperLeft.X != 0 {
			tt.Errorf("expected x = 0, received x = %d", d.DownRight().X)
		}
		if d.outerContainer.UpperLeft.Y != 0 {
			tt.Errorf("expected y = 0, received y = %d", d.DownRight().Y)
		}
		d.MoveTo(NewPoint(10, 10))
		if d.UpLeft().X != 10 {
			tt.Errorf("expected x = 10, received x = %d", d.UpLeft().X)
		}
		if d.UpLeft().Y != 10 {
			tt.Errorf("expected y = 10, received y = %d", d.UpLeft().Y)
		}
		if d.DownRight().X != 20 {
			tt.Errorf("expected x = 20, received x = %d", d.DownRight().X)
		}
		if d.DownRight().Y != 20 {
			tt.Errorf("expected y = 20, received y = %d", d.DownRight().Y)
		}
		if d.outerContainer.UpperLeft.X != 5 {
			tt.Errorf("expected x = 5, received x = %d", d.DownRight().X)
		}
		if d.outerContainer.UpperLeft.Y != 5 {
			tt.Errorf("expected y = 5, received y = %d", d.DownRight().Y)
		}
	})
}
