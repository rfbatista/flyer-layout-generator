package entities

import "time"

type LayoutRequest struct {
	ID        int32     `json:"id,omitempty"`
	DesignID  int32     `json:"design_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	StartedAt time.Time `json:"started_at,omitempty"`
	StoppedAt time.Time `json:"stopped_at,omitempty"`
	ErrorAt   time.Time `json:"error_at,omitempty"`
	Status    string    `json:"status,omitempty"`
}

type LayoutRequestConfig struct {
	LimitSizerPerElement  bool  `json:"limit_sizer_per_element,omitempty"`
	AnchorElements        bool  `json:"anchor_elements,omitempty"`
	ShowGrid              bool  `json:"show_grid,omitempty"`
	MinimiumComponentSize int32 `json:"minimium_component_size,omitempty"`
	MinimiumTextSize      int32 `json:"minimium_text_size,omitempty"`
	SlotsX                int32 `json:"slots_x,omitempty"`
	SlotsY                int32 `json:"slots_y,omitempty"`
	Grid                  Grid  `json:"grid,omitempty"`
	Padding               int32 `json:"padding,omitempty"`
	KeepProportions       bool  `json:"keep_proportions,omitempty"`
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
