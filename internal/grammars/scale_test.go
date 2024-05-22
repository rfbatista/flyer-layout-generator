package grammars

import (
	"algvisual/internal/entities"
	"reflect"
	"testing"
)

func TestScaleComponent(t *testing.T) {
	tests := []struct {
		name       string
		world      World
		prancheta  entities.Prancheta
		id         int32
		wantWorld  World
		wantPranch entities.Prancheta
	}{
		{
			name: "should scale",
			world: World{
				Components: []entities.DesignComponent{
					{
						ID:       1,
						Width:    50,
						Height:   50,
						Xi:       25,
						Yi:       25,
						Xii:      75,
						Yii:      75,
						Elements: []entities.DesignElement{entities.DesignElement{ID: 1, Width: 50, Height: 50, Xi: 25, Xii: 75, Yi: 25, Yii: 75}},
					},
				},
				OriginalDesign: entities.DesignFile{Width: 100, Height: 100},
			},
			prancheta: entities.Prancheta{
				Width:  200,
				Height: 200,
				Components: []entities.DesignComponent{
					{
						ID:       1,
						Width:    50,
						Height:   50,
						Xi:       25,
						Yi:       25,
						Xii:      75,
						Yii:      75,
						Elements: []entities.DesignElement{entities.DesignElement{ID: 1, Width: 50, Height: 50, Xi: 25, Xii: 75, Yi: 25, Yii: 75}},
					},
				},
			},
			id: 1,
			wantWorld: World{
				Components: []entities.DesignComponent{
					{
						ID:       1,
						Width:    50,
						Height:   50,
						Xi:       25,
						Yi:       25,
						Xii:      75,
						Yii:      75,
						Elements: []entities.DesignElement{entities.DesignElement{ID: 1, Width: 50, Height: 50, Xi: 25, Xii: 75, Yi: 25, Yii: 75}},
					},
				},
				OriginalDesign: entities.DesignFile{Width: 100, Height: 100},
			},
			wantPranch: entities.Prancheta{
				Width:  200,
				Height: 200,
				Components: []entities.DesignComponent{
					{
						ID:       1,
						Width:    100,
						Height:   100,
						Xi:       25,
						Yi:       25,
						Xii:      125,
						Yii:      125,
						Elements: []entities.DesignElement{entities.DesignElement{ID: 1, Width: 100, Height: 100, Xi: 25, Xii: 125, Yi: 25, Yii: 125}},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWorld, gotPranch := ScaleComponent(tt.world, tt.prancheta, tt.id)
			if !reflect.DeepEqual(gotWorld, tt.wantWorld) {
				t.Errorf("ScaleComponent() gotWorld = %v, want %v", gotWorld, tt.wantWorld)
			}
			if !reflect.DeepEqual(gotPranch, tt.wantPranch) {
				t.Errorf("ScaleComponent() \n\tgotPrancheta = %+v, \n\twant %+v", gotPranch, tt.wantPranch)
			}
		})
	}
}
