package mapper

import "algvisual/internal/domain/entities"

func LayoutToDto(e entities.Layout) entities.LayoutDTO {
	var cc []entities.LayoutComponentDTO
	for _, c := range e.Components {
		cc = append(cc, DesignComponentToDto(c))
	}
	var bg *entities.LayoutComponentDTO
	if e.Background != nil {
		bge := DesignComponentToDto(*e.Background)
		bg = &bge
	}
	return entities.LayoutDTO{
		ID:         e.ID,
		Background: bg,
		ImageURL:   e.ImageURL,
		DesignID:   e.DesignID,
		Width:      e.Width,
		Height:     e.Height,
		Components: cc,
		Template:   e.Template,
		Grid:       GridToDto(e.Grid),
	}
}
