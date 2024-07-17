package grammars

import "algvisual/internal/entities"

func ExpandElements(
	original entities.Layout,
	prevLayout *entities.Layout,
	template entities.Template,
	prevGrid *entities.Grid,
) (*entities.Layout, *entities.Grid, error) {
	var out entities.Layout
	var stageComponents []entities.LayoutComponent
	// TODO: this can be improved
	for _, c := range prevLayout.Components {
		if !prevGrid.CantItGrow(c.Positions, c.ID) {
			c.ApplyPadding(original.Config.Padding)
			stageComponents = append(stageComponents, c)
			continue
		}
		cont, _ := prevGrid.FindSpaceToGrow(c.Pivot, c.InnerContainer, c.ID)
		w := cont.Width()
		h := cont.Height()
		c.ScaleToFitInSize(w, h)
		gcrid := prevGrid.ContainerToPositions(*cont)
		prevGrid.OcupyByPositionList(gcrid, c.ID)
		c.MoveTo(cont.UpperLeft)
		c.CenterInContainer(*cont)
		c.ApplyPadding(original.Config.Padding)
		innerW := c.Width()
		innerH := c.Height()
		if innerW <= 50 {
			continue
		}
		if innerH <= 50 {
			continue
		}
		stageComponents = append(stageComponents, c)
	}
	out.Components = stageComponents
	out.Template = template
	out.DesignID = original.DesignID
	out.Width = template.Width
	out.Height = template.Height
	out.Grid = *prevGrid
	return &out, prevGrid, nil
}
