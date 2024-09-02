package entities

import (
	"time"
)

type RenderJobStatus string

const (
	RenderJobStatusPending  RenderJobStatus = "pending"
	RenderJobStatusStarted  RenderJobStatus = "started"
	RenderJobStatusFinished RenderJobStatus = "finished"
	RenderJobStatusError    RenderJobStatus = "error"
	RenderJobStatusUnknown  RenderJobStatus = "unknown"
)

func (s RenderJobStatus) String() string {
	return string(s)
}

type RenderJob struct {
	ID           int64           `json:"id,omitempty"`
	LayoutID     int32           `json:"layout_id,omitempty"`
	AdaptationID int32           `json:"adaptation_id,omitempty"`
	Status       RenderJobStatus `json:"status,omitempty"`
	CreatedAt    time.Time       `json:"created_at,omitempty"`
	StartedAt    time.Time       `json:"started_at,omitempty"`
	StoppedAt    time.Time       `json:"stopped_at,omitempty"`
	FinishedAt   time.Time       `json:"finished_at,omitempty"`
	ImageURL     string          `json:"image_url,omitempty"`
	ErrorAt      time.Time       `json:"error_at,omitempty"`
	Log          string          `json:"log,omitempty"`
}
