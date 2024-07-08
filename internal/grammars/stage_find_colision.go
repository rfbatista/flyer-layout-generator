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
	stagegrid, _ := entities.NewGrid(
		entities.WithDefault(template.Width, template.Height),
		entities.WithCells(prevGrid.SlotsX, prevGrid.SlotsY),
	)
	for _, c := range prevLayout.Components {
		if !prevGrid.HaveColisionInList(c.Positions, c.ID) {
			stageComponents = append(stageComponents, c)
			continue
		}
		positions, err := stagegrid.FindFreePositionsToFitBasedOnPivot(
			c.Pivot,
			c.InnerContainer,
		)
		if stagegrid.IsPositionListOcupiedByOtherThanThisId(c.Positions, c.ID) {
			stagegrid.RemoveFromAllCells(c.ID)
			continue
		}
		if err != nil {
			stagegrid.RemoveFromAllCells(c.ID)
			continue
		}
		cont := stagegrid.PositionsToContainer(positions)
		stagegrid.RemoveFromAllCells(c.ID)
		stagegrid.OcupyByPositionList(positions, c.ID)
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
	out.Grid = *stagegrid
	return &out, stagegrid, nil
}
