package entities

import (
	"time"
)

type LayoutJobStatus string

const (
	LayoutJobStatusPending  LayoutJobStatus = "pending"
	LayoutJobStatusStarted  LayoutJobStatus = "started"
	LayoutJobStatusFinished LayoutJobStatus = "finished"
	LayoutJobStatusError    LayoutJobStatus = "error"
	LayoutJobStatusUnknown  LayoutJobStatus = "unknown"
)

func (s LayoutJobStatus) String() string {
	return string(s)
}

type LayoutJob struct {
	ID              int64           `json:"id,omitempty"`
	BasedOnLayoutID int32           `json:"layout_id,omitempty"`
	CreatedLayoutID int32           `json:"created_layout_id,omitempty"`
	UserID          int32           `json:"user_id,omitempty"`
	TemplateID      int32           `json:"template_id,omitempty"`
	AdaptationID    int32           `json:"adaptation_id,omitempty"`
	Status          LayoutJobStatus `json:"status,omitempty"`
	CreatedAt       time.Time       `json:"created_at,omitempty"`
	StartedAt       time.Time       `json:"started_at,omitempty"`
	StoppedAt       time.Time       `json:"stopped_at,omitempty"`
	FinishedAt      time.Time       `json:"finished_at,omitempty"`
	ImageURL        string          `json:"image_url,omitempty"`
	ErrorAt         time.Time       `json:"error_at,omitempty"`
	UpdatedAt       time.Time       `json:"updated_at,omitempty"`
	Log             string          `json:"log,omitempty"`
	Config          LayoutJobConfig `json:"config,omitempty"`
}
