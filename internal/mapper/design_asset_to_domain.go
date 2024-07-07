package mapper

import (
	"algvisual/database"
	"algvisual/internal/entities"
)

func DesignAssetToDomain(r database.DesignAsset) entities.DesignAsset {
	return entities.DesignAsset{
		ID:        r.ID,
		DesignID:  r.DesignID.Int32,
		ProjectID: r.ProjectID.Int32,
		PathURL:   r.AssetPath.String,
		AssetURL:  r.AssetUrl.String,
		Type:      entities.StringToDesignAssetType(string(r.Type.DesignAssetType)),
		Width:     r.Width.Int32,
		Height:    r.Height.Int32,
	}
}
