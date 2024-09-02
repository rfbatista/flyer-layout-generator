package shared

import "context"

type ApplicationEvent struct {
	ID      string
	Receipt string
	Body    string
}

type Consumer = func(ctx context.Context, event *ApplicationEvent) error
