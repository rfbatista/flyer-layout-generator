package entities

import (
	"fmt"
	"time"
)

type LayoutRequestStatus int

const (
	RequestStatusNotStarted LayoutRequestStatus = iota
	RequestStatusRunning
	RequestStatusStopped
	RequestStatusFinished
	RequestStatusError
)

func (s LayoutRequestStatus) String() string {
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

func (s LayoutRequestStatus) Text() string {
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

type LayoutRequest struct {
	ID         int32              `json:"id,omitempty"`
	DesignID   int32              `json:"design_id,omitempty"`
	CreatedAt  *time.Time         `json:"created_at,omitempty"`
	Total      int32              `json:"total"`
	Done       int32              `json:"done"`
	StartedAt  *time.Time         `json:"started_at,omitempty"`
	StoppedAt  *time.Time         `json:"stopped_at,omitempty"`
	ErrorAt    *time.Time         `json:"error_at,omitempty"`
	FinishedAt *time.Time         `json:"finished_at,omitempty"`
	Jobs       []LayoutRequestJob `json:"jobs,omitempty"`
	Status     string             `json:"status,omitempty"`
}

func (l *LayoutRequest) StatusMessage() string {
	var isRunning bool
	var isFinished bool
	for _, j := range l.Jobs {
		if j.IsRunning() {
			isRunning = true
		}
		if j.IsCompleted() {
			isFinished = true
		}
	}
	if isRunning {
		return "Em execução"
	}
	if isFinished && !isRunning {
		return "Concluido"
	}
	return "Não iniciado"
}

func (l *LayoutRequest) CreatedAtText() string {
	if l.CreatedAt == nil {
		return ""
	}
	return l.CreatedAt.Format(timeformat)
}

func (l *LayoutRequest) StartedAtText() string {
	if l.StartedAt == nil {
		return ""
	}
	return l.StartedAt.Format(timeformat)
}

func (l *LayoutRequest) StoppedAtText() string {
	if l.StoppedAt == nil {
		return ""
	}
	return l.StoppedAt.Format(timeformat)
}

func (l *LayoutRequest) ErrorAtText() string {
	if l.ErrorAt == nil {
		return ""
	}
	return l.ErrorAt.Format(timeformat)
}

func (l *LayoutRequest) FinishedAtText() string {
	if !l.IsCompleted() {
		return ""
	}
	var isRunning bool
	var finishedAt *time.Time
	for _, j := range l.Jobs {
		if finishedAt == nil {
			finishedAt = j.FinishedAt
		}
		if j.IsRunning() {
			isRunning = true
		}
		if j.FinishedAt == nil {
			continue
		}
		if finishedAt.Before(*j.FinishedAt) {
			finishedAt = j.FinishedAt
		}
	}
	if l.ErrorAt != nil {
		return l.ErrorAtText()
	}
	if isRunning || finishedAt == nil {
		return ""
	}
	return finishedAt.Format(timeformat)
}

func (l *LayoutRequest) IsFailure() bool {
	return l.ErrorAt != nil
}

func (l *LayoutRequest) IsCompleted() bool {
	if l.IsRunning() {
		return false
	}
	var isRunning bool
	var isFinished bool
	for _, j := range l.Jobs {
		if j.IsRunning() {
			isRunning = true
		}
		if j.IsCompleted() {
			isFinished = true
		}
	}
	return isFinished && !isRunning
}

func (l *LayoutRequest) IsRunning() bool {
	var isRunning bool
	for _, j := range l.Jobs {
		if j.IsRunning() {
			isRunning = true
		}
		if j.NotStarted() {
			isRunning = false
		}
	}
	return isRunning
}

func (l *LayoutRequest) TotalJobsText() string {
	return fmt.Sprintf("%d", len(l.Jobs))
}

func (l *LayoutRequest) TotalFinishedJobsText() string {
	finished := 0
	for _, j := range l.Jobs {
		if j.IsCompleted() {
			finished += 1
		}
	}
	return fmt.Sprintf("%d", finished)
}

func (l *LayoutRequest) DurationText() string {
	var isRunning bool
	var finishedAt *time.Time
	var startedAt *time.Time
	for _, j := range l.Jobs {
		if finishedAt == nil {
			finishedAt = j.FinishedAt
		}
		if startedAt == nil {
			startedAt = j.StartedAt
		}
		if j.IsRunning() {
			isRunning = true
		}
		if j.FinishedAt == nil {
			continue
		}
		if finishedAt.Before(*j.FinishedAt) {
			finishedAt = j.FinishedAt
		}
		if j.StartedAt == nil {
			continue
		}
		if startedAt.After(*j.StartedAt) {
			startedAt = j.StartedAt
		}
	}
	if isRunning {
		if finishedAt != nil && startedAt != nil {
			t := time.Time{}.Add(finishedAt.Sub(*startedAt))
			return fmt.Sprintf("%dm%d", t.Minute(), t.Second())
		}
	}
	if finishedAt != nil {
		return time.Time{}.Add(finishedAt.Sub(*startedAt)).Format("5.000s")
	}
	return ""
}

func NewLayoutRequestConfigPriority(pr []string) map[string]int {
	m := make(map[string]int)
	for idx, s := range pr {
		m[s] = idx
	}
	return m
}
