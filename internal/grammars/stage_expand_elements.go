package grammars

import "algvisual/internal/entities"

func Stage4(
	original entities.Layout,
	prevLayout *entities.Layout,
	template entities.Template,
	prevGrid *entities.Grid,
) (*entities.Layout, *entities.Grid, error) {
	var out entities.Layout
	var stageComponents []entities.LayoutComponent
	stageGrid, _ := entities.NewGrid(
		entities.WithDefault(template.Width, template.Height),
		entities.WithCells(prevGrid.SlotsX, prevGrid.SlotsY),
	)
	for _, c := range prevLayout.Components {
		if !prevGrid.CantItGrow(c.Positions[0], c.InnerContainer, c.ID) {
			c.ApplyPadding(original.Config.Padding)
			stageComponents = append(stageComponents, c)
			continue
		}
		cont, err := prevGrid.FindSpaceToGrow(c.Positions[0], c.InnerContainer, c.ID)
		if err != nil || cont == nil {
			continue
		}
		gcrid := prevGrid.ContainerToPositions(*cont)
		prevGrid.OcupyByPositionList(gcrid, c.ID)
		c.MoveTo(cont.UpperLeft)
		c.ScaleToFitInSize(cont.Width(), cont.Height())
		c.CenterInContainer(*cont)
		c.ApplyPadding(original.Config.Padding)
		stageComponents = append(stageComponents, c)
	}
	out.Components = stageComponents
	out.Template = template
	out.DesignID = original.DesignID
	out.Width = template.Width
	out.Height = template.Height
	out.Grid = *stageGrid
	return &out, stageGrid, nil
}
