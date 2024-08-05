package entities

type CompanyAPICredential struct {
	ID       int32  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	APIKey   string `json:"api_key,omitempty"`
	CompayID int32  `json:"compay_id,omitempty"`
}
