package grammars

import "algvisual/internal/entities"

func StageFindColision(
	original entities.Layout,
	prevLayout entities.Layout,
	template entities.Template,
	prevGrid entities.Grid,
) (*entities.Layout, *entities.Grid, error) {
	var out entities.Layout
	var stageComponents []entities.LayoutComponent
	for _, c := range prevLayout.Components {
		if !prevGrid.HaveColisionInList(c.Positions, c.ID) {
			stageComponents = append(stageComponents, c)
			continue
		}
		positions, err := prevGrid.FindFreePositionsToFitBasedOnPivot(
			c.Pivot,
			c.InnerContainer,
		)
		if prevGrid.IsPositionListOcupiedByOtherThanThisId(c.Positions, c.ID) {
			prevGrid.RemoveFromAllCells(c.ID)
			continue
		}
		if err != nil {
			prevGrid.RemoveFromAllCells(c.ID)
			continue
		}
		cont := prevGrid.PositionsToContainer(positions)
		prevGrid.RemoveFromAllCells(c.ID)
		prevGrid.OcupyByPositionList(positions, c.ID)
		c.Positions = positions
		c.ScaleToFitInSize(cont.Width(), cont.Height())
		c.MoveTo(cont.UpperLeft)
		c.CenterInContainer(cont)
		c.GridContainer = cont
		stageComponents = append(stageComponents, c)
	}
	out.Components = stageComponents
	out.Template = template
	out.DesignID = original.DesignID
	out.Width = template.Width
	out.Height = template.Height
	out.Grid = prevGrid
	return &out, &prevGrid, nil
}
