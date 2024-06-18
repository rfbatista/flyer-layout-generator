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
	}
	return isRunning
}

func (l *LayoutRequest) TotalJobsText() string {
	return fmt.Sprintf("%d", len(l.Jobs))
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
		return ""
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

type LayoutRequestConfig struct {
	LimitSizerPerElement  bool           `json:"limit_sizer_per_element,omitempty"`
	AnchorElements        bool           `json:"anchor_elements,omitempty"`
	ShowGrid              bool           `json:"show_grid,omitempty"`
	MinimiumComponentSize int32          `json:"minimium_component_size,omitempty"`
	MinimiumTextSize      int32          `json:"minimium_text_size,omitempty"`
	SlotsX                int32          `json:"slots_x,omitempty"`
	SlotsY                int32          `json:"slots_y,omitempty"`
	Grid                  Grid           `json:"grid,omitempty"`
	Padding               int32          `json:"padding,omitempty"`
	KeepProportions       bool           `json:"keep_proportions,omitempty"`
	Priorities            map[string]int `json:"priorities,omitempty"`
}

type LayoutRequestJob struct {
	ID         int32                `json:"id,omitempty"`
	RequestID  int32                `json:"request_id,omitempty"`
	TemplateID int32                `json:"template_id,omitempty"`
	CreatedAt  *time.Time           `json:"created_at,omitempty"`
	StartedAt  *time.Time           `json:"started_at,omitempty"`
	StoppedAt  *time.Time           `json:"stopped_at,omitempty"`
	FinishedAt *time.Time           `json:"finished_at,omitempty"`
	ImageURL   string               `json:"image_url,omitempty"`
	Config     *LayoutRequestConfig `json:"config,omitempty"`
	ErrorAt    *time.Time           `json:"error_at,omitempty"`
	Status     string               `json:"status,omitempty"`
	Log        string               `json:"log,omitempty"`
}

func (l *LayoutRequestJob) StatusMessage() string {
	return l.Status
}

func (l *LayoutRequestJob) CreatedAtText() string {
	if l.CreatedAt == nil {
		return ""
	}
	return l.CreatedAt.Format(timeformat)
}

func (l *LayoutRequestJob) StartedAtText() string {
	if l.StartedAt == nil {
		return ""
	}
	return l.StartedAt.Format(timeformat)
}

func (l *LayoutRequestJob) StoppedAtText() string {
	if l.StoppedAt == nil {
		return ""
	}
	return l.StoppedAt.Format(timeformat)
}

func (l *LayoutRequestJob) ErrorAtText() string {
	if l.ErrorAt == nil {
		return ""
	}
	return l.ErrorAt.Format(timeformat)
}

func (l *LayoutRequestJob) FinishedAtText() string {
	if l.ErrorAt != nil {
		return l.ErrorAtText()
	}
	if l.FinishedAt == nil {
		return ""
	}
	return l.FinishedAt.Format(timeformat)
}

func (l *LayoutRequestJob) IsFailure() bool {
	return l.ErrorAt != nil
}

func (l *LayoutRequestJob) IsCompleted() bool {
	return l.FinishedAt != nil
}

func (l *LayoutRequestJob) IsRunning() bool {
	return l.FinishedAt == nil && l.StartedAt != nil && l.ErrorAt == nil
}

func (l *LayoutRequestJob) DurationText() string {
	if l.ErrorAt != nil {
		return time.Time{}.Add(l.ErrorAt.Sub(*l.StartedAt)).Format("5.000s")
	}
	if l.FinishedAt != nil {
		return time.Time{}.Add(l.FinishedAt.Sub(*l.StartedAt)).Format("5.000s")
	}
	return ""
}
