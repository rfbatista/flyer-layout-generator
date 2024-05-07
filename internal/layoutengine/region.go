package layoutengine

type DesignTemplateRegion struct {
	Xi  int
	Yi  int
	Xii int
	Yii int
}

type ComponentStatus struct {
	C        Component
	PixelCnt int
}

func defineComponentsPerRegion(
	regions []DesignTemplateRegion,
	components []Component,
) []DesignTemplateRegion {
	componentsIn := make(map[int]bool)
	for _, reg := range regions {
		statusComponents := make([]ComponentStatus, 0)
		for _, c := range components {
			if _, ok := componentsIn[c.ID]; !ok {
				statusComponents = append(statusComponents, ComponentStatus{C: c})
			}
		}
		for x := reg.Xi; x < reg.Xii; x++ {
			for y := reg.Yi; y < reg.Yii; y++ {
				for range statusComponents {
				}
			}
		}
		var chosenComponent *ComponentStatus
		for range statusComponents {
		}
		if chosenComponent != nil {
		}
	}
	return regions
}
