package entities

type DesignAssetType uint8

const (
	DesignAssetTypeText DesignAssetType = iota
	DesignAssetTypeSmartObject
	DesignAssetTypePSDImage
	DesignAssetTypeShape
	DesignAssetTypePixel
	DesignAssetTypeGroup
	DesignAssetTypeUnknown
)

func (d DesignAssetType) ToString() string {
	switch d {
	case DesignAssetTypeText:
		return "type"
	case DesignAssetTypeSmartObject:
		return "smartobject"
	case DesignAssetTypePSDImage:
		return "psdimage"
	case DesignAssetTypeShape:
		return "shape"
	case DesignAssetTypePixel:
		return "pixel"
	case DesignAssetTypeGroup:
		return "group"
	}
	return "unknown"
}

func StringToDesignAssetType(s string) DesignAssetType {
	switch s {
	case "type":
		return DesignAssetTypeText
	case "smartobject":
		return DesignAssetTypeSmartObject
	case "psdimage":
		return DesignAssetTypePSDImage
	case "shape":
		return DesignAssetTypeShape
	case "pixel":
		return DesignAssetTypePixel
	case "group":
		return DesignAssetTypeGroup
	}

	return DesignAssetTypeUnknown
}
