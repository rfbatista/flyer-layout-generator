package entities

import "time"

type Advertiser struct {
	ID         int64      `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	CompanyID  int32      `json:"company_id,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	DeleteedAt *time.Time `json:"deleteed_at,omitempty"`
}
