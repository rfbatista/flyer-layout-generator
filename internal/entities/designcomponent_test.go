package entities

import "testing"

func TestDesignComponent(t *testing.T) {
	t.Run("should move design component", func(tt *testing.T) {
		d := DesignComponent{
			InnerContainer: NewContainer(NewPoint(5, 5), NewPoint(15, 15)),
			OuterContainer: NewContainer(NewPoint(0, 0), NewPoint(20, 20)),
		}
		if d.OuterContainer.UpperLeft.X != 0 {
			tt.Errorf("expected x = 0, received x = %d", d.DownRight().X)
		}
		if d.OuterContainer.UpperLeft.Y != 0 {
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
		if d.OuterContainer.UpperLeft.X != 5 {
			tt.Errorf("expected x = 5, received x = %d", d.DownRight().X)
		}
		if d.OuterContainer.UpperLeft.Y != 5 {
			tt.Errorf("expected y = 5, received y = %d", d.DownRight().Y)
		}
	})

	t.Run("should scale correctly", func(tt *testing.T) {
		scale := calculateScaleFactor(100, 100, 200, 200)
		if scale != float64(2) {
			tt.Errorf("expected 2 received %f", scale)
		}
		scale2 := calculateScaleFactor(100, 100, 300, 200)
		if scale2 != float64(2) {
			tt.Errorf("expected 2 received %f", scale2)
		}
	})

	t.Run("should scale component correctly", func(tt *testing.T) {
		d := DesignComponent{
			InnerContainer: NewContainer(NewPoint(5, 5), NewPoint(15, 15)),
			OuterContainer: NewContainer(NewPoint(0, 0), NewPoint(20, 20)),
		}
		d.ScaleToFitInSize(20, 30)
		if d.InnerContainer.Width() != 20 {
			tt.Errorf("expected width = 20, received x = %d", d.InnerContainer.Width())
		}
		if d.InnerContainer.Height() != 20 {
			tt.Errorf("expected height = 20, received x = %d", d.InnerContainer.Height())
		}
	})

}
