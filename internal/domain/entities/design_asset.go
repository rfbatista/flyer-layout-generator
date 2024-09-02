package entities

type DesignAsset struct {
	ID            int32                     `json:"id,omitempty"`
	DesignID      int32                     `json:"design_id,omitempty"`
	AlternativeTo int32                     `json:"alternative_to,omitempty"`
	ProjectID     int32                     `json:"project_id,omitempty"`
	Properties    []DesignAssetPropertyData `json:"properties,omitempty"`
	PathURL       string                    `json:"path_url,omitempty"`
	AssetURL      string                    `json:"asset_url,omitempty"`
	Type          DesignAssetType           `json:"type,omitempty"`
	Width         int32                     `json:"width,omitempty"`
	Height        int32                     `json:"height,omitempty"`
}

type DesignAssetPropertyData struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
