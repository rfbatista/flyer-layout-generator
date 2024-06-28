package entities

import "time"

type Project struct {
	ID         int32      `json:"id,omitempty"`
	Client     Client     `json:"client,omitempty"`
	Advertiser Advertiser `json:"advertiser,omitempty"`
	Name       string     `json:"name,omitempty"`
	Briefing   string     `json:"briefing,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	DeleteedAt *time.Time `json:"deleteed_at,omitempty"`
}
