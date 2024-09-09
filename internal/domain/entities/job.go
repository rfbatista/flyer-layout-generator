package entities

import (
	"time"
)

type AdaptationBatchStatus string

const (
	AdaptationBatchStatusPending         AdaptationBatchStatus = "pending"
	AdaptationBatchStatusStarted         AdaptationBatchStatus = "started"
	AdaptationBatchStatusFinished        AdaptationBatchStatus = "finished"
	AdaptationBatchStatusRenderingImages AdaptationBatchStatus = "rendering_images"
	AdaptationBatchStatusError           AdaptationBatchStatus = "error"
	AdaptationBatchStatusCanceled        AdaptationBatchStatus = "canceled"
	AdaptationBatchStatusClosed          AdaptationBatchStatus = "closed"
	AdaptationBatchStatusUnknown         AdaptationBatchStatus = "unknown"
)

type JobType string

const (
	JobTypeAdaptation  JobType = "adaptation"
	JobTypeReplication JobType = "replication"
	JobTypeUnknown     JobType = "unknown"
)

func (s AdaptationBatchStatus) String() string {
	return string(s)
}

type Job struct {
	ID              int64                 `json:"id,omitempty"`
	Type            JobType               `json:"type,omitempty"`
	UserID          int64                 `json:"user_id,omitempty"`
	CreatedAt       time.Time             `json:"created_at,omitempty"`
	StartedAt       time.Time             `json:"started_at,omitempty"`
	StoppedAt       time.Time             `json:"stopped_at,omitempty"`
	UpdatedAt       time.Time             `json:"updated_at,omitempty"`
	ErrorAt         time.Time             `json:"error_at,omitempty"`
	FinishedAt      time.Time             `json:"finished_at,omitempty"`
	Jobs            []RenderJob           `json:"jobs,omitempty"`
	RemovedSimilars bool                  `json:"removed_similars,omitempty"`
	Status          AdaptationBatchStatus `json:"status,omitempty"`
	LayoutID        int32                 `json:"layout_id,omitempty"`
	RequestID       int32                 `json:"request_id,omitempty"`
	TemplateID      int32                 `json:"template_id,omitempty"`
	Summary         JobSummary            `json:"summary,omitempty"`
	Log             string                `json:"log,omitempty"`
}

type JobSummary struct {
	Total int64 `json:"total"`
	Done  int64 `json:"done"`
}
