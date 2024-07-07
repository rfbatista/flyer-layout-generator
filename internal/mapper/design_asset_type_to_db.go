package mapper

import (
	"algvisual/database"
	"algvisual/internal/entities"
)

func DesignAssetTypeToDB(d entities.DesignAssetType) database.DesignAssetType {
	switch d {
	case entities.DesignAssetTypeText:
		return database.DesignAssetTypeText
	case entities.DesignAssetTypePixel:
		return database.DesignAssetTypePixel
	case entities.DesignAssetTypeGroup:
		return database.DesignAssetTypeGroup
	case entities.DesignAssetTypeSmartObject:
		return database.DesignAssetTypeSmartobject
	}
	return database.DesignAssetTypeUnknown
}
