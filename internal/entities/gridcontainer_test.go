package entities

import "testing"

func TestGridContainer(tt *testing.T) {
	tt.Run("should create a container", func(t *testing.T) {
		gc := NewGridContainer(NewPosition(0, 0), NewPosition(1, 1))
		cont := gc.ToContainer(100, 100)
		if cont.Width() != 200 {
			t.Errorf("expected 200 received %d", cont.Width())
		}
		if cont.Height() != 200 {
			t.Errorf("expected 200 received %d", cont.Height())
		}
	})
}
