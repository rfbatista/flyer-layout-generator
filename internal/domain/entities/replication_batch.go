package entities

import (
	"time"
)

type ReplicationBatchStatus int

const (
	RequestStatusNotStarted ReplicationBatchStatus = iota
	RequestStatusRunning
	RequestStatusStopped
	RequestStatusFinished
	RequestStatusError
)

func (s ReplicationBatchStatus) String() string {
	switch s {
	case RequestStatusNotStarted:
		return "not_started"
	case RequestStatusRunning:
		return "running"
	case RequestStatusStopped:
		return "stopped"
	case RequestStatusFinished:
		return "finished"
	case RequestStatusError:
		return "error"
	}
	return "unknown"
}

func (s ReplicationBatchStatus) Text() string {
	switch s {
	case RequestStatusNotStarted:
		return "Não iniciado"
	case RequestStatusRunning:
		return "Excutando"
	case RequestStatusStopped:
		return "Pausado"
	case RequestStatusFinished:
		return "Finalizado"
	case RequestStatusError:
		return "Error"
	}
	return "Desconhecido"
}

type ReplicationBatch struct {
	ID         int32               `json:"id,omitempty"`
	DesignID   int32               `json:"design_id,omitempty"`
	CreatedAt  *time.Time          `json:"created_at,omitempty"`
	Total      int32               `json:"total,omitempty"`
	Done       int32               `json:"done,omitempty"`
	StartedAt  *time.Time          `json:"started_at,omitempty"`
	StoppedAt  *time.Time          `json:"stopped_at,omitempty"`
	ErrorAt    *time.Time          `json:"error_at,omitempty"`
	FinishedAt *time.Time          `json:"finished_at,omitempty"`
	Config     LayoutRequestConfig `json:"config,omitempty"`
	Jobs       []LayoutRequestJob  `json:"jobs,omitempty"`
	Status     string              `json:"status,omitempty"`
}

func NewLayoutRequestConfigPriority(pr []string) map[string]int {
	m := make(map[string]int)
	for idx, s := range pr {
		m[s] = idx
	}
	return m
}
