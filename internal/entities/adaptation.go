package entities

import (
	"time"
)

type AdaptationBatchStatus string

const (
	AdaptationBatchStatusPending  AdaptationBatchStatus = "pending"
	AdaptationBatchStatusStarted  AdaptationBatchStatus = "started"
	AdaptationBatchStatusFinished AdaptationBatchStatus = "finished"
	AdaptationBatchStatusError    AdaptationBatchStatus = "error"
	AdaptationBatchStatusClosed   AdaptationBatchStatus = "closed"
)

func (s AdaptationBatchStatus) String() string {
	return string(s)
}

type AdaptationBatch struct {
	ID         int64 `json:"id,omitempty"`
	UserID     int64
	DesignID   int32                 `json:"design_id,omitempty"`
	CreatedAt  time.Time             `json:"created_at,omitempty"`
	Total      int32                 `json:"total,omitempty"`
	Done       int32                 `json:"done,omitempty"`
	StartedAt  time.Time             `json:"started_at,omitempty"`
	StoppedAt  time.Time             `json:"stopped_at,omitempty"`
	UpdatedAt  time.Time             `json:"updated_at,omitempty"`
	ErrorAt    time.Time             `json:"error_at,omitempty"`
	FinishedAt time.Time             `json:"finished_at,omitempty"`
	Jobs       []AdaptationJob       `json:"jobs,omitempty"`
	Status     AdaptationBatchStatus `json:"status,omitempty"`
	LayoutID   int32                 `json:"layout_id,omitempty"`
	RequestID  int32                 `json:"request_id,omitempty"`
	TemplateID int32                 `json:"template_id,omitempty"`
	Log        string                `json:"log,omitempty"`
}
