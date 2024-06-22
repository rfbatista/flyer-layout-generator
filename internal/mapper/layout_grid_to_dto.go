package mapper

import "algvisual/internal/entities"

func GridToDto(g entities.Grid) entities.GridDTO {
	return entities.GridDTO{
		AllCells: g.Cells(),
	}
}
