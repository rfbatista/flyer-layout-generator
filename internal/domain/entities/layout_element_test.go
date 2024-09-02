package entities

import (
	"testing"
)

func TestLayoutElement(tt *testing.T) {
	tt.Run("test if element scale", func(t *testing.T) {
		c := NewLayoutElement(0, 100, 0, 120)
		c.InnerContainer = NewContainer(NewPoint(10, 10), NewPoint(90, 110))
		c.Scale(0.5)
		if c.OWidth() != 50 {
			t.Errorf("expected 50 but received %d", c.OWidth())
		}
		if c.OHeight() != 60 {
			t.Errorf("expected 60 but received %d", c.OHeight())
		}
		if c.Width() != 40 {
			t.Errorf("expected 40 but received %d", c.Width())
		}
		if c.Height() != 50 {
			t.Errorf("expected 50 but received %d", c.Height())
		}
	})
}
