package entities

type DesginAssetProperty uint8

const (
	DesignAssetPropertyText DesginAssetProperty = iota
	DesginAssetPropertyFontName
	DesginAssetPropertyFontSize

	DesignAssetPropertyUnknown
)

func (d DesginAssetProperty) ToString() string {
	switch d {
	case DesignAssetPropertyText:
		return "text"
	}

	return "unknown"
}
