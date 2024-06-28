package entities

import "time"

type Advertiser struct {
	ID         int32      `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	DeleteedAt *time.Time `json:"deleteed_at,omitempty"`
}
