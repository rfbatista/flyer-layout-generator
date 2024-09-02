package mapper

import "algvisual/internal/domain/entities"

func GridToDto(g entities.Grid) entities.GridDTO {
	return entities.GridDTO{
		AllCells: g.Cells(),
	}
}
