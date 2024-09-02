package entities

import "time"

type User struct {
	ID         int32
	Name       string
	Email      string
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	DeleteedAt *time.Time `json:"deleteed_at,omitempty"`
}
