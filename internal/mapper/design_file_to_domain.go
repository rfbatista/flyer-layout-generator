package mapper

import (
	"algvisual/internal/database"
	"algvisual/internal/entities"
)

func DesignFileToDomain(raw database.Design) entities.DesignFile {
	return entities.DesignFile{
		ID:             raw.ID,
		Name:           raw.Name,
		LayoutID:       raw.LayoutID.Int32,
		ProjectID:      raw.ProjectID.Int32,
		Filepath:       raw.FileUrl.String,
		FileExtension:  raw.FileExtension.String,
		ImagePath:      raw.ImageUrl.String,
		ImageURL:       raw.ImageUrl.String,
		ImageExtension: raw.ImageExtension.String,
		Width:          raw.Width.Int32,
		Height:         raw.Height.Int32,
		CreatedAt:      raw.CreatedAt.Time,
		IsProcessed:    raw.IsProccessed.Bool,
	}
}

func TodesignEntitie(raw database.Design) entities.DesignFile {
	return entities.DesignFile{
		ID:             raw.ID,
		ProjectID:      raw.ProjectID.Int32,
		Width:          raw.Width.Int32,
		LayoutID:       raw.LayoutID.Int32,
		Height:         raw.Height.Int32,
		Name:           raw.Name,
		Filepath:       raw.FileUrl.String,
		ImageExtension: raw.ImageExtension.String,
		ImagePath:      raw.ImageUrl.String,
		CreatedAt:      raw.CreatedAt.Time,
		IsProcessed:    raw.IsProccessed.Bool,
	}
}
